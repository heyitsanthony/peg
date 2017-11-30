#!/bin/bash

for a in $@; do
	dat=`(time ./golang 2>/dev/null <$a | wc -l) 2>&1 | xargs echo | awk ' { print $1" "$3 } '`
	echo $a $dat
done
