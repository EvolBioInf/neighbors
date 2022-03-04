BEGIN {
  if(db == "") {
    print "Usage: awk -f searchDb.awk -v db=<db> index.out"
    exit
  }
  tq = "select replicons from genome join taxon where taxon.taxid=%d and genome.taxid=taxon.taxid"
  tmpl = "sqlite3 %s \"%s\""
}
!/^#/ {
  id = $2
}
/^t/ {
  type = "t"
  orgn = $4
  for(i = 5; i <= NF; i++)
    orgn = orgn " " $i
}
/^n/ {
  type = "n"
  orgn = $3
  for(i = 4; i <= NF; i++)
    orgn = orgn " " $i
}
!/^#/ {
  rep = ""
  q = sprintf(tq, id)
  cmd = sprintf(tmpl, db, q)
  while(cmd | getline)
    if($1 != "-")
      rep = " " $1
  close(cmd)
  if(length(rep) > 0)
    printf "%s\t%s\t%s\t%s\n", type, id, orgn, rep
}
