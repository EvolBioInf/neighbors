BEGIN{
  # Terminal nodes
  for (i=1; i<=25; i++)
    printf("\\psfrag{ T%d}{\\rnode{t%d}{$T_{%d}$}}\n", i, i, i)
  # Internal nodes
  for (i=1; i<=24; i++)
    printf("\\psfrag{ %d}{\\rnode{n%d}{\\footnotesize$%d$}}\n", i, i, i)
}
