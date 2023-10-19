FROM alpine:3.18

###############################################################################
#                                INSTALLATION
###############################################################################
ENV WORKDIR                 /app
WORKDIR $WORKDIR
ADD hack                                     $WORKDIR/hack/
ADD manifest/config/casbin.conf              $WORKDIR/manifest/config/casbin.conf
ADD manifest/config/config.yaml              $WORKDIR/config.yaml
ADD resource                                 $WORKDIR/resource/
ADD ./temp/linux_amd64/grata                 $WORKDIR/grata
ADD ./manifest/docker/entrypoint.sh          $WORKDIR/entrypoint.sh
RUN chmod +x $WORKDIR/grata
RUN chmod +x $WORKDIR/entrypoint.sh


###############################################################################
#                                   START
###############################################################################


CMD ./entrypoint.sh

