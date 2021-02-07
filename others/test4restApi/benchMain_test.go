package main

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//mockFile 代表测试样本数据
func send(t *testing.T, reqUrl, method string, reqBody []byte, headers http.Header) {
	Request := gorequest.New()
	Request.Header = headers
	var resp *http.Response
	var errs []error

	if method == "POST" {
		resp, _, errs = Request.Post(reqUrl).Type("json").Send(reqBody).EndBytes()
	} else {
		resp, _, errs = Request.Get(reqUrl).Type("json").EndBytes()
	}
	errStr := make([]string, 0)
	for _, err := range errs {
		assert.Nil(t, err)
		if err != nil {
			errStr = append(errStr, err.Error())
			return
		}
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func BenchmarkHelloWorld(b *testing.B) {
	b.ReportAllocs()
	t := new(testing.T)
	headers := map[string][]string{
		"Content-Type":  {"application/json"},
		"Authorization": {""},
	}
	var by []byte
	for i := 0; i < b.N; i++ {
		send(t, "/hello?id=1000", http.MethodGet, by, headers)
	}
}

//注意命名规范 Benchmark+首字母大写的方法名 参数固定
func BenchmarkAuthMail(b *testing.B) {
	params, _ := json.Marshal(map[string]interface{}{
		"email":    "5303221@gmail.com",
		"password": "123456",
	})
	headers := map[string][]string{
		"Content-Type":  {"application/json"},
		"Authorization": {""},
	}
	t := new(testing.T)
	for i := 0; i < b.N; i++ {
		send(t, "/auth/mail", http.MethodPost, params, headers)
	}
}
