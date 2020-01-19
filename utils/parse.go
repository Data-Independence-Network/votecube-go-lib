package utils

import (
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"strconv"
)

func ParseInt64Param(
	paramName string,
	ctx *fasthttp.RequestCtx,
) (int64, bool) {
	number, parseError := strconv.ParseInt(ctx.UserValue(paramName).(string), 10, 64)
	if parseError != nil {
		log.Printf("Processing %s - Invalid %s: %s", ctx.UserValue("recordType"), paramName, ctx.UserValue(paramName))
		log.Print(parseError)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return 0, false
	}

	return number, true
}

func ParseUint64Param(
	paramName string,
	ctx *fasthttp.RequestCtx,
) (uint64, bool) {
	number, parseError := strconv.ParseUint(ctx.UserValue(paramName).(string), 10, 64)
	if parseError != nil {
		log.Printf("Processing %s - Invalid %s: %s", ctx.UserValue("recordType"), paramName, ctx.UserValue(paramName))
		log.Print(parseError)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return 0, false
	}

	return number, true
}
