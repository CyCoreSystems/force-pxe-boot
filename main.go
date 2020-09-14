package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"

	gofish "github.com/stmcginnis/gofish/school"
)

var host string
var username string
var password string

func init() {
	flag.StringVar(&host, "-h", "", "Hostname or IP address of BMC")
	flag.StringVar(&username, "-u", "ADMIN", "Username for BMC authentication")
	flag.StringVar(&password, "-p", "ADMIN", "Password for BMC authentication")
}

func httpClientForSelfSigned() (client *http.Client, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{
		Transport: tr,
	}

	return client, err
}

func main() {
	flag.Parse()

	if host == "" {
		log.Fatalln("hostname is required")
	}

	hc, err := httpClientForSelfSigned()
	if err != nil {
		log.Fatalln("failed to create http client:", err)
	}

	ac, err := gofish.APIClient(fmt.Sprintf("https://%s", host), hc)
	if err != nil {
		log.Fatalln("failed to create API client:", err)
	}

	svc, err := gofish.ServiceRoot(ac)
	if err != nil {
		log.Fatalln("failed to attach ServiceRoot to API Client:", err)
	}

	sess, err := svc.CreateSession(username, password)
	if err != nil {
		log.Fatalln("failed to create authenticated session:", err)
	}
	defer svc.DeleteSession(sess.Session)

	ac.Token = sess.Token

	list, err := svc.Systems()
	if err != nil {
		log.Fatalln("failed to get Systems interface:", err)
	}
	if len(list) < 1 {
		log.Fatalln("no systems found")
	}

	list[0].FIXME - NEED - ACTION

}
