#!/bin/bash -e

export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

echo "  * Building protos for tests (without building compiler)"
echo ""

go run $(ls -1 ../go/membufc/*.go | grep -v _test.go) -m `find ./types -name "*.proto"`

echo "  * Running tests"
echo ""

go test -count=1 .
