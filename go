#!/bin/bash
RED="\e[31m"
GREEN="\e[32m"
END="\e[0m"

bindir="/home/$USER/bin/"
inc="/home/$USER/bin/.inc"

myfile="$inc/all.txt"
hcol="$inc/clihead.txt"
lgo="$inc/cclii.txt"
i=1

header() {
cat "$lgo" 
printf "\n"
}

go() {

header
printf "\n"
#printf "=%.0s"  $(seq 1 63) 

printf "\n"
printf  "Enter command keyword: "
read -r kword
printf "\n"
printf "Keyword entered: '${RED}$kword${END}' \n\n"
echo -e "Results:\n"
extract $kword

}

extract(){

grep "$kword" "$myfile" | while read -r line;

do

if [ "$i" = 1 ]; then
cat "$hcol"
fi

echo -e "\n"
printf "$line"
echo -e "\n"
i=$(($i+1))
done | column -t -s "|"
ask
}

ask(){
printf "\n"
read -p "Continue? [y/n] : " opt

if [ "$opt" = "y" ]; then
    cont
else
    echo "Byers"
fi
}

cont() {

printf "\n"
printf "=%.0s"  $(seq 1 63) 

printf "\n"
printf  "Enter command keyword: "
read -r kword
printf "\n"
printf "Keyword entered: '${RED}$kword${END}' \n\n"
echo -e "Results:\n"
extract $kword

}

go
