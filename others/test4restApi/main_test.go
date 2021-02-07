package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 封装的通用请求方法
func newRequest(method, path string, header map[string]string, body io.Reader) *httptest.ResponseRecorder {
	// 直接复用定义好 gin 路由实例
	router := SetupRouter()
	req, _ := http.NewRequest(method, path, body)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	router.ServeHTTP(w, req)
	defer w.Result().Body.Close()
	return w
}

// 模拟GET请求
func TestHelloWorld(t *testing.T) {
	w := newRequest(http.MethodGet, "/hello?id=1000", map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "",
	}, nil)

	assert.Equal(t, http.StatusOK, w.Code)

	var response Response
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, response.Code, 0)
	if data, ok := response.Data.(map[string]string); ok {
		assert.Equal(t, 1000, data["id"])
	}
	t.Log("响应内容", response.Data)

}

func TestAuthMail(t *testing.T) {
	params, err := json.Marshal(map[string]interface{}{
		"email":    "5303221@gmail.com",
		"password": "123456",
	})
	assert.Nil(t, err)
	w := newRequest(
		http.MethodPost,
		"/auth/mail",
		map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "",
		},
		bytes.NewBuffer(params))

	assert.Equal(t, http.StatusOK, w.Code)

	var response Response
	err = json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, response.Code, 0)

	if data, ok := response.Data.(map[string]string); ok {
		assert.Equal(t, "5303221@gmail.com", data["email"])
	}
	t.Log("响应内容", response.Data)
}