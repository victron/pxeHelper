# VMware

## confert disk from virtual box to VMware
`vmkfstools -i VB_disk.vmdk  -d thin vmware_disk.vmdk`
[source](https://www.vionblog.com/migrate-virtualbox-vmware-esxi-6-5/)
[ddb.adapterType = "id" to ddb.adapterType = "lsilogic"](https://kb.vmware.com/s/article/1016192?docid=1028042)