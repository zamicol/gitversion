//Package gitversion uses git to generate a version string useful for binary versioning.
package gitversion

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
)

// Get uses git to construct a version string using tag and commit.
// This can be useful where deployed go binaries versioning is important.
//
//
// For a clean git commit, Get() will return a simple hash.
//
//     026249145dab6c65dbfeedf7d01aa2720f51a815
//
// If there has been any change to tracked files, `(uncommited)` will be
// appended to commit hash.
//
//     026249145dab6c65dbfeedf7d01aa2720f51a815 (uncommited)
//
// If there is tag information, the tag name will be prepended before the
// commit hash.
//
//     v1.0.0 026249145dab6c65dbfeedf7d01aa2720f51a815
//
// Or if there are uncommited changes:
//
//     v1.0.0 026249145dab6c65dbfeedf7d01aa2720f51a815 (uncommited)
//
func Get() (string, error) {

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
	uncommited := reg.FindString(out.String())
	if uncommited != "" {
		uncommited = " (uncommited)"
	}

	// construct our version and return.
	return tag + hash + uncommited, nil
}
