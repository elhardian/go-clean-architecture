package httpClient

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/elhardian/go-clean-architecture/libs/config"
	"github.com/rs/zerolog/log"
)

type Http interface {
	Connect()
	CallURL(method, url string, header map[string]string, rawData []byte) ([]byte, error)
}

type Options struct {
	timeout int
	http    *http.Client
}

func NewHttp(cfg *config.Config) Http {
	opt := new(Options)
	opt.timeout = cfg.HttpTimeout
	return opt
}

func (o *Options) Connect() {
	httpClient := &http.Client{
		Timeout: time.Duration(o.timeout) * time.Second,
	}

	o.http = httpClient
}

func (o *Options) CallURL(method, url string, header map[string]string, rawData []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(rawData))
	if err != nil {
		log.Error().Err(err).Msg("Failed To Prepare Request Client HTTP")
		return nil, err
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	res, err := o.http.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Failed To Request Client HTTP")
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed To Read Result Client HTTP")
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Error().Err(err).Msg("Error Status Code Not 200")
		return body, errors.New("code not 200")
	}

	return body, nil
}
