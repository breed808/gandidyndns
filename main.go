package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	gandi "github.com/breed808/gandiapi"
)

func getPublicIP() (string, error) {
	ipClient := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://ifconfig.me/ip", nil)
	if err != nil {
		return "", err
	}

	resp, err := ipClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	fqdn := os.Args[1]

	publicIP, err := getPublicIP()
	if err != nil {
		fmt.Printf("Error acquiring public ip address: %v", err)
		os.Exit(1)
	}

	client := gandi.NewClient(os.Getenv("GANDI_API_KEY"), false, false)

	for _, name := range os.Args[2:] {
		record := gandi.Record{
			RrsetType:   "A",
			RrsetTTL:    3600,
			RrsetName:   name,
			RrsetHref:   "",
			RrsetValues: []string{publicIP},
		}

		_, err = client.UpdateRecord(record, fqdn)
		if err != nil {
			fmt.Printf("Failed request: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Updated record for %s.%s\n", record.RrsetName, fqdn)
	}
}
