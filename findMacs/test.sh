t=$(mktemp tXXX)
n=$(mktemp nXXX)
stan -s 3 -T $t -N $n -o
r=$(mktemp rXXX.out)
./findMacs $t $n 2>/dev/null > $r
d=$(diff $r r1.txt)
if [[ "$d" == "" ]]
then
    echo "findMacs - mashmap OK"
else
    echo "Error - mashmap: " $d
fi
./findMacs -t 8 $t $n 2>/dev/null > $r
d=$(diff $r r1.txt)
if [[ "$d" == "" ]]
then
    echo "findMacs - mashmap -t OK"
else
    echo "Error - mashmap -t: " $d
fi
./findMacs -n $t $n > $r
d=$(diff $r r2.txt)
if [[ "$d" == "" ]]
then
    echo "findMacs - nucmer OK"
else
    echo "Error - nucmer " $d
fi
./findMacs -n -t 8 $t $n > $r
d=$(diff $r r2.txt)
if [[ "$d" == "" ]]
then
    echo "findMacs - nucmer -t OK"
else
    echo "Error - nucmer -t " $d
fi
b=$(mktemp bXXX)
bash bedtools.sh > $b
d=$(diff $b r3.txt)
if [[ "$d" == "" ]]
then
    echo "findMacs - bedtools.sh OK"
else
    echo "Error - bedtools.sh " $d
fi
rm -r $t $n $r $b
