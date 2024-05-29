wget guanine.evolbio.mpg.de/neighbors/neidb_210923.tgz
tar -xvzf neidb_210923.tgz
rm neidb_210923.tgz
cd neidb_210923
makeNeiDb
cd ../
mv neidb_210923/neidb .
rm -r neidb_210923
