package utils

import (
	"bytes"
	"github.com/klauspost/compress/gzip"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	// https://blog.klauspost.com/gzip-performance-for-go-webservers/
	gzippers = sync.Pool{New: func() interface{} {
		return gzip.NewWriter(nil)
	}}
	gunzippers = sync.Pool{New: func() interface{} {
		reader, _ := gzip.NewReader(nil)

		return reader
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
	defer gz.Close()

	if _, err := gz.Write(dataBytes); err != nil {
		log.Print("Unable to gzip %s", ctx.UserValue("recordType"))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return buf, false
	}

	return buf, true
}

func UnzipToNewZipper(
	compressedData []byte,
) (bytes.Buffer, *gzip.Writer, bool) {
	var buf bytes.Buffer

	// Get a Writer from the Pool
	// https://blog.klauspost.com/gzip-performance-for-go-webservers/
	gz := gzippers.Get().(*gzip.Writer)
	gz.Reset(&buf)

	if !UnzipToZipper(compressedData, gz) {
		gz.Close()
		gzippers.Put(gz)

		return buf, nil, false
	}

	return buf, gz, true
}

func UnzipToZipper(
	compressedData []byte,
	gz *gzip.Writer,
) bool {
	if compressedData != nil {
		compressedDataReader := bytes.NewReader(compressedData)
		gunz := gunzippers.Get().(*gzip.Reader)
		gunz.Reset(compressedDataReader)

		defer gunzippers.Put(gunz)
		defer gunz.Close()

		if _, err := io.Copy(gz, gunz); err != nil {
			log.Print(err)

			return false
		}
	}

	return true
}

func CloseZipper(
	gz *gzip.Writer,
) {
	gz.Close()
	gzippers.Put(gz)
}
