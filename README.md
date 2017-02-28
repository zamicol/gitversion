# gitversion


[Godoc](https://godoc.org/github.com/zamicol/gitversion)

Use git for semantic like project versioning.  

The "version" is
  1. the tag if exists,
  2. The git commit hash
  3. The string "uncommitted" if the current build is not committed.
For example:

   `1.0.2 136540c6f09ba9783c6d6de89a7d9f2e038f6c26 uncommitted`

Since tag will be prepended, tag can be useful to specify a semantic version.
Use tag as normal as the semantic version.  

   `git tag 0.0.1`

Will result in:

   0.0.1 ef8f94357058ce9cba81909016b138e6d54c0381


## Go Quickstart

` import github.com/zamicol/gitversion`
