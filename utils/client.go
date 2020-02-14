package utils

import "github.com/valyala/fasthttp"

func PutJson(url string, jsonString string) string {
	return updateJson(url, "PUT", jsonString)
}

func PostJson(url string, jsonString string) string {
	return updateJson(url, "POST", jsonString)
}

func updateJson(url string, method string, jsonString string) string {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")
	req.SetBodyString(jsonString)

	req.SetRequestURI(url)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()

	return string(bodyBytes)
	// User-Agent: fasthttp
	// Body:
}

func Get(url string) string {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(url)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()

	return string(bodyBytes)
	// User-Agent: fasthttp
	// Body:
}
