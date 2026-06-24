p=./cmd/fintac
# O111:H8 from paper
$p -t 991910 ../data/phyl_cleaned.nwk
$p -t 991910 -u 562 ../data/phyl_cleaned.nwk
# O104:H4 from pangenomics chapter
$p -t "1038927|1048254|1133852|1133853|1134782" ../data/phyl_cleaned.nwk
$p -t "1038927|1048254|1133852|1133853|1134782" -u 562 ../data/phyl_cleaned.nwk
# O157:H7
$p -t 83334 ../data/phyl_cleaned.nwk
$p -t 83334_ -u 562_ ../data/phyl_cleaned.nwk
# O25 want clade 434
$p -t 941322_ ../data/phyl_cleaned.nwk
$p -t 941322_ -u 562_ ../data/phyl_cleaned.nwk
