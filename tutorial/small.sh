mkdir small
cd small
neighbors -L complete,chromosome -g -t 91891 ../neidb | less
neighbors -L complete,chromosome -g -t 91891 -l ../neidb
neighbors -L complete,chromosome -g -t 91891 \
            -l ../neidb > acc.txt
grep -c '^t' acc.txt
grep -c '^n' acc.txt
grep '^t' acc.txt | awk '{print $2}' > tacc.txt
grep '^n' acc.txt | awk '{print $2}' > nacc.txt
datasets download genome accession \
           --inputfile tacc.txt \
           --exclude-atypical \
           --dehydrated \
           --filename tdata.zip
datasets download genome accession \
           --inputfile nacc.txt \
           --exclude-atypical \
           --dehydrated \
           --filename ndata.zip
unzip tdata.zip -d tdata
unzip ndata.zip -d ndata
datasets rehydrate --directory tdata
datasets rehydrate --directory ndata
mkdir all
for a in tdata/ncbi_dataset/data/*/*.fna
do
    b=$(basename $a)
    mv $a all/t$b
done
for a in ndata/ncbi_dataset/data/*/*.fna
do
    b=$(basename $a)
    mv $a all/n$b
done
phylonium all/* > lpn.dist
nj lpn.dist | midRoot | land > lpn.nwk
sed -E 's/([nt])[^f]*fna/\1/g' lpn.nwk |
    plotTree
head -n 1 all/*.fna |
    grep -B 1 fraseri |
    head -n 2
# Up   Node               Branch Length   Cumul...
5      1                  0               0.032798
4      2                  0.0095          0.023298
3      8                  0.0165          0.006798
2      9                  0.000348        0.00645
1      10                 0.00354         0.00291
0      nGCA_003003865...  0.00291         0
pickle 8 lpn.nwk |
    grep -v '^#' |
    while read a b; do
          head -n 1 all/$a
    done
pickle 3 lpn.nwk |
    grep -v '^#' |
    while read a b; do
          head -n 1 all/$a
    done
pickle 16 lpn.nwk |
    grep -c '^n'
fintac lpn.nwk
mkdir targets
pickle 16 lpn.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a $(pwd)/targets/$a
    done
mkdir neighbors
pickle -c 16 lpn.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a $(pwd)/neighbors/$a
    done
makeFurDb -t targets/ -n neighbors/ -d lpn.db
fur -d lpn.db > lpn.fasta
We are done with our small analysis and return to the parent directory.
cd ../
