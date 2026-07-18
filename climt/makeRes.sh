f="test.nwk"
s="303"
p="./cmd/climt"
$p $s $f > r1.txt
$p -d 2 $s $f > r2.txt
$p -D "|" -d 2 $s $f > r3.txt
s="^30[34]$"
$p -r $s $f > r4.txt
