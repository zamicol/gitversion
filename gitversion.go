// Package gitversion uses git to generate a version string useful for binary
// versioning.
//
// This can be useful where deployed go binaries versioning is important.
//
// For a clean git commit, Get() will return a simple hash.
//
//	v0.0.0 26249145DAB6C65DBFEEDF7D01AA2720F51A815
//
// If there has been any change to tracked files, `uncommitted` will be appended
// to commit hash.
//
//	v0.0.0 26249145DAB6C65DBFEEDF7D01AA2720F51A815 uncommitted
//
// If there is tag information, the tag name will be prepended before the commit
// hash.
//
//	v1.0.0 26249145DAB6C65DBFEEDF7D01AA2720F51A815
//
// Or if there are uncommitted changes:
//
//	v1.0.0 26249145DAB6C65DBFEEDF7D01AA2720F51A815 uncommitted
package gitversion

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// DefaultName is the default name for VERSION file.
const DefaultName string = "VERSION"

// version uses git to construct a version string using tag and commit.
func Version() (string, error) {

	// commit hash
	cmd := exec.Command("git", "log", "-1") // First log record
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.New("gitversion: git command failed.  (Need to 'git init' and commit?)")
	}

	// git log will return a few lines, We only want the line about commit.
	var reg = regexp.MustCompile(`^commit (\w{40})`)
	matches := reg.FindStringSubmatch(out.String())

	if len(matches) != 2 || len(matches[1]) != 40 {
		return "", errors.New("gitversion: unable to get git commit hash")
	}

	hash := strings.ToUpper(matches[1]) // Upper case Hex.

	// tag
	cmd = exec.Command("git", "tag", "--sort=-taggerdate") // If there are no tags, git returns nothing

	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", errors.New("gitversion: problem with git tag")
	}

	reg = regexp.MustCompile(`\A.*`) // Grab first line
	tag := reg.FindString(out.String())
	if tag != "" {
		tag = tag + " "
	}

	// uncommitted changes
	cmd = exec.Command("git", "status")
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", errors.New("gitversion: problem with git status")
	}

	reg = regexp.MustCompile(`(Changes to be committed)|(Changes not staged for commit)`)
	uncommitted := reg.FindString(out.String())
	if uncommitted != "" {
		uncommitted = " uncommitted"
	}

	// construct our version and return.
	return tag + hash + uncommitted, nil
}

// Now returns human readable time.
// 2006/01/02 15:04:05
// for comparison, go mod uses the format: yyyymmddhhmmss
func Now() string {
	currentTime := time.Now()
	return (currentTime.Format("2006/01/02 15:04:05"))
}

// Write writes out a version file.
// A version file looks like this:
//
//	0.0.1 EF8F94357058CE9CBA81909016B138E6D54C0381 uncommitted
//	2006/01/02 15:04:05
//
// Where the first line is the version, the second line is build date.
func Write(f string) (err error) {
	if f == "" {
		f = DefaultName
	}

	// O_WRONLY is write only.
	file, err := os.OpenFile(f, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		return (err)
	}
	version, err := version()
	if err != nil {
		return (err)
	}

	file.WriteString(version + "\n" + Now())
	return nil
}

// Get gets the current version from the VERSION file
func Get(f string) (version string, date string, err error) {
	if f == "" {
		f = DefaultName
	}

	file, err := os.Open(f)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	version = scanner.Text()
	scanner.Scan()
	date = scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", "", err
	}
	return version, date, nil
}

// GetJSON gets from VERSION file and returns as JSON.
func GetJSON(f string) (JSON string, err error) {
	v, d, err := Get(f)
	if err != nil {
		return "", err
	}
	s := strings.Split(v, " ")
	JSON = fmt.Sprintf(`{"tag":"%s","hash":"%s",`, s[0], s[1])
	if len(s) == 3 {
		JSON += fmt.Sprintf(`"committed":"%s",`, s[2])
	}
	JSON += fmt.Sprintf(`"build_date":"%s"}`, d)
	return
}
