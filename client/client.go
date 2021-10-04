package client

import (
	"encoding/base64"
	"net/http"
	"time"
)

const (
	APIPath   = "https://online.moysklad.ru/api/remap/1.2"
	LimitRows = 1000
)

type moySklad struct {
	credentials  string
	requestLimit chan int
	api          *http.Client
}

type MS interface {
	Product() Product
}

func MoySkladClient(login, password string) MS {
	var result MS
	encodeCreds := base64.StdEncoding.EncodeToString([]byte(login + ":" + password))

	result = &moySklad{
		credentials:  encodeCreds,
		requestLimit: make(chan int, 5),
		api: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       5,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		},
	}

	return result
}

func (ms *moySklad) Product() Product {
	return productsClient(ms)
}
