#!/usr/bin/env bash

<< COMMENTOUT
sudo apt update -y && sudo apt install -y git && rm -rf ./setup && git clone https://github.com/enoatu/setup && ./setup/debian_setup.sh && rm -rf ./ArroganciA3 && git clone https://github.com/enoatu/ArroganciA3 && ./ArroganciA3/setup/debian_setup.sh
COMMENTOUT

# debian
# 1. install docker
# ref: https://matsuand.github.io/docs.docker.jp.onthefly/engine/install/debian/
sudo apt update -y
sudo apt install -y \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg \
  lsb-release

curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

 echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt update -y
sudo apt install -y docker-ce docker-ce-cli containerd.io
# 2. start service
sudo systemctl enable docker
sudo systemctl restart docker
# 3. enable no-sudo
sudo usermod -a -G docker $USER
# 4. install docker-compose and enable no-sudo
sudo curl -L https://github.com/docker/compose/releases/download/1.29.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose
# 5. yarn last setup
sudo curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt install -y nodejs
sudo npm install -g yarn
cd ArroganciA3
touch .env # for docker-compose up
/usr/bin/yarn install
<< COMMENTOUT
exec $SHELL -l
COMMENTOUT
