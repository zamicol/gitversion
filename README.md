# gitversion
Use git for binary versioning.  

See the godoc at https://godoc.org/github.com/zamicol/gitversion


## Quickstart

` import github.com/zamicol/gitversion`

## Examples
Get uses git to construct a version string using tag and commit.

For a clean git commit, Get() will return a simple hash.

    026249145dab6c65dbfeedf7d01aa2720f51a815


If there has been any change to tracked files, `(uncommited)` will be
appended to commit hash.

    026249145dab6c65dbfeedf7d01aa2720f51a815 (uncommited)

If there is tag information, the tag name will be prepended before the
commit hash.

    v1.0 026249145dab6c65dbfeedf7d01aa2720f51a815

Or if there are uncommited changes:

    v1.0 026249145dab6c65dbfeedf7d01aa2720f51a815 (uncommited)
