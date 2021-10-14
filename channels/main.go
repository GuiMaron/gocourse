package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkSite (site string, channel chan string) {

	r, e := http.Get(site)
	m	 := ""

	if e != nil {
		m = fmt.Sprintf("%s failed (%s)\n", site, e)
		fmt.Print(m)
		channel <- site
		return
	}

	if r.StatusCode != http.StatusOK {
		m = fmt.Sprintf("%s responded with status code %d\n", site, r.StatusCode)
		fmt.Print(m)
		channel <- site
		return
	}

	m = fmt.Sprintf("%s is OK\n", site)
	fmt.Print(m)
	channel <- site

}

func main () {

	sites := []string {
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	channel := make(chan string)

	for _, site := range sites {
		go checkSite(site, channel)
	}

	for site := range channel  {

		go func (s string) {
			time.Sleep(333 * time.Millisecond)
			checkSite(s, channel)
		}(site)

	}

}
