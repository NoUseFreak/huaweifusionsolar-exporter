package huaweifusionsolar

import (
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type HuaweiInverter struct {
	baseAPI  string
	username string
	password string

	token string

	client *http.Client
}

func New(username string, password string, api_endpoint string) (HuaweiInverter, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	return HuaweiInverter{
		baseAPI:  api_endpoint,
		username: username,
		password: password,

		client: &http.Client{
			Jar: jar,
		},
	}, err
}
