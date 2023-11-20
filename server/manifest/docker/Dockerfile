FROM alpine:3.18

RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone


ENV WORKDIR                 /app
WORKDIR $WORKDIR
ADD hack                                     $WORKDIR/hack/
ADD i18n                                     $WORKDIR/i18n/
ADD manifest/config                          $WORKDIR/manifest/config/
ADD resource                                 $WORKDIR/resource/
ADD ./temp/linux_amd64/grata                 $WORKDIR/grata
RUN chmod +xr $WORKDIR/grata


CMD ["./grata"]

