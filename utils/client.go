package utils

import (
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func PostObject(
	url string,
	in interface{},
	out interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	return updateObject(url, "POST", in, out, ctx)
}

func PutObject(
	url string,
	in interface{},
	out interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	return updateObject(url, "PUT", in, out, ctx)
}

func updateObject(
	url string,
	method string,
	in interface{},
	out interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	marshalledBytes, ok := Marshal(in, ctx)
	if !ok {
		return false
	}

	response, ok := updateJson(url, method, string(marshalledBytes), ctx)

	if !ok {
		return false
	}

	if !Unmarshal(response, out, ctx) {
		return false
	}

	return true
}

func PostJson(
	url string,
	jsonString string,
	ctx *fasthttp.RequestCtx,
) (string, bool) {
	data, ok := updateJson(url, "POST", jsonString, ctx)

	if !ok {
		return "", false
	}

	return string(data), ok
}

func PutJson(
	url string,
	jsonString string,
	ctx *fasthttp.RequestCtx,
) (string, bool) {
	data, ok := updateJson(url, "PUT", jsonString, ctx)

	if !ok {
		return "", false
	}

	return string(data), ok
}

func updateJson(
	url string,
	method string,
	jsonString string,
	ctx *fasthttp.RequestCtx,
) ([]byte, bool) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")
	req.SetBodyString(jsonString)

	req.SetRequestURI(url)

	error := fasthttp.Do(req, resp)

	if error != nil {
		log.Printf("Unable to %s to %s\n", method, url)
		log.Print(error)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return nil, false
	}

	bodyBytes := resp.Body()

	return bodyBytes, true
}

func GetObject(
	url string,
	out interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	response, ok := Get(url, ctx)

	if !ok {
		return false
	}

	if !Unmarshal(response, out, ctx) {
		return false
	}

	return true
}

func Get(
	url string,
	ctx *fasthttp.RequestCtx,
) ([]byte, bool) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)

	error := fasthttp.Do(req, resp)

	if error != nil {
		log.Printf("Unable to GET from %s\n", url)
		log.Print(error)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return nil, false
	}

	bodyBytes := resp.Body()

	return bodyBytes, true
}
