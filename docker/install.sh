#!/bin/bash

# install linux dependencies
apt-get update
apt-get -y upgrade
apt-get -y install apt-utils sudo wget git make cmake autoconf build-essential noweb pkg-config golang
apt-get clean

# install graphviz to make plots
apt-get -y install graphviz

# install Entrez Direct (efetch)
cd /home/jdoe
sh -c "$(wget -q ftp://ftp.ncbi.nlm.nih.gov/entrez/entrezdirect/install-edirect.sh -O -)"
cp /root/edirect/* /usr/local/bin
rm -rf /root/edirect

# install neighbors
git clone https://github.com/EvolBioInf/neighbors.git
cd neighbors
make
cp ./bin/* /usr/local/bin
cd /home/jdoe
rm -rf neighbors

# install phylonium
apt-get install -y libdivsufsort-dev libdivsufsort3

git clone https://github.com/evolbioinf/phylonium
cd phylonium
autoreconf -fi -Im4
./configure
make
make install
cd /home/jdoe
rm -rf phylonium

# install macle
git clone https://github.com/simongog/sdsl-lite.git
cd sdsl-lite
./install.sh
cd /home/jdoe
rm -rf sdsl-lite

git clone https://github.com/EvolBioInf/macle.git
cd macle
make
cp ./build/macle /usr/local/bin/macle
cd /home/jdoe
rm -rf macle

# install fur
apt-get -y install gnuplot libbsd-dev libbsd-dev libgsl-dev libsdsl-dev ncbi-blast+ primer3
git clone https://github.com/EvolBioInf/fur
cd fur
make
cp ./build/* /usr/local/bin
cd /home/jdoe
rm -rf fur

# install biobox
git clone https://github.com/evolbioinf/biobox
cd biobox
make
cp ./bin/* /usr/local/bin
cd /home/jdoe
rm -rf biobox

# make a folder in /home/jdoe to share it between the container and the host later
mkdir neighbors_share

# remove unused dependencies
rm -rf /install.sh /usr/local/go /root/go /root/.cache \
    /usr/lib/go-1.15 /usr/share/go-1.15 /usr/share/icons

apt-get -y remove \
    git texlive-latex-extra texlive-fonts-recommended texlive-fonts-extra texlive-base \
    texlive-latex-recommended texlive-pstricks texlive-science
apt-get -y autoremove
apt-get clean

