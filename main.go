package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Table []map[string]string

// var mu sync.Mutex
var table Table

func handlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pxeHelper is UP... \n you requeted URL.Path = %q\n", r.URL.Path)
}

func handlerKS(w http.ResponseWriter, r *http.Request) {
	// mu.Lock()
	ip_remote := strings.Split(r.RemoteAddr, ":")[0]
	data, err := table.searchHost(ip_remote, *matchKey)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	tmpl, err := template.ParseFiles(*templateKS)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	tmpl.Execute(w, data)
	// mu.Unlock()
}

// parsing csv during startup
func csvReader() []map[string]string {
	file, e := os.Open(*csvFile)
	defer file.Close()

	if e != nil {
		log.Fatal(e)
	}

	reader := csv.NewReader(file)
	records, e := reader.ReadAll()
	if e != nil {
		log.Fatal(e)
	}

	csv_len := len(records)
	if csv_len < 2 {
		log.Fatal("in .csv should be min 2 lines")
	}

	var table Table
	header := records[0]

	for i, row := range records[1:] {
		table = append(table, make(map[string]string))
		for j, cell := range row {
			table[i][header[j]] = cell
		}
	}
	return table
}

// ip: ip address to search in table
// key: column name
func (table Table) searchHost(ip string, key string) (map[string]string, error) {
	for _, row := range table {
		if row[key] == ip {
			return row, nil
		}
	}
	return nil, errors.New("in .csv missing {key:val} \"ip\":" + ip)
}

func main() {
	table = csvReader()
	genDnsmasqConf(*templateDnsmasq)
	if *dnsmasqOnly {
		os.Exit(0)
	}
	fileServer := http.FileServer(http.Dir(*imageFolder))
	http.HandleFunc("/", handlerDefault)
	http.HandleFunc("/"+*pathKs, handlerKS)
	http.Handle("/"+*imageUrl+"/", http.StripPrefix("/"+*imageUrl+"/", fileServer))
	log.Println("Listening on:", *listen_adr_port, "...")
	log.Fatal(http.ListenAndServe(*listen_adr_port, nil))
}
