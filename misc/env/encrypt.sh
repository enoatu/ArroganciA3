#!/usr/bin/env bash

# not libressl, require openssl
read -s -p "Enter Password: " PASS1
echo
read -s -p "ReEnter Password: " PASS2
echo
[ $PASS1 != $PASS2 ] && echo "mismatched password" && exit 1
docker-compose exec -T app bash -c "cd /go/src/arrogancia && openssl enc -e -aes-256-cbc -md sha256 -iter 10000 -salt -k $PASS1 -in .env -out .env.encrypt"
