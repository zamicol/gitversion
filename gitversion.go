//Package gitversion uses git to generate a version string useful for binary versioning.
package gitversion

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

// DefaultFile is the default name for where version informaion is kept.
const DefaultFile string = "VERSION"

// TimeStamp is for using TimeStamp or not.
// For fully deterministic versions, set to false.
var TimeStamp = true

// Version uses git to construct a version string using tag and commit.
// This can be useful where deployed go binaries versioning is important.
//
//
// For a clean git commit, Get() will return a simple hash.
//
//     v0.0.0 026249145dab6c65dbfeedf7d01aa2720f51a815
//
// If there has been any change to tracked files, `uncommitted` will be
// appended to commit hash.
//
//     v0.0.0 026249145dab6c65dbfeedf7d01aa2720f51a815 uncommitted
//
// If there is tag information, the tag name will be prepended before the
// commit hash.
//
//     v1.0.0 026249145dab6c65dbfeedf7d01aa2720f51a815
//
// Or if there are uncommitted changes:
//
//     v1.0.0 026249145dab6c65dbfeedf7d01aa2720f51a815 uncommitted
//
// Now this is a problem with uncommitted, since it will be the last hash.
// TODO allow for deploy dir hasing.
func version() (string, error) {

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

	hash := matches[1]

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

// Now returns time like the go mod format.
//	yyyymmddhhmmss
// Usefull for knowing build date.
func Now() string {
	currentTime := time.Now()
	return (currentTime.Format("20060102150405"))
}

// Write writes out a version file.
// A version file looks like this:
//
//	0.0.1 ef8f94357058ce9cba81909016b138e6d54c0381 uncommitted
//	2017-02-28T19:49:11-0700
//
// Where the first line is the version, the second line is build date.
func Write(f string) (err error) {
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

	// Add timestamp
	if TimeStamp {
		version = version + "\n" + Now()
	}

	file.WriteString(version)
	return nil
}

// Get gets the current version from the VERSION file
func Get(f string) (version string, date string, err error) {
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

// Dir is for deterministic directory versioning.
// See https://github.com/golang/mod/blob/ce943fd02449f621243c9ea6e64098e84752b92b/sumdb/dirhash/hash_test.go#L71
func Dir(path string) string {
	hash := sha256.New()
	var previous []byte

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		fmt.Println("Walking : " + path)
		if err != nil {
			fmt.Println("err")
			return err
		}

		if !info.Mode().IsRegular() {
			fmt.Println("Not regular")
			fmt.Println(info)
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		previous = append(previous, data...)
		previous = hash.Sum(previous)

		fmt.Println("Previous: " + fmt.Sprintf("%x", previous))

		return nil
	})

	fmt.Println("Outside walk")

	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", previous)

}
