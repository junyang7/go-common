#!/usr/bin/env sh

ROOT=$(cd "$(dirname "$0")";pwd)

for item in `ls ${ROOT}`
do
  if [ -d ${item} ]
  then
    cd ${ROOT}/${item}
    go test
    cd ${ROOT}
  fi
done
