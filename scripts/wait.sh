#!/usr/bin/env bash

for ((i=1; i<=$#; i++))
do
  if [ ${!i} == '--' ]
  then
    cmd=""
    for c in ${@:($i+1)}
    do
      cmd=$cmd' '$c
    done
    $cmd
    break
  else
#    wget -qO- https://ghproxy.com/https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for | sh -s -- 127.0.0.1:${!i} -- echo success
    bash /root/wait-for 127.0.0.1:${!i} -- echo success
  fi
done