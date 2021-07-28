#!/usr/bin/env bash

<< COMMENTOUT
sudo yum update -y && sudo yum install -y git && rm -rf ./setup && git clone https://github.com/enoatu/setup && ./setup/centos_setup.sh && rm -rf ./ArroganciA3 && git clone https://github.com/enoatu/ArroganciA3 && ./ArroganciA3/setup/centos_setup.sh
COMMENTOUT

# amazon linux2
# 1. install docker
sudo amazon-linux-extras install docker
# 2. start service
sudo service docker start
# 3. enable no-sudo
sudo usermod -a -G docker ec2-user
# 4. install docker-compose and enable no-sudo
sudo curl -L https://github.com/docker/compose/releases/download/1.29.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose
# 5. yarn last setup
sudo curl -sL https://rpm.nodesource.com/setup_16.x | sudo bash -
sudo yum install -y nodejs
exsh
sudo npm install -g yarn
exsh
cd ArroganciA3
yarn install
