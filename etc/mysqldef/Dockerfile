ARG BASE=alpine:3.12
FROM ${BASE} as builder

ENV VERSION 0.10.16

ADD https://github.com/k0kubun/sqldef/releases/download/v${VERSION}/mysqldef_linux_amd64.tar.gz /tmp/
RUN cd /tmp \
    && tar zxvf mysqldef_linux_amd64.tar.gz -C /bin \
    && rm -rf mysqldef_linux_amd64.tar.gz

ENTRYPOINT [ "mysqldef" ]
