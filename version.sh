#!/usr/bin/env bash

# This script generates a `VERSION` file.
#
# 1. First line is the "version".
# 2. Second line is the build date.
version=""
#rfc-3339ish, runs on older versions of 'date'
buildDate=$(date +%Y-%m-%dT%T%z)

# commit hash
log=$(git log -1)
# line with 'commit' and a 40 hex character string
reg='^commit ([[:xdigit:]]{40})'
if [[ $log =~ $reg ]]; then
  version="${BASH_REMATCH[1]}"
fi

# tag
# Only show tag in the version number if the current commit is hashed.
headHash=$(git rev-parse HEAD)
tag=`git tag --points-at $headHash`
if [[ ! -z $tag ]]; then
  version="$tag $version"
fi

# Optionally, always show the "latest" tag.
# tag=$(git tag --sort=-taggerdate)
# # grab everything up to the first white space.
# reg='^[^[:space:]]+'
# if [[ $tag =~ $reg ]]; then
#   version="${BASH_REMATCH[0]} $version"
# fi

# uncommited
status=$(git status)
reg='(Changes to be committed)|(Changes not staged for commit)'
if [[ $status =~ $reg ]]; then
  version="$version uncommitted"
fi

# Write this information to a version file in the directory.
echo -e "$version\n$buildDate" > VERSION
