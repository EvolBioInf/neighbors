#!/bin/bash

# install dependencies and additional programs
apt-get update
apt-get -y upgrade
apt-get -y install apt-utils sudo wget git make cmake unzip \
	autoconf build-essential less pkg-config \
	libgsl-dev libdivsufsort-dev libdivsufsort3 \
	libbsd-dev libbsd-dev libgsl-dev libsdsl-dev \
	ncbi-blast+ graphviz
apt-get clean

# find and install current golang version
gotar=$(wget -O- https://go.dev/dl/?mode=json | grep -o 'go.*.linux-amd64.tar.gz' | head -n 1)
wget https://go.dev/dl/$gotar
tar -C /usr/local -xzf $gotar
export PATH=$PATH:/usr/local/go/bin
rm $gotar

# install NCBI datasets
wget "https://ftp.ncbi.nlm.nih.gov/pub/datasets/command-line/v2/linux-amd64/datasets"
mv datasets /usr/local/bin/
chmod +x /usr/local/bin/datasets

cd /home/jdoe

# install biobox
git clone https://github.com/evolbioinf/biobox
cd biobox
make
cp bin/* /usr/local/bin
cd /home/jdoe
rm -rf biobox
   
# install neighbors
git clone https://github.com/EvolBioInf/neighbors
cd neighbors
make
cp ./bin/* /usr/local/bin
cd /home/jdoe
rm -rf neighbors

# downlod neighborsDoc with the tutorial
cd /home/jdoe
wget -O neighborsDoc.pdf https://owncloud.gwdg.de/index.php/s/iBtUxwYqrUA2rbj/download

# install phylonium
git clone https://github.com/evolbioinf/phylonium
cd phylonium
autoreconf -fi -Im4
./configure
make
make install
cd /home/jdoe
rm -rf phylonium

# install fur
apt-get -y install gnuplot 
git clone https://github.com/EvolBioInf/fur
cd fur
make
cp bin/* /usr/local/bin
cd /home/jdoe
rm -rf fur

# make a folder in /home/jdoe to share it with the host
mkdir neighbors_share

# remove unnecessary files
rm -rf /install.sh /usr/local/go /root/go* /root/.cache \
    /usr/lib/go* /usr/share/go* /usr/share/icons
