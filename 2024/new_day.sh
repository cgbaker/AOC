#!/bin/bash

if [ $# -ne 1 ]; then
	echo "usage: $0 <day>"
	exit 1
fi

NEWDIR=$(printf 'day%2.2d' $1)
if [ -d $NEWDIR ]; then 
	echo "directory $NEWDIR already exists"
	exit 1
fi

cp -r day_template $NEWDIR
sed -i "s/DAY/$NEWDIR/" $NEWDIR/go.mod
