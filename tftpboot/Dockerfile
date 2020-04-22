FROM alpine

WORKDIR /var/lib/tftpboot
COPY ./ /var/lib/tftpboot/
VOLUME /var/lib/tftpboot/pxeHelper.d
RUN apk add dnsmasq && \
    apk add libc6-compat && \
    apk add xorriso

CMD osirrox -indev *.iso -extract / mnt && \ 
    ./pxeHelper -dnsonly -dnstpl pxeHelper.d/template_dnsmasq.conf && \
    dnsmasq --conf-file=./dnsmasq.conf --dhcp-broadcast && \
    ./pxeHelper -dnstpl pxeHelper.d/template_dnsmasq.conf -kstpl pxeHelper.d/templateKS.conf -image mnt
    
    

# NOTE: 
# docker run -d --privileged --net host -v ./:./pxeHelper.d victron/pxehelper:XXXX
