#!/bin/sh
./version.sh
version=$(head -1 VERSION)
buildDate=$(tail -1 VERSION)
# Build
`go build -ldflags "-X 'main.Version=$version' -X 'main.BuildDate=$buildDate'"`
