#!/bin/bash

set -e

echo "Builing package..."
make build

SUBDIR=TDD.sh
DIR=/tmp/$SUBDIR
ARCHIVE=/tmp/TDDsh.tar.gz

echo "Cleaning /tmp..."
rm $ARCHIVE -f
rm $DIR -rf
mkdir $DIR

echo "Copying files to /tmp"
cp _build/linux/tdd $DIR/
cp README.md $DIR/

echo "Creating archive..."
cd /tmp
tar -czvf $ARCHIVE $SUBDIR/*
cd -

echo "Archive available at " $ARCHIVE
echo "Done!"
