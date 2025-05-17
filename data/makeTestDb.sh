echo "Download raw data"
rm -rf neidb_280425
wget owncloud.gwdg.de/index.php/s/sNVcgHAkc1JHVPq/download \
     -O neidb_280425.tgz
tar -xvzf neidb_280425.tgz
cd neidb_280425

echo "Construct master database"
makeNeiDb

echo "Construct small data sets"
rm -f *Test.*
for taxid in 207598 730 3344738; do
    dree -l $taxid neidb |
	tail -n +2 |
	while read a b; do
	    awk -v t=$a '$1==t' names.dmp
	done >> namesTest.dmp

    dree -l $taxid neidb |
	tail -n +2 |
	while read a b; do
	    awk -v t=$a '$1==t' nodes.dmp
	done >> nodesTest.dmp

    dree -l $taxid neidb |
	tail -n +2 |
	while read a b; do
	    awk -v t=$a '$6==t' assembly_summary_genbank.txt
	done >> gbTest.txt

    dree -l $taxid neidb |
	tail -n +2 |
	while read a b; do
	    awk -v t=$a '$6==t' assembly_summary_refseq.txt
	done >> rsTest.txt
done

# Ensure that Hominiae (207598) is root
sed 's/9604/207598/' nodesTest.dmp > t && mv t nodesTest.dmp

echo "Build test.db"
makeNeiDb -a namesTest.dmp -d test.db -g gbTest.txt -o nodesTest.dmp -r rsTest.txt
cp namesTest.dmp gbTest.txt nodesTest.dmp rsTest.txt ../
cp test.db ../
cp neidb ../
cd ../
