package hkclient

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/remind101/emp/Godeps/_workspace/src/github.com/remind101/empire/pkg/heroku"
)

type Clients struct {
	ApiURL string
	Client *heroku.Client
}

func New(nrc *NetRc, agent string) (*Clients, error) {
	userAgent := agent + " " + heroku.DefaultUserAgent
	ste := Clients{}

	ste.ApiURL = os.Getenv("EMPIRE_API_URL")
	if ste.ApiURL == "" {
		return nil, errors.New("EMPIRE_API_URL must be set")
	}

	disableSSLVerify := false

	apiURL, err := url.Parse(ste.ApiURL)
	if err != nil {
		return nil, err
	}

	user, pass, err := nrc.GetCreds(apiURL)
	if err != nil {
		return nil, err
	}

	debug := os.Getenv("HKDEBUG") != ""
	ste.Client = &heroku.Client{
		URL:       ste.ApiURL,
		Username:  user,
		Password:  pass,
		UserAgent: userAgent,
		Debug:     debug,
	}

	tr := &http.Transport{}
	ste.Client.HTTP = &http.Client{Transport: tr}

	if disableSSLVerify || os.Getenv("HEROKU_SSL_VERIFY") == "disable" {
		tr.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	ste.Client.AdditionalHeaders = http.Header{}
	for _, h := range strings.Split(os.Getenv("HKHEADER"), "\n") {
		if i := strings.Index(h, ":"); i >= 0 {
			ste.Client.AdditionalHeaders.Set(
				strings.TrimSpace(h[:i]),
				strings.TrimSpace(h[i+1:]),
			)
		}
	}

	return &ste, nil
}
