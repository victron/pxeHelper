FROM alpine

WORKDIR /var/lib/tftpboot
COPY ./ /var/lib/tftpboot/
VOLUME /var/lib/tftpboot/pxeHelper.d
RUN apk add dnsmasq && \
    apk add libc6-compat && \
    apk add xorriso

CMD osirrox -indev pxeHelper.d/*.iso -extract / mnt && \ 
    ./pxeHelper -dnsonly -dnstpl pxeHelper.d/template_dnsmasq.conf -csv pxeHelper.d/table.csv && \
    dnsmasq --conf-file=./dnsmasq.conf --dhcp-broadcast && \
    ./pxeHelper -dnstpl pxeHelper.d/template_dnsmasq.conf -csv pxeHelper.d/table.csv -kstpl pxeHelper.d/templateKS.conf -image mnt
    
    

# NOTE: 
# docker run -it --privileged --net host -v `pwd`:/var/lib/tftpboot/pxeHelper.d victron/pxehelper:latest
