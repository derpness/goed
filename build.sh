#!/bin/bash

# Bundle file resources into the goed binary and build
date +%s > res/resources_version.txt
go-bindata -pkg core -o core/resources_gen.go res/...

go install ./...
