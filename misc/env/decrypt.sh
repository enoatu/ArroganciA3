#!/usr/bin/env bash

read -s -p "Enter Password: " PASS
openssl enc -d -aes-256-cbc -salt -k $PASS -in .env.encrypt -out .env
