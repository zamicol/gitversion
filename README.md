# gitversion

Semantic like project versioning using git.  

The "version" is
  1. The git tag if it exists.
  2. The git commit hash.
  3. The string "uncommitted" if the current git repo is not committed.

For example:

    1.0.2 136540c6f09ba9783c6d6de89a7d9f2e038f6c26 uncommitted

We recommend that the `git tag`'s value be set to the semantic version.

For example `git tag 0.0.1` will result in:

    0.0.1 ef8f94357058ce9cba81909016b138e6d54c0381


## Go Quickstart
[Godoc](https://godoc.org/github.com/zamicol/gitversion)

` import github.com/zamicol/gitversion`
