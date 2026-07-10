prog=./cmd/fintac
tree=../data/eco7k.nwk
neidb=~/Data/neidb/neidb
echo "***991910, O111:H8 from paper, -u"
$prog -t 991910_ $tree
$prog -t 991910_ -u 562_ $tree
echo "***O104:H4"
$prog -t "1038927|1048254|1133852|1133853|1134782" $tree
$prog -t "1038927|1048254|1133852|1133853|1134782" -u 562 $tree
echo "***O157:H7, -u, -H"
$prog -t 83334 $tree
$prog -t 83334_ -u 562_ $tree
$prog -t 83334_ -u 562_ -H $neidb $tree
echo "***O25:H4-ST131, -u, -w"
$prog -t 941322_ $tree
$prog -t 941322_ -u 562_ $tree
echo "***O16:H48 from pangenomics chapter and paper: include K12, -u, -H"
$prog -t 2605619_ $tree
$prog -t '2605619_|83333_' $tree
$prog -t '2605619_|83333_' -u 562_ $tree
$prog -t '2605619_|83333_' -u 562_ -H $neidb $tree
