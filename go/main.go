package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

func main() {
	p := bluemonday.UGCPolicy()
	resty := resty.New()
	r := gin.Default()
	r.POST("/js", func(c *gin.Context) {

		c.String(http.StatusOK, "go[bluemonday]: %s\n", p.Sanitize(c.PostForm("js")))

		re, _ := httpClientJS(resty, c.PostForm("js"))
		c.String(http.StatusOK, "js[sanitize-html]: %s\n", re)

		re, _ = httpClientDOMpurify(resty, c.PostForm("js"))
		c.String(http.StatusOK, "js[dompurify]: %s\n", re)

		re, _ = httpClientPython(resty, c.PostForm("js"))
		c.String(http.StatusOK, "python[html_sanitizer]: %s\n", re)
		
		
	})
	r.Run("0.0.0.0:9090")

}

func httpClientJS(r *resty.Client, value string) (string, error) {
	res, errRes := r.R().SetFormData(map[string]string{
		"js": value,
	}).Post("http://node:3000")

	if errRes != nil {
		fmt.Println(errRes)
		return "", errRes
	}
	return res.String(), nil
}

func httpClientDOMpurify(r *resty.Client, value string) (string, error) {
	res, errRes := r.R().SetFormData(map[string]string{
		"js": value,
	}).Post("http://node:3000/dompurify")

	if errRes != nil {
		fmt.Println(errRes)
		return "", errRes
	}
	return res.String(), nil
}

func httpClientPython(r *resty.Client, value string) (string, error) {
	res, errRes := r.R().SetFormData(map[string]string{
		"js": value,
	}).Post("http://python:8181")

	if errRes != nil {
		fmt.Println(errRes)
		return "", errRes
	}
	fmt.Println(res.RawResponse.Body)
	return res.String(), nil
}
