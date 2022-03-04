BEGIN {
  tmpl = "esearch -db nucleotide -query \"%s [orgn] complete genome [titl]\" | efetch -format acc"
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
  acc = ""
  cmd = sprintf(tmpl, orgn)
  while(cmd | getline)
    acc = " " $1
  close(cmd)
  if(length(acc) > 0)
    printf "%s\t%s\t%s\t%s\n", type, id, orgn, acc
}
