package utils

import (
	"bytes"
	"github.com/klauspost/compress/gzip"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"sync"
)

var (
	gzippers = sync.Pool{New: func() interface{} {
		return gzip.NewWriter(nil)
	}}
)

func MarshalZip(
	v interface{},
	ctx *fasthttp.RequestCtx,
) (bytes.Buffer, bool) {
	dataBytes, ok := Marshal(v, ctx)
	if !ok {
		return bytes.Buffer{}, false
	}

	return Zip(dataBytes, ctx)
}

func Zip(
	dataBytes []byte,
	ctx *fasthttp.RequestCtx,
) (bytes.Buffer, bool) {
	var buf bytes.Buffer
	// https://blog.klauspost.com/gzip-performance-for-go-webservers/
	gz := gzippers.Get().(*gzip.Writer)
	gz.Reset(&buf)

	defer gzippers.Put(gz)

	if _, err := gz.Write(dataBytes); err != nil {
		log.Print("Unable to gzip %s", ctx.UserValue("recordType"))
		log.Print(err)
		gz.Close()
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return buf, false
	}
	gz.Close()

	return buf, true
}
