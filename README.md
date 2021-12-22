# gitversion

Semantic like project versioning using git.  

The "version" is
  1. The git tag if it exists, if not 'v0.0.0'
  2. The git commit hash.
  3. The string "uncommitted" if the current git repo is not committed.

For example:

    1.0.2 136540C6F09BA9783C6D6DE89A7D9F2E038F6C26 uncommitted

It's recommend `git tag`'s value be set to the semantic version.

For example `git tag 0.0.1` will result in:

    0.0.1 EF8F94357058CE9CBA81909016B138E6D54C0381


## Go Quickstart
[Godoc](https://godoc.org/github.com/zamicol/gitversion)

` import github.com/zamicol/gitversion`


# Why when there's `go mod`?

- go mod does not use git's hashes in total for versioning.
- There is no command to easily get the current project's mod information.
- For running binaries on deployment servers.
- Does not rehash the project.  



