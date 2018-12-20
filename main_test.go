package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000", nil)
	if err != nil {
		t.Fatalf("could not load index")
	}

	wri := httptest.NewRecorder()

	indexHandler(wri, req)

	res := wri.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %v", res.Status)

	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response because: %v", err)
	}
	msg := string(bytes.TrimSpace(body))

	fmt.Println(msg)

}
