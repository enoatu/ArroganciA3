#!/usr/bin/env bash

read -s -p "Enter Password: " PASS
openssl enc -e -aes-256-cbc -salt -k $PASS -in .env -out .env.encrypt
