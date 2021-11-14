#!/bin/bash -eu

make_list=$(find . -type f -name Makefile)
for m in $make_list; do
    d=$(dirname $m)
    pushd $d
    make all
    popd
    echo
done
