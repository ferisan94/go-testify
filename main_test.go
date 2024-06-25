package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "ожидался статус-код 200")
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	require.Len(t, list, totalCount, "ожидалось %d кафе", totalCount)
}

func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "ожидался статус-код 200")
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "тело ответа не должно быть пустым")
	list := strings.Split(body, ",")
	assert.Len(t, list, 2, "ожидалось 2 кафе")
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=unknown", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "ожидался статус-код 400")
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body, "ожидалось сообщение 'wrong city value'")
}
