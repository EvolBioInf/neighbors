h=$(history | tail | grep update)
if [[ $h == "" ]]; then
    apt update
fi
s=$(which sudo)
if [[ $s == "" ]]; then
    apt -y install sudo
fi
sudo apt -y install golang make phylonium sqlite3 tar wget
