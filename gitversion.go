package gitversion

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
)

//Get Use git to construct a version number using the latest tag and commit hash
func Get() string {

	////////////////////
	// commit hash
	////////////////////
	cmd := exec.Command("git", "log", "-1") // First log record
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		err = errors.New("git command failed, and unable to get git version.  (Need to 'git init' and commit?)")
		panic(err)
	}

	// git log will return a few lines, We only want the line about commit.
	var reg = regexp.MustCompile(`^commit (\w{40})`)
	matches := reg.FindStringSubmatch(out.String())

	if len(matches) != 2 || len(matches[1]) != 40 {
		err = errors.New("Problem getting git commit. ")
		panic(err)
	}

	hash := matches[1]

	////////////////////
	// tag
	////////////////////
	cmd = exec.Command("git", "tag")
	//If there are no tags, git will return nothing
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	reg = regexp.MustCompile(`\A.*`) //Grab first line
	tag := reg.FindString(out.String()) + " "

	////////////////////
	// uncommitted changes
	////////////////////
	cmd = exec.Command("git", "status")
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	reg = regexp.MustCompile(`Changes to be committed`)
	uncommited := reg.FindString(out.String())
	if uncommited != "" {
		uncommited = " (uncommited)"
	}

	////////////////////
	// construct our version and return.
	////////////////////
	return tag + hash + uncommited
}
