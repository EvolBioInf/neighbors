## Introduction

This container includes pre-compiled programs from the EvolBioInf
github repository (https://github/evolbioinf) for marker sequence
search using whole genome sequences:

- `Neighbors` package (find taxonomic neighbors)
- `phylonium` (fast genetic distance calculation)
- `fur` package (find unique genomic regions)
- `biobox` package (manipulate sequences, trees, etc.)

Auxiliary tools on-board:

- `datasets` package (access NCBI databases from Unix CLI)
- `blast+` (pairwise sequence alignment)
- `gnuplot` (draw plots, in particular, phylogenetic trees)

## Installation
    
To install the container:

    docker pull itsers/neighbors

## Preparing host machine
    
Enable X11 connections to allow docker containers to display plots on
the screen:

    xhost +

This message should appear upon successful execution:

    'access control disabled, clients can connect from any host'.

## Running the container

We recommend these options to run the container:

    docker run -it --env="DISPLAY" --net=host -v
    ~/neighbors_share:/home/jdoe/neighbors_share -h neighbors
    --detach-keys="ctrl-@" itsers/neighbors

When container is turned on successfully, you will see your username
changed to `jdoe`, and the host name changed to `neighbors` in the
current shell instance:

    >jdoe@neighbors:~$

The initial user password for jdoe is: `password`.

To stop the container and return to the host shell, type `exit` or
press `Ctrl+D`.

The flags usage explained:

- `-it` runs the container in CLI interactive mode

- `--env="DISPLAY"` and `--net=host`: these are needed to display
    graphic output from inside the container

- `-v ~/neighbors_share:/home/jdoe/neighbors`: this creates a folder
shared between the neighbors container (`/home/jdoe/neighbors_share`)
and your `/home` directory (`/home/username/neighbors_share`). With
this, you can use the container to operate with your own data and save
the results on your machine that hosts the container. Just make sure
that you have saved your results at `/home/jdoe/neighbors_share`
before stopping the container.

- `--detach-keys="ctrl-@"` is used here to override the default key
sequence (`Ctrl+P, Ctrl+Q`) with `Ctrl+Shift+2` for switching between
interactive mode and daemon (background) mode. This frees the sequence
`Ctrl+P` to be used alongside with `Ctrl+N` to navigate the shell
command history within the container.

`-h neighbors` sets the host name inside the container.

## Documentation and tutorial

There is a file `neighborsDoc.pdf` at `/home/jdoe` describing the
programs of the `neqighbors` package and providing a tutorial on
neighbors-based analysis. You can move or copy the file to the
`/home/jdoe/neighbors_share` folder and open it in your host operating
system. If you want to open the PDF within the container, install a
PDF viewer (for example, evince) inside the container:

	apt-get install evince
	evince neighborsDoc.pdf

Make sure that X11 connections are enabled beforehand (see above).

## Contacts

If you experience problems with the container, please open an issue at
`https://github.com/EvolBioInf/neighbors`.  To report bugs and errors
from `neighbors`, `fur`, or `biobox` specifically, please open issues
in the respective GitHub repos:
`https://github.com/EvolBioInf/neighbors`,
`https://github.com/EvolBioInf/fur`, or
`https://github.com/EvolBioInf/biobox`.

## References
**Haubold, B., Klötzl, F., Hellberg, L., Thompson, D., &
Cavalar, M. (2021)**. Fur: Find unique genomic regions for diagnostic
PCR. *Bioinformatics, 37(15),
2081-2087*. (https://doi.org/10.1093/bioinformatics/btab059)

**Klötzl, F., & Haubold, B. (2020)**. Phylonium: fast estimation of
  evolutionary distances from large samples of similar
  genomes. *Bioinformatics, 36(7),
  2040-2046*. (https://doi.org/10.1093/bioinformatics/btz903)

https://github.com/EvolBioInf/biobox

**Kans, J. (2022)**. Entrez direct: E-utilities on the UNIX command
  line. *In Entrez Programming Utilities Help [Internet]. National
  Center for Biotechnology Information
  (US)*. (https://www.ncbi.nlm.nih.gov/books/NBK179288/)

**National Center for Biotechnology Information (US), & Camacho,
  C. (2008)**. BLAST (r) Command Line Applications User Manual
  (p. 30). *National Center for Biotechnology Information
  (US)*. (https://www.ncbi.nlm.nih.gov/books/NBK279690/)