package main

import (
	"flag"
)

var imageFolder, pathKs, imageUrl, matchKey, csvFile *string
var templateKS, listen_adr_port, templateDnsmasq *string
var dnsmasqOnly *bool

func init() {
	imageFolder = flag.String("image", "../mnt", "path to folder with image")
	pathKs = flag.String("ks", "ks", "url to ks (kick start) file")
	imageUrl = flag.String("imageurl", "mnt", "url to image, used in boot.cfg in prefix keyword")
	matchKey = flag.String("key", "ip", "column name to do matching in csv file")
	csvFile = flag.String("csv", "./table.csv", "path to csv file")
	templateKS = flag.String("kstpl", "./templateKS.conf", "path to KS template")
	templateDnsmasq = flag.String("dnstpl", "./template_dnsmasq.conf", "path to dnsmasq template")
	listen_adr_port = flag.String("adrPort", ":80", "listen adress port. Example: localhost:80")
	dnsmasqOnly = flag.Bool("dnsonly", false, "generate dnsmasq config and exit")
	flag.Parse()
}
