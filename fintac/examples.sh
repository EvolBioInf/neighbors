p=./cmd/fintac
echo "***991910, O111:H8 from paper, -u"
$p -t 991910 ../data/phyl_cleaned.nwk
$p -t 991910 -u 562 ../data/phyl_cleaned.nwk
echo "***O104:H4"
$p -t "1038927|1048254|1133852|1133853|1134782" ../data/phyl_cleaned.nwk
$p -t "1038927|1048254|1133852|1133853|1134782" -u 562 ../data/phyl_cleaned.nwk
echo "***O157:H7, -u, -H"
$p -t 83334 ../data/phyl_cleaned.nwk
$p -t 83334_ -u 562_ ../data/phyl_cleaned.nwk
$p -t 83334_ -u 562_ -H ~/Data/neidb/neidb ../data/phyl_cleaned.nwk
echo "***O25:H4-ST131, -u, -w"
$p -t 941322_ ../data/phyl_cleaned.nwk
$p -t 941322_ -u 562_ ../data/phyl_cleaned.nwk
$p -t 941322_ -u 562_ -w 3 ../data/phyl_cleaned.nwk
echo "***O16:H48 from pangenomics chapter and paper: include K12, -u, -H"
$p -t 2605619_ ../data/phyl_cleaned.nwk
$p -t '2605619_|83333_' ../data/phyl_cleaned.nwk
$p -t '2605619_|83333_' -u 562_ ../data/phyl_cleaned.nwk
$p -t '2605619_|83333_' -u 562_ -H ~/Data/neidb/neidb ../data/phyl_cleaned.nwk
