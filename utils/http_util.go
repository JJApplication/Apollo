/*
   Create: 2024/1/28
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import (
	"bytes"
	"io"
	"net/http"
)

// HttpGet 发送get请求
func HttpGet(url string, headers map[string]string) ([]byte, error) {
	rest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		rest.Header.Set(k, v)
	}
	cli := new(http.Client)
	res, err := cli.Do(rest)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
		return io.ReadAll(res.Body)
	}

	return nil, err
}

// HttpPost 发送post请求 rest请求
func HttpPost(url string, body []byte) ([]byte, error) {
	data := bytes.NewReader(body)
	rest, err := http.NewRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}
	rest.Header.Set("Content-Type", "application/json")
	rest.Header.Set("Accept", "application/json")
	cli := new(http.Client)
	res, err := cli.Do(rest)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
		return io.ReadAll(res.Body)
	}

	return nil, err
}
