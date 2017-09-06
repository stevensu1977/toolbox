package net

import "testing"

func TestHTTPSimpleREST(t *testing.T) {
	client := &HTTPSimpleREST{}
	err := client.Get("http://localhost:9999", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := &map[string]interface{}{}

	err = client.Get("http://localhost:9999", resp)
	t.Log(resp)

	data := map[string]string{}
	data["email"] = "demo01@demo.com"

	err = client.Post("http://localhost:9999", data, resp)
	t.Log(resp)

	client.Headers = map[string]string{}
	client.Headers["X-Identity-Domain"] = "cncsmtrail3578"
	err = client.Put("http://localhost:9999", data, resp)
	t.Log(resp)

	client2 := NewHTTPSimpleREST()
	client2.Headers["X-Identity-Domain"] = "cncsmtrail3578"
	err = client2.Put("http://localhost:9999", data, resp)
	t.Log(resp)

}
