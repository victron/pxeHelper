## Files to build image
need to put in this folder actual files
- table.csv
- template_dnsmasq.conf
- templateKS.conf
- boot.cfg
- .iso image

## binaries
all binaries in `tftpboot`
all actual configs and template in `docker` dir

## exec
`docker run -d --privileged --net host victron/pxehelper:XXXX`