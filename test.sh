pushd .

echo ""
echo "***** TESTING LIBRARY *****"
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