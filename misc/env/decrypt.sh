#!/usr/bin/env bash

# not libressl, require openssl
read -s -p "Enter Password: " PASS
cd app/arrogancia
docker-compose exec -T app bash -c "openssl enc -d -aes-256-cbc -md sha256 -iter 10000 -salt -k $PASS -in .env.encrypt -out .env"
