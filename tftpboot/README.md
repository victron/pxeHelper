# pxeHelper

Dynamically generate `kickstart` based on source ip.  
For pxe boot, dhcp server providing ip for new server. This IP based on mac address. 
Then normally provided `boot.cfg` via tftp, where present link to ks file. 
`pxeHelper` generate `ks` file from template looking into source IP. 
Specific settings located in `.csv` file.
Additionally `pxeHelper` provide static url for image upload.

# state less image
Just prepared environment. All configuration doing during run.
All binary licated in this folder.

## Files to run container

need to put in project folder.
[example](example/docker_VMware)

- table.csv
- template_dnsmasq.conf
- templateKS.conf
- boot.cfg
- .iso image

## run
`docker run -d --privileged --net host -v `pwd`:/var/lib/tftpboot/pxeHelper.d victron/pxehelper:latest`

`--privileged --net host` - requirement of dnsmasq


# state full image (outdated, just as reference)
all configs, templates should be available at the moment to build image

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

## run
`docker run -d --privileged --net host victron/pxehelper:XXXX`
