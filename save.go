package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	sf "github.com/varuuntiwari/save-site/sf"
)

var (
	// depth int64
	site string
	domain string
)

func validSite() bool {
	if !strings.HasPrefix(site, "http://") && !strings.HasPrefix(site, "https://") {
		conn, err := http.Head("http://" + site)
		if err == nil && conn.StatusCode == 200 {
			domain = site
			site = "http://" + site
			return true
		}
		conn, err = http.Head("https://" + site)
		if err == nil && conn.StatusCode == 200 {
			domain = site
			site = "https://" + site
			return true
		}
		return false
	} else {
		_, err := http.Head(site)
		if strings.HasPrefix(site, "http://") {
			domain = strings.ReplaceAll(site, "http://", "")
		} else {
			domain = strings.ReplaceAll(site, "https://", "")
		}
		return err == nil
	}
}

func main() {
	fmt.Print("+----------------------------------+\n")
	fmt.Print("| save-site by varuuntiwari@GitHub |\n")
	fmt.Print("+----------------------------------+\n\n")
	flag.StringVar(&site, "site", "", "domain name to save")
	flag.Parse()
	if site == "" {
		panic("[-] site not specified")
	}

	if validSite() {
		fmt.Printf("[+] saving %v\n", domain)
	} else {
		fmt.Printf("[-] cannot reach %s\n", domain)
	}

	if sf.SaveSite(site, domain) {
		fmt.Println("[+] save successful")
	} else {
		fmt.Println("[-] some error occurred, try again")
	}	
}