package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w, err
}

func TestApi(t *testing.T) {
	assert := assert.New(t)
	r := gin.Default()
	serve(r)

	t.Run("index.html", func(t *testing.T) {
		res, err := performRequest(r, "GET", "/")
		assert.Equal(err, nil, "should be equal")

		body, _ := ioutil.ReadAll(res.Body)

		html, err := ioutil.ReadFile("./public/index.html")
		assert.Equal(err, nil, "should be equal")

		assert.Equal(res.Result().StatusCode, 200, "should be equal")
		assert.Equal(string(body), string(html), "should be equal")
	})

	t.Run("404 error", func(t *testing.T) {
		res, err := performRequest(r, "GET", "/notfound")

		body, _ := ioutil.ReadAll(res.Body)

		assert.Equal(err, nil, "should be equal")
		assert.Equal(res.Result().StatusCode, 404, "should be equal")
		assert.Equal(string(body), `{"message":"Resource not found","success":false}`, "should be equal")
	})
}
