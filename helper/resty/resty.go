package resty

import "github.com/go-resty/resty/v2"

func NewResty() *resty.Client {
	return resty.New()
}
