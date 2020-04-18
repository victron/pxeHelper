// generate config for dnsmasq

package main

import (
	"log"
	"os"
	"text/template"
)

const DNSMASQ_CONFIG_OUTPUT = "./dnsmasq.cfg"

func genDnsmasqConf(templateFile string) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	writer, err := os.Create(DNSMASQ_CONFIG_OUTPUT)
	defer writer.Close()
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	tmpl.Execute(writer, table)
}
