package net

import (
	"fmt"
	"testing"
)

func TestFakeResponseWriter(t *testing.T) {

	w := NewFakeResponseWriter()
	data := map[string]string{}

	data["API"] = "1.0"

	w.Header().Add("Content-Type", "application/json")
	//ServerJSON(w, data)
	ServerRAW(w, data)
	fmt.Printf("Header %v \t Body %s\n", w.Header(), w.Body())
}
