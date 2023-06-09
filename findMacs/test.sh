t=$(mktemp tXXX)
n=$(mktemp nXXX)
stan -s 3 -T $t -N $n -o
r=$(mktemp rXXX.out)
./findMacs $t $n > $r
d=$(diff $r r1.txt)
if [ "$d" = "" ]
then
    echo "findMacs OK"
else
    echo "Error: " $d
fi
rm -r $t $n $r
