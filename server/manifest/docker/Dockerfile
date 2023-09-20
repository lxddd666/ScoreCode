FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
ADD hack                $WORKDIR/hack/
ADD manifest/config                $WORKDIR/manifest/config/
ADD resource                $WORKDIR/resource/
ADD ./temp/linux_amd64/grata $WORKDIR/grata
ADD ./manifest/docker/entrypoint.sh $WORKDIR/entrypoint.sh
RUN chmod +x $WORKDIR/grata
RUN chmod +x $WORKDIR/entrypoint.sh


###############################################################################
#                                   START
###############################################################################

WORKDIR $WORKDIR
CMD ./entrypoint.sh

