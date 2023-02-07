#!/usr/bin/env bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

build_func() {
  echo 'Building '$1'-server...'
  docker build -f docker/$1-server/Dockerfile -t vgalaxy/$1-server . > /dev/null 2>& 1
  if [ $? == 0 ]
  then
    echo -e "${GREEN}Success${NC}"
  else
    echo -e "${RED}failed${NC}"
  fi
}

servers=('api' 'user' 'comment' 'publish' 'feed' 'relation')

for server in ${servers[*]}
do
  build_func $server
done