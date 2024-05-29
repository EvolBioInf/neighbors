taxi -s "O157:H7" neidb | head
taxi -s "O157:H7" neidb |
    tail -n +2 |
    wc -l
taxi -s O157:H7 neidb |
    tail -n +2 |
    awk '{print $2}' |
    sort |
    uniq -c
dree 83334 neidb | dot -T x11
ants 83334 neidb
dree -l 83334 neidb | head
dree -l -n 83334 neidb | head
dree -n -l 83334 neidb |
    tail -n +2 |
    awk '{s += $3}END{print s}'
dree -n -l 83334 neidb |
    tail -n +2 |
    sort -k 3 -n -r
printf 83334 | neighbors -g neidb
printf 83334 |
    neighbors -l neidb |
    tail -n +2 |
    grep '^t' |
    wc -l
printf 83334 | neighbors -l neidb | head
printf 83334 | neighbors -l neidb > acc.txt
grep -c '^n' acc.txt
grep '^t' acc.txt | awk '{print $2}' > tacc.txt
grep '^n' acc.txt | awk '{print $2}' > nacc.txt
datasets download genome accession \
           --inputfile tacc.txt \
           --assembly-level complete \
           --exclude-atypical \
           --dehydrated \
           --filename tdata.zip
datasets download genome accession \
           --inputfile nacc.txt \
           --assembly-level complete \
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
phylonium all/* > o157.dist
nj o157.dist | midRoot | land > o157.nwk
plotTree -d 1200,1200 o157.nwk
sed 's/n[^f]*fna/n/g' o157.nwk |
    plotTree -d 1200,1200
sed 's/n[^f]*fna/n/g' o157.nwk |
    pickle -t 269 |
    plotTree -d 1200,1200 
sed 's/n[^f]*fna/n/g' o157.nwk |
    pickle -t 293 |
    plotTree -d 1200,1200 
pickle 301 o157.nwk |
    grep -c '^n'
sed 's/n[^f]*fna/n/g' o157.nwk |
    pickle -t 292 |
    pickle -t -c 312 |
    plotTree -d 1200,1200
sed -E 's/([nt])[^f]*fna/\1/g' o157.nwk |
    pickle -t 292 |
    pickle -t -c 312 |
    plotTree -d 1200,1200
fintac o157.nwk
fintac -a o157.nwk | head
pickle -c 301 o157.nwk | grep '^t'
mkdir targets
pickle 301 o157.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a $(pwd)/targets/$a
    done
mkdir neighbors
pickle 294 o157.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a $(pwd)/neighbors/$a
    done
makeFurDb -t targets/ -n neighbors/ -d 301_294.db
fur -d 301_294.db > 301_294.fasta
rm neighbors/*
pickle -c 301 o157.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a $(pwd)/neighbors/$a
    done
makeFurDb -t targets -n neighbors -d 301.db
fur -d 301.db
for a in targets/*.fna; do
    echo -n $a ' ';cres $a |
          grep To |
          awk '{print $2}'
done > tlen.dat
awk '{print $2}' tlen.dat | outliers
awk '$2==4.869019e+06' tlen.dat
rm targets/tGCA_030908645.1_ASM3090864v1_genomic.fna
makeFurDb -t targets -n neighbors -o -d 301.db
fur -d 301.db
fur -m -d 301.db
rm neighbors/*
pickle 294 o157.nwk |
        grep -v '^#' |
        while read a; do
            ln -s $(pwd)/all/$a $(pwd)/neighbors/$a
        done
makeFurDb -t targets -n neighbors -d 301_294.db -o
fur -d 301_294.db > 301_294.fasta
rm targets/*
pickle 293 o157.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a $(pwd)/targets/$a
    done
rm targets/tGCA_030908645.1_ASM3090864v1_genomic.fna
rm neighbors/*
pickle -c 293 o157.nwk |
    grep -v '^#' |
    while read a; do
          ln -s $(pwd)/all/$a $(pwd)/neighbors/$a
    done
makeFurDb -t targets -n neighbors -d 293.db
fur -m -d 293.db > 293.fasta
