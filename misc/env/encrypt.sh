#!/usr/bin/env bash

# not libressl, require openssl
read -s -p "Enter Password: " PASS1
read -s -p "ReEnter Password: " PASS2
[$PASS1 != $PASS2] && echo "mismatched password" && exit 1
docker-compose exec -T app bash -c "cd /go/src/arrogancia && openssl enc -e -aes-256-cbc -md sha256 -iter 10000 -salt -k $PASS -in .env -out .env.encrypt"
