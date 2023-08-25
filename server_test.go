package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"io"
	//"github.com/assertgo/assert" не работает нихуя
)

func TestHandlerRequestEasy(t *testing.T) {
	expected := "Hello, new request w.write"
	// TODO: how you check url?????
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	HandlerRequestEasy(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	
	if err != nil {
			t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
			t.Errorf("Expected Hello john but got %v", string(data))
	}
	if res.StatusCode != 200 {
		t.Error("Expected status code 200")
	}
}

func TestUsdHandlerRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	UsdHandlerRequest(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	t.Log(string(data[:]))
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if res.StatusCode != 200 {
		t.Error("Expected status code 200")
	}

}




