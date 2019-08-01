package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func Test_Hello_ReturnsStatus200(t *testing.T) {
	hostIP := os.Getenv("ADDRESS")
	address := fmt.Sprintf("http://%s/demo/hello", hostIP)
	resp, err := http.Get(address)
	if err != nil {
		t.Error(err.Error())
	} else if resp == nil {
		t.Error("Got no response")
	} else if resp.StatusCode != 200 {
		t.Errorf("Got response %d", resp.StatusCode)
	}
}
