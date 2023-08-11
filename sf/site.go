package sf

import (
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	dir string
)

func SaveSite(site, url string) bool {
	dir = "sites/" + strings.ReplaceAll(url, "/", "_")
	
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return false
	}

	resp, err := http.Get(site)
	if err != nil {
		return false
	}
	
	f, _ := os.Create(dir + "/index.html")
	defer f.Close()
	
	io.Copy(f, resp.Body)

	return true
}