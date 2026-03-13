package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCafeWhenOk(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	requests := []string{
		"/cafe?count=2&city=moscow",
		"/cafe?city=tula",
		"/cafe?city=moscow&search=ложка",
	}

	for _, v := range requests {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", v, nil)

		handler.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusOK)
		fmt.Println(response.Body.String())
	}
}

func TestCafeWhenNotOk(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	tableTest := []struct {
		url     string
		code    int
		message string
	}{
		{url: "/cafe", code: http.StatusBadRequest, message: "unknown city"},
		{url: "/cafe?city=omsk", code: http.StatusBadRequest, message: "unknown city"},
		{url: "/cafe?city=tula&count=na", code: http.StatusBadRequest, message: "incorrect count"},
	}

	for _, tt := range tableTest {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", tt.url, nil)

		handler.ServeHTTP(response, request)

		assert.Equal(t, response.Code, tt.code)
		assert.Equal(t, strings.TrimSpace(response.Body.String()), tt.message)
	}
}
