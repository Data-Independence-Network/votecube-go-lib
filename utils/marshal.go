package utils

import (
	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"reflect"
)

/*
func Unmarshal(
	data []byte,
	v interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal(data, v); err != nil {
		log.Printf("Unable to unmarshal %s\n", reflect.TypeOf(v))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return false
	}

	return true
}
*/

// http://jsoniter.com/migrate-from-go-std.html
func Unmarshal(
	data []byte,
	v interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	iter := jsoniter.ConfigFastest.BorrowIterator(data)
	defer jsoniter.ConfigFastest.ReturnIterator(iter)
	iter.ReadVal(&v)
	if iter.Error != nil {
		log.Printf("Unable to unmarshal %s\n", reflect.TypeOf(v))
		log.Print(iter.Error)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
	}

	return true
}

func MarshalWithPreciseFloats(
	v interface{},
	ctx *fasthttp.RequestCtx,
) ([]byte, bool) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	dataBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("Unable to marshal %s\n", reflect.TypeOf(v))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return nil, false
	}

	return dataBytes, true
}

// http://jsoniter.com/migrate-from-go-std.html
// marshals floats with 6 digits precision (lossy), which is significantly faster
func Marshal(
	v interface{},
	ctx *fasthttp.RequestCtx,
) ([]byte, bool) {
	stream := jsoniter.ConfigFastest.BorrowStream(nil)
	defer jsoniter.ConfigFastest.ReturnStream(stream)
	stream.WriteVal(v)
	if stream.Error != nil {
		log.Printf("Unable to marshal %s\n", reflect.TypeOf(v))
		log.Print(stream.Error)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return nil, false
	}

	return stream.Buffer(), true
}

func MarshalAux(
	v interface{},
	errorMessage string,
) ([]byte, bool) {
	stream := jsoniter.ConfigFastest.BorrowStream(nil)
	defer jsoniter.ConfigFastest.ReturnStream(stream)
	stream.WriteVal(v)
	if stream.Error != nil {
		log.Printf(errorMessage)
		log.Printf("Unable to marshal %s\n", reflect.TypeOf(v))
		log.Print(stream.Error)

		return nil, false
	}

	return stream.Buffer(), true
}
