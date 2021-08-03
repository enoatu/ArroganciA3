#!/usr/bin/env bash
<< COMMENTOUT
rm -rf ./setup && git clone https://github.com/enoatu/setup && ./setup/centos_setup.sh && rm -rf ./ArroganciA3 && git clone https://github.com/enoatu/ArroganciA3 && ./ArroganciA3/setup/mac_setup.sh
COMMENTOUT

# ready docker-compose
cd ArroganciA3
touch .env # for docker-compose up
yarn install
/bin/yarn install
