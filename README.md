# `neighbors`
## Description
Identify target and neighbor genomes for marker discovery.
## Introduction
Genetic markers are genomic regions that are common to a set of target
organisms and absent from all other organisms. This absence is often
tested by running marker candidates against comprehensive sequence
databases like [GenBank](https://www.ncbi.nlm.nih.gov/genbank/). Any
hit outside the targets is interpreted as cross-reactivity and
removed. However, for large candidate sets a search of
[GenBank](https://www.ncbi.nlm.nih.gov/genbank/) can be time
consuming. It is also not necessary, as evolutionary biology tells us
that the vast majority of cross-reactive material is contained in the
targets' closest phylogenetic neighbors. The programs in `neighbors`
help identify the target and neighbor genomes currently
available in [RefSeq](https://www.ncbi.nlm.nih.gov/refseq/). 

Given a sample of target genomes and a sample of neighbor genomes
discovered with `neighbors`, the regions common to the targets that
are absent form the neighbors are good marker candidates. The program
[`fur`](https://github.com/evolbioinf/fur) is one example of a program
for identifying such regions. Once found, the marker candidates can be
further analyzed *in silico* and
*in vitro* to extract genetic markers.  
## Author
[Bernhard Haubold](http://guanine.evolbio.mpg.de/), `haubold@evolbio.mpg.de`
## Make the Programs
Make sure you've installed the packages `git`, `golang`,
`make`, `noweb`, and `sqlite3`.  
  `$ make`  
  The directory `bin` now contains the seven executables of the
  package, additional scripts are in
  `scripts`.
## Make the Documentation
Make sure you've installed the packages `git`, `make`, `noweb`, `texlive-science`,
`texlive-pstricks`, `texlive-latex-extra`,
and `texlive-fonts-extra`. Then execute  
  `$ make doc`  
  The documentation is now in `doc/neighborsDoc.pdf`. 
## Run Docker Container 
As an alternative to building `neighbors` from scratch, we also post it as a [docker
  container](https://hub.docker.com/r/itsers/neighbors). The container
  includes all programs needed to work through the tutorial in `~/tutorial.txt`.
  -  `$ docker pull itsers/neighbors`
  -  `$ docker container run --detach-keys="ctrl-@" -h neighbors -it itsers/neighbors`
## Make the Docker Container
The command  
`$ make docker`  
pulls the repository
`https://github.com/IvanTsers/neighbors-docker`, and starts building a
local copy of the docker image.
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
