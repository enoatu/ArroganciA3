FROM golang:1.16.3-buster
MAINTAINER enoatu

RUN apt-get update
RUN apt-get install -y \
    git \
    unzip \
    tree \
    cron \
    daemontools \
# サイズをへらす
    && apt-get clean

RUN go get -u github.com/go-sql-driver/mysql

COPY ./batch /go/src/batch
WORKDIR /go/src/batch

RUN go mod init

#RUN go mod download
## run時にgithub.com/astaxie/beegoで'missing go.sum entry'が出るので予め不要パッケージ削除
#RUN go mod tidy

# タイムゾーン設定
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && dpkg-reconfigure -f noninteractive tzdata
