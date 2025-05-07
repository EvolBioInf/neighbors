echo "Delete old test.db"
rm -f test.db
echo "Build new test.db"
makeNeiDb -a namesTest.dmp -d test.db -g gbTest.txt -o nodesTest.dmp -r rsTest.txt
