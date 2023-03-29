## Introduction

This container includes all necessary tools for marker development using whole genome sequences 
(from https://github.com/EvolBioInf/) pre-installed together with all required dependencies: 

- `Neighbors` package (find neighbor genomes based on NCBI taxonomy)
- `phylonium` (fast phylogeny reconstruction from multiple genomic sequences)
- `fur` package (find unique genomic regions)
- `biobox` package (a comprehensive toolset for working with sequences, trees, etc.)

Auxiliary tools on-board:

- `edirect` package (access NCBI databases from Unix CLI)
- `blast+` (pairwise sequence alignment)
- `gnuplot` (draw plots, in particular, phylogenetic trees)

## Installation
    
To install the container:

    docker pull itsers/neighbors

## Preparing host machine
    
Enable X11 connections to allow docker containers to display plots on the screen:

    xhost +

This message should appear upon successful execution:

    'access control disabled, clients can connect from any host'.

## Running the container
    
**Minimalistic command** to run the container in CLI interactive mode:

    docker run -it itsers/neighbors

When container is turned on successfully, you will see your username changes to `jdoe` in the current shell instance:

    >jdoe@your_machine_name:~$

The initial sudo password for jdoe is: `password`.
To turn the container off and return to your host shell, press `Ctrl+D`.

**Full functionality (recommended)** of the container is revealed with the following options:

    docker run -it --env="DISPLAY" --net=host -v ~/neighbors_share:/home/jdoe/neighbors_share itsers/neighbors

Flag info:
`--env="DISPLAY"` and `--net=host`: this helps the container to connect to the host screen (needed to display graphic output)

`-v ~/neighbors_share:/home/jdoe/neighbors`: this makes a folder shared between the neighbors container (`/home/jdoe/neighbors_share`) and your `/home` directory (`/home/username/neighbors_share`). With this shared folder, you can use the neighbors container to operate with your own data and save the results on your machine that hosts the container. Just make sure that you have saved your results in the `/home/jdoe/neighbors_share` direcotry BEFORE exiting the container.

## Tutorial
    
There is a plain-text tutorial on `neighbors` inside the container. When the container is on, run

    more tutorial.txt

Press `Enter` to scroll the text down. Press `Up arrow (↑)` to scroll the text up.
You may exit `more` (press `Ctrl+C`) to copy and paste commands from the tutorial in the container command line.

## Contacts

If you experience some problems with the container, contact Ivan Tsers: `tsers@evolbio.mpg.de`.
To report bugs and errors from `neighbors`, `fur`, or `biobox`, please open an issue in the respective GitHub repos: `https://github.com/EvolBioInf/neighbors`, `https://github.com/EvolBioInf/fur`, or `https://github.com/EvolBioInf/biobox`.

## References

**Haubold, B., Tsers, I., & Denker, S. (2022)**. Neighbors: Using Taxonomic Neighborhood for Marker Discovery (*forthcoming*)

**Haubold, B., Klötzl, F., Hellberg, L., Thompson, D., & Cavalar, M. (2021)**. Fur: Find unique genomic regions for diagnostic PCR. *Bioinformatics, 37(15), 2081-2087*. (https://doi.org/10.1093/bioinformatics/btab059)

**Klötzl, F., & Haubold, B. (2020)**. Phylonium: fast estimation of evolutionary distances from large samples of similar genomes. *Bioinformatics, 36(7), 2040-2046*. (https://doi.org/10.1093/bioinformatics/btz903)

https://github.com/EvolBioInf/biobox

**Kans, J. (2022)**. Entrez direct: E-utilities on the UNIX command line. *In Entrez Programming Utilities Help [Internet]. National Center for Biotechnology Information (US)*. (https://www.ncbi.nlm.nih.gov/books/NBK179288/)

**National Center for Biotechnology Information (US), & Camacho, C. (2008)**. BLAST (r) Command Line Applications User Manual (p. 30). *National Center for Biotechnology Information (US)*. (https://www.ncbi.nlm.nih.gov/books/NBK279690/)
