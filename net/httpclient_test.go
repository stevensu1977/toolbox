package net

import (
	"net/http"
	"net/url"
	"testing"
)

type EchoModel struct {
	Method string
	Header http.Header
	Body   string
}

func TestBuildRawQuery(t *testing.T) {
	p := make(map[string]string)
	p["token"] = "123123"
	p["appID"] = "abcd"
	p["appName"] = "汉字,http://demo"
	baseURL := "https://api.github.com"
	BuildRawQuery(baseURL, p)

}

func TestClientUnmarshal(t *testing.T) {
	r := &map[string]interface{}{}
	GetUnmarshal("http://localhost:9999", r)
	t.Logf("%v", r)
}

func TestHTTPSimpleClient(t *testing.T) {
	client := HTTPSimpleClient{
		Headers: map[string]string{
			"Content-Type":      "json",
			"X-Identity-domain": "apacsales",
		},
	}
	s, err := client.Get("http://localhost:9999")

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d %s", s.StatusCode, string(s.Body))

}

func TestHTTPSimpleClientPostForm(t *testing.T) {
	client := &HTTPSimpleClient{}
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	r, err := client.PostFormData("http://localhost:9999", v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(r.Body))

}

func TestHTTPFetch(t *testing.T) {
	Fetch("http://www.unixstickers.com/image/cache/data/stickers/golang/golang.sh-600x600.png", "../upload")
	Fetch("http://dl.google.com/chrome/install/696.57/chrome_installer.exe", "../upload")
}
