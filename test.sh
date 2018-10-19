pushd .

echo ""
echo "***** TESTING GO LIBRARY *****"
echo ""
echo "  Running ./go/test_lib.sh"
echo ""

cd ./go
./test_lib.sh

echo ""
echo "***** TESTING COMPILER *****"
echo ""
echo "  Running ./go/membufc/test_compiler.sh"
echo ""

cd ./membufc
./test_compiler.sh

popd

pushd .

echo ""
echo "***** TESTING JS LIBRARY *****"
echo ""
echo "  Running ./javascript/test_lib.sh"
echo ""

cd ./javascript
./test_lib.sh

popd