# `neighbors`
## Description
Identify neighbor genomes for marker discovery with
[`fur`](https://github.com/evolbioinf/fur).
## Author
[Bernhard Haubold](http://guanine.evolbio.mpg.de/), `haubold@evolbio.mpg.de`
## Make the Programs
Make sure you've installed the packages `git`, `golang`, `make`, and `noweb`.  
  `$ make`  
  The directory `bin` now contains the binaries, scripts are in
  `scripts`.
## Make the Documentation
Make sure you've installed the packages `git`, `make`, `noweb`, `texlive-science`,
`texlive-pstricks`, `texlive-latex-extra`,
and `texlive-fonts-extra`.  
  `$ make doc`  
  The documentation is now in `doc/neighborsDoc.pdf`.
## Docker Container 
As an alternative to building `neighbors` from scratch, we also post it as a [docker
  container](https://hub.docker.com/r/itsers/neighbors). The container
  includes all programs needed to work through the tutorial in `~/tutorial.txt`.
  -  `$ docker pull itsers/neighbors`
  -  `$ docker container run --detach-keys="ctrl-@" -h neighbors -it itsers/neighbors`
## License
[GNU General Public License](https://www.gnu.org/licenses/gpl.html)
