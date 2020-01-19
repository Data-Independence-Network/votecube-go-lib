package utils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func Unmarshal(
	data []byte,
	v interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	if err := json.Unmarshal(data, v); err != nil {
		log.Printf("Unable to unmarshal %s", ctx.UserValue("recordType"))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return false
	}

	return true
}

func Marshal(
	v interface{},
	ctx *fasthttp.RequestCtx,
) ([]byte, bool) {
	dataBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("Unable to marshal %s", ctx.UserValue("recordType"))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return nil, false
	}

	return dataBytes, true
}
