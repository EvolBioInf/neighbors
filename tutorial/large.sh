mkdir large
cd large
neighbors -t 446 -l -L complete,chromosome ../neidb |
    grep '^t' |
    awk '{print $2}' > uacc.txt
wc -l uacc.txt
grep -v -f ../small/tacc.txt uacc.txt > tmp.txt
wc -l tmp.txt
mv tmp.txt uacc.txt
grep -v -f ../small/nacc.txt uacc.txt > tmp.txt
wc -l tmp.txt
mv tmp.txt uacc.txt
datasets download genome accession\
           --inputfile uacc.txt \
           --exclude-atypical \
           --dehydrated \
           --filename udata.zip
unzip udata.zip -d udata
datasets rehydrate --directory udata
cp -r ../small/all .
for a in udata/ncbi_dataset/data/*/*.fna
do
    b=$(basename $a)
    mv $a all/u$b
done
ls all | wc -l
phylonium all/*.fna > lpn.dist
nj lpn.dist | midRoot | land > lpn.nwk
fintac -n "^n" lpn.nwk
pickle 25 lpn.nwk |
    grep -c '^u'
pickle -c 25 lpn.nwk |
    grep -c '^u'
mkdir targets
pickle 25 lpn.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a targets/$a
    done
mkdir neighbors
pickle -c 25 lpn.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a neighbors/$a
    done
ls targets/ | wc -l
ls neighbors/ | wc -l
makeFurDb -t targets/ -n neighbors/ -d lpn.db
fur -d lpn.db > lpn.fasta
cd ../
