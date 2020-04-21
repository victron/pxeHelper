FROM alpine

WORKDIR /var/lib/tftpboot
COPY ./ /var/lib/tftpboot/
RUN apk add dnsmasq && \
    apk add libc6-compat && \
    apk add xorriso && \
    apk add tree && \
    wget https://github.com/victron/pxeHelper/releases/download/2020-04-18/tftpboot.tgz && \
    wget https://github.com/victron/pxeHelper/releases/download/2020-04-18/tftpboot.tgz.md5 && \
    md5sum -c tftpboot.tgz.md5 && \
    tar xvzf tftpboot.tgz && \
    ./pxeHelper -dnsonly && \
    osirrox -indev *.iso -extract / mnt && \
    rm *.iso

CMD dnsmasq --conf-file=./dnsmasq.conf --dhcp-broadcast && ./pxeHelper -image mnt

# NOTE: 
# docker run -d --privileged --net host victron/pxehelper:XXXX
