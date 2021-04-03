/* ------------ License ------------ */
// BSD 3-Clause License
//
// Copyright (c) 2021, Seongmin Kim
// All rights reserved.
/* --------------------------------- */

package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpGet(url string, header map[string]string, data url.Values) ([]byte, error) {
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(data.Encode()))
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("[HttpGet] HTTP GET Error", err, resp.StatusCode)
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpPOST(url string, header map[string]string, data url.Values) ([]byte, error) {
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//println(data.Encode())
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("[HttpPost] HTTP POST Error", err, resp.StatusCode)
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
