package services

import (
	"fmt"
	"os"
	"time"

	"resty.dev/v3"
)

var client *resty.Client

func init() {
	apikey := os.Getenv("VIRUSTOTAL_APIKEY")
	client = resty.New()
	client.SetBaseURL("https://www.virustotal.com/api/v3")
	client.SetHeader("x-apikey", apikey)
	client.SetTimeout(10 * time.Second)
}

func testURL(url string) (bool, error) {
	var scanResult struct {
		Data struct {
			Id string `json:"id"`
		} `json:"data"`
	}

	resp, err := client.R().SetFormData(map[string]string{"url": url}).SetResult(&scanResult).Post("/urls")
	if err != nil { return false, fmt.Errorf("Failed to request scan: %w", err) }
	if resp.IsError() { return false, fmt.Errorf("Failed to request scan: %s (body: %s)", resp.Status(), resp.String()) }

	var report struct {
		Data struct {
			Attributes struct {
				Stats struct {
					Undetected int `json:"undetected"`
					Harmless int `json:"harmless"`
					Malicious int `json:"malicious"`
					Suspicious int `json:"suspicious"`
				} `json:"stats"`
			} `json:"attributes"`
		} `json:"data"`
	}

	resp, err = client.R().SetResult(&report).Get("/analyses/" + scanResult.Data.Id)
	if err != nil { return false, err }
	if resp.IsError() { return false, fmt.Errorf("report request failed: %s (body: %s)", resp.Status(), resp.String()) }

	s := report.Data.Attributes.Stats

	if s.Malicious > 0 || s.Suspicious > 0 { return false, fmt.Errorf("Malicious URL detected") }

	return true, nil
}
