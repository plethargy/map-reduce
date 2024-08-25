#!/usr/bin/bash

tempFileForPath="tempfileforpath.txt" #name this whatever you want as long as it doesn't interfere with your current structure.
mainGoSrcFile="mapreduce.go"
echo $GOPATH > $tempFileForPath
goPathCount=$(grep -c $(pwd) $tempFileForPath)
goPathEmpty=$(grep -c . $tempFileForPath)
echo $goPathCount
if [ $goPathCount -eq 0 ]; then
	if [ $goPathEmpty -eq 0 ]; then
		export GOPATH=$(pwd) #Need to confirm whether this clears after shell script finishes executing
	else
		export GOPATH="${GOPATH}:$(pwd)"
	fi
else
	echo "It already exists in the path"
fi 
echo "Gopath is set to ${GOPATH}"
rm -rf $tempFileForPath #clean up filep

goBin=$(which go)
modulerootdir=$(pwd)
packageList=("cli" "io" "worker" "partition" "log" "map" "reduce" "emit")

for package in ${packageList[@]}; do 
	cd "${modulerootdir}/src/${package}"
	echo "Running go build in $(pwd)"
	$goBin build
done 

echo "Running go install"
cd "${modulerootdir}/src"
$goBin install $mainGoSrcFile

cp config.json ../bin/config.json
