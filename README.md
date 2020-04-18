# pxeHelper

Dynamically generate `kickstart` based on source ip. 
For pxe boot, dhcp server providing ip for new server. This IP based on mac address. 
Then normally provided `boot.cfg` via tftp, where present link to ks file. 
`pxeHelper` generate `ks` file from template looking into source IP. 
Specific settings located in `.csv` file.
Additionally `pxeHelper` provide static url for image upload.

## Depends
- dnsmasq - as dhcp and tftp server (or any other)
- libc6-compat - for alpine )))


## Files
Required files
```
.
├── pxeHelper
├── table.csv
├── templateKS.cfg
└── template_dnsmasq.conf
```
- `pxeHelper` - on alpine need `libc6-compat` and `x` flag
- `table.csv` - info about individual settings for every new host
- `template_dnsmasq.conf` - template for `dnsmas`
Additionally in every start `pxeHelper` generates `dnsmasq.conf` from template `dnsmasq.conf`. You can: 
- replace `/etc/dnsmasq.conf` by it and restart `dnsmasq`
- or put in `/etc/dnsmasq.d/` (if your distro support it)
- or put line `conf-file=path_to_generated_config` in `/etc/dnsmasq.conf` [details](https://www.linux.com/topic/networking/advanced-dnsmasq-tips-and-tricks/)


## Quick start
```
# run in background in any dir
$ nohup ./pxeHelper -adrPort :8000 -csv test.csv -ksTpl ks.cfg & echo $! > pidfile                                                  
[19:07:06] 2020/04/17 19:08:05 Listening on: :8000 ...
.....
# to kill
kill `cat pidfile`
```

## Implementation example (Legasy BIOS only **not UEFI**)
based on `alpine` linux

### packages
`apk add dnsmasq`
[some info about apk](example/INFO/apk.md)

### [prepare network configuration](./example/interfaces)

### Files summary structure
```
/var/lib/tftpboot# tree
.
├── ESXI
│   ├── VMware-VMvisor-Installer-6.0.0.update02-3620759.x86_64.iso
│   ├── boot_esxi.cfg
│   ├── mnt
│   └── pxeHelper
│       ├── dnsmasq.conf
│       ├── pxeHelper
│       ├── table.csv
│       ├── templateKS.conf
│       └── template_dnsmasq.conf
├── gpxelinux.0
├── mboot.c32
├── menu.c32
├── pxelinux.0
└── pxelinux.cfg
    └── default
```
- `dnsmasq.conf` generated based on [table.csv](./example/table.csv)
`./pxeHelper -dnsonly`
[generated `dnsmasq.conf`](./example/dnsmasq.conf)
      - configure dnsmasq to use it.
      ```
      # cat /etc/dnsmasq.conf
      conf-file=/var/lib/tftpboot/ESXI/pxeHelper/dnsmasq.conf
      ```
      *restart `dnsmasq`*
      ```
      rc-update add dnsmasq default
      servise dnsmasq start
      ```
- `ESXI` - folder for project (specific OS)
- `.iso` - original iso image
- `mnt` - empty folder for mount point
- binary files:
    - `gpxelinux.0` or `pxelinux.0` - Obtain from SYSLINUX version 3.86 (vmware recomendation)
    [syslinux-3.86.zip](https://mirrors.edge.kernel.org/pub/linux/utils/boot/syslinux/3.xx/syslinux-3.86.zip)  
    *NOTE: it's better to use `gpxelinux` instead `pxelinux`. This allow use http for files transfer. (tftp used only to send initial small files). With `tftp` can be observed `fail sending` in reason of network.*
    - `mboot.c32` (optional for menu `menu.c32`) copy from iso image
- folder `pxelinux.cfg`. During boot server looking into this folder on tftp server. Server trying find individual for him files based on mac address and so on. [Details](https://wiki.syslinux.org/wiki/index.php?title=PXELINUX). Also it very visible in `dnsmasq` logs. At the end server looking `default` file. We are providing same - `pxelinux.cfg/default` for any pxe host.
[pxelinux.cfg/default](./example/default)
```
# cat /var/lib/tftpboot/pxelinux.cfg/default
DEFAULT install
NOHALT 1
LABEL install
  KERNEL mboot.c32
  APPEND -c ESXI/boot_esxi.cfg
```
- `boot_esxi.cfg` - taken from original image. [boot_esxi.cfg](./example/boot_esxi.cfg)
      - Added `prefix` as path to unpacked image folder. 
      - Deleted `/` before every file name.
      - Added path to kick start file ks=http://192.168.15.1/ks
- depends on IP, pxeHelper generate different `ks` files
[table.csv](./example/table.csv)

### mounting iso
```
cd /var/lib/tftpboot/ESXI

modprobe loop
LOOP=`losetup -f`
losetup $LOOP VMware-VMvisor-Installer-6.0.0.update02-3620759.x86_64.iso
mount -t iso9660 -o ro $LOOP mnt/
....
# reminder how to umount )))
umount /mnt
losetup -d $LOOP
```

### start pxeHelper
```
./pxeHelper
```

## References

### PXE principles
[PXE Booting the ESXi Installer](example/INFO/vsphere-esxi-67-upgrade-guide.pdf)
[Create an Installer ISO Image with a Custom Installation or Upgrade Script](https://docs.vmware.com/en/VMware-vSphere/6.0/com.vmware.vsphere.upgrade.doc/GUID-C03EADEA-A192-4AB4-9B71-9256A9CB1F9C.html)

### ks file
[official - mandatory params](https://docs.vmware.com/en/VMware-vSphere/6.7/com.vmware.esxi.install.doc/GUID-C3F32E0F-297B-4B75-8B3E-C28BD08680C8.html)
[Automated installation with VMware ESXi 5.5/6.0/6.5 ](https://be-virtual.net/automated-installation-with-vmware-esxi-5-56-06-5/)
[Customizing ESXi installation with kickstart files and PXE boot ](https://rudimartinsen.com/2018/06/09/customizing-esxi-installation-with-kickstart-files-and-pxe-boot/) vlan in kernel params
[Automatically Install VMware ESXi 6.7 through PXE boot](https://xenappblog.com/2018/automatically-install-vmware-esxi-6-7/)
[multiple firstboot sections](https://www.altaro.com/vmware/scripted-deployment-esxi-part-1/)

### UEFI
[Syslinux boot loaders](https://wiki.syslinux.org/wiki/index.php?title=Config)

### Alpine
[Configure Networking](https://wiki.alpinelinux.org/wiki/Configure_Networking)
[Burning ISOs](https://wiki.alpinelinux.org/wiki/Burning_ISOs)
[How to enable and start services on Alpine Linux](https://www.cyberciti.biz/faq/how-to-enable-and-start-services-on-alpine-linux/)

### dnsmasq
[Configuring Dnsmasq to Support PXE Clients](https://docs.oracle.com/en/operating-systems/oracle-linux/7/install/ol7-install-pxe-dnsmasq.html)
[PXE](https://wiki.archlinux.org/index.php/PXE_(%D0%A0%D1%83%D1%81%D1%81%D0%BA%D0%B8%D0%B9))
[Advanced Dnsmasq Tips and Tricks](https://www.linux.com/topic/networking/advanced-dnsmasq-tips-and-tricks/) - config tree.

