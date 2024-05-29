sudo apt -y install phylonium zip
wget https://ftp.ncbi.nlm.nih.gov/pub/datasets/command-line/v2/linux-amd64/datasets
chmod +x datasets
if [[ ! -d ~/bin ]]; then
    mkdir ~/bin
    source ~/.profile
fi
mv datasets ~/bin
git clone https://github.com/evolbioinf/biobox
cd biobox
bash scripts/setup.sh
make
cp bin/* ~/bin
cd -
git clone https://github.com/evolbioinf/fur
cd fur
bash scripts/setup.sh
make
cp bin/* ~/bin
cd -
