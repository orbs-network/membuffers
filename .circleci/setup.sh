#!/bin/bash -e

cd javascript && npm install
cd ..
go get -u github.com/gobuffalo/packr/packr