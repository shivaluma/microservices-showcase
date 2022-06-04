package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestApplicationRunnable(t *testing.T) {
	go main()

	time.Sleep(time.Second * 1)

	res, err := http.Get("http://0.0.0.0:1001")

	if err != nil {
		t.Errorf("Error when calling server: %v", err)
	}

	defer res.Body.Close()
	out, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, "Hello, World From Fiber FastHTTP!", string(out), "they should be equal")
}
