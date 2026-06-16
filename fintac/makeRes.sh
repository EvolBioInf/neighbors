f="./test.nwk"
p="./fintac"
$p $f > r1.txt
$p -a $f > r2.txt
$p -t "^tGC[AF]" $f > r3.txt
$p -n "^nGC[AF]" $f > r4.txt
$p -u "^n" $f > r5.txt
$p -w 0 $f > r6.txt
