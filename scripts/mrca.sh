# Calcualte most recent common ancestor from a set of ancestries.
# Example:
## Get ancestries of rat, mouse and human
# ants 10117 data/neidb > rr.txt # Rattus rattus
# ants 10090 data/neidb > mm.txt # Mus musculus
# ants 9606  data/neidb > hs.txt # Homo sapiens
## Calculate MRCA
# bash scripts/mrca rr.txt mm.txt
# bash scripts/mrca rr.txt mm.txt hs.txt
paste $@ |
    tail -n +2 |
    sed -E 's/([^ ]) ([^ ])/\1_\2/g'  |
    awk '{for(i=2;i<=NF-3;i+=3)if($i!=$(i+3))exit; print}' |
    tail -n 1 |
    awk '{print $2, $3}'
