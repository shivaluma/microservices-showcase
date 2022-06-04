package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestApplicationRunnable(t *testing.T) {
	go main()
	res, err := http.Get("http://127.0.0.1:1001/")

	if err != nil {
		t.Errorf("Error when calling server: %v", err)
	}

	defer res.Body.Close()
	out, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, "Hello, World From Fiber FastHTTP!", string(out), "they should be equal")
}
