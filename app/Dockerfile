FROM golang:1.16.3-buster
MAINTAINER enoatu

RUN apt-get update
RUN apt-get install -y \
    git \
    unzip \
    tree \
# サイズをへらす
    && apt-get clean

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/beego/bee/v2@v2.0.2
RUN go get -u github.com/beego/beego/v2@v2.0.1

COPY ./arrogancia /go/src/arrogancia
WORKDIR /go/src/arrogancia

# 初回時
#
# routerなしのエラーが出る
# bee generate controller base

# 依存ライブラリをダウンロードする
RUN go mod download
# run時にgithub.com/astaxie/beegoで'missing go.sum entry'が出るので予め不要パッケージ削除
RUN go mod tidy

# タイムゾーン設定
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && dpkg-reconfigure -f noninteractive tzdata

# bee server起動
ENTRYPOINT ["/go/bin/bee", "run"]
CMD ["-downdoc=true", "-gendoc=true"]
