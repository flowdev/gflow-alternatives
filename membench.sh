#!/usr/bin/bash

outfile="membench.out"
rm -f $outfile
for alt in classic lambdas oopmix rets ; do
	go build "$alt/cmd/${alt}mem/${alt}mem.go"
	/usr/bin/time -f "$alt	%MKb" ./${alt}mem &>> $outfile
done
