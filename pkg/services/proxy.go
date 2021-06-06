package services

import (
	"net/http"
	"net/url"
)

type ProxyService interface {
	New() func(*http.Request) (*url.URL, error)
}
