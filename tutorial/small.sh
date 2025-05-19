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
climt -r 003003865 lpn.nwk
pickle 8 lpn.nwk |
    grep -v '^#' |
    while read a; do
          head -n 1 all/$a
    done
pickle 3 lpn.nwk |
    grep -v '^#' |
    while read a; do
          head -n 1 all/$a
    done
pickle 16 lpn.nwk |
    grep -c '^n'
fintac lpn.nwk
mkdir targets
pickle 16 lpn.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a targets/$a
    done
mkdir neighbors
pickle -c 16 lpn.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a neighbors/$a
    done
makeFurDb -t targets/ -n neighbors/ -d lpn.db
fur -d lpn.db > lpn.fasta
cd ../
