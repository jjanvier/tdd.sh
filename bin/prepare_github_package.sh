#!/bin/bash

set -e

echo "Builing package..."
make build

SUBDIR=TDD.sh
DIR=/tmp/$SUBDIR
ARCHIVE=/tmp/TDD.sh.tar.gz

echo "Cleaning /tmp..."
rm $ARCHIVE -f
rm $DIR -rf
mkdir $DIR

echo "Copying files to /tmp"
mv tdd $DIR/tdd.sh

echo "Creating archive..."
cd /tmp
tar -czvf $ARCHIVE $SUBDIR/*
cd -

echo "Archive available at " $ARCHIVE
echo "Done!"
