package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func ReturnId(
	id int64,
	ctx *fasthttp.RequestCtx,
) {
	byteMask, idSignificantBytes, ok := encodeInt64(id, ctx)
	if !ok {
		return
	}

	// https://github.com/valyala/fasthttp/issues/444
	ctx.Response.Reset()
	ctx.SetStatusCode(http.StatusCreated)
	ctx.SetContentType("vcb")
	ctx.Response.AppendBody([]byte{byteMask})
	ctx.Response.AppendBody(idSignificantBytes)
}

func ReturnIdAndCreateEs(
	id int64,
	createEs int64,
	ctx *fasthttp.RequestCtx,
) {
	byteMask, idSignificantBytes, createEsSignificantBytes, ok := encodeIdAndCreateEs(id, createEs, ctx)
	if !ok {
		return
	}

	// https://github.com/valyala/fasthttp/issues/444
	ctx.Response.Reset()
	ctx.SetStatusCode(http.StatusCreated)
	ctx.SetContentType("vcb")
	ctx.Response.AppendBody([]byte{byteMask})
	ctx.Response.AppendBody(idSignificantBytes)
	ctx.Response.AppendBody(createEsSignificantBytes)
}

func ReturnIdAndCreateEsAndVersion(
	id int64,
	createEs int64,
	version int16,
	ctx *fasthttp.RequestCtx,
) {
	byteMask, idSignificantBytes, createEsSignificantBytes, versionSignificantBytes,
		ok := encodeIdAndCreateEsAndVersion(id, createEs, version, ctx)
	if !ok {
		return
	}

	// https://github.com/valyala/fasthttp/issues/444
	ctx.Response.Reset()
	ctx.SetStatusCode(http.StatusCreated)
	ctx.SetContentType("vcb")
	ctx.Response.AppendBody([]byte{byteMask})
	ctx.Response.AppendBody(idSignificantBytes)
	ctx.Response.AppendBody(createEsSignificantBytes)
	ctx.Response.AppendBody(versionSignificantBytes)
}

func ReturnPartitionPeriodAndIdAndVersion(
	partitionPeriod int32,
	id int64,
	version int32,
	ctx *fasthttp.RequestCtx,
) {
	byteMask, partitionPeriodBytes, idSignificantBytes, versionSignificantBytes,
		ok := encodePartitionPeriodAndIdAndVersion(partitionPeriod, id, version, ctx)
	if !ok {
		return
	}

	// https://github.com/valyala/fasthttp/issues/444
	ctx.Response.Reset()
	ctx.SetStatusCode(http.StatusCreated)
	ctx.SetContentType("vcb")
	ctx.Response.AppendBody([]byte{byteMask})
	ctx.Response.AppendBody(partitionPeriodBytes)
	ctx.Response.AppendBody(idSignificantBytes)
	ctx.Response.AppendBody(versionSignificantBytes)
}

func encodeIdAndCreateEsAndVersion(
	id int64,
	createEs int64,
	version int16,
	ctx *fasthttp.RequestCtx,
) (byte, []byte, []byte, []byte, bool) {
	idByteMask, idSignificantBytes, ok := encodeInt64(id, ctx)
	if !ok {
		return 0, nil, nil, nil, false
	}

	esByteMask, createEsSignificantBytes, ok := encodeEpochSeconds(createEs, ctx)
	if !ok {
		return 0, nil, nil, nil, false
	}
	versionByteMask, versionSignificantBytes, ok := encodeUint16(version, ctx)
	if !ok {
		return 0, nil, nil, nil, false
	}

	finalByteMask := versionByteMask<<4 + esByteMask<<3 + idByteMask

	fmt.Println("")
	fmt.Println("id:       %d", id)
	fmt.Println("createEs: %d", createEs)
	fmt.Println("version: %d", version)
	fmt.Printf("%d ", finalByteMask)
	for _, n := range idSignificantBytes {
		fmt.Printf("%d ", n)
	}
	for _, n := range createEsSignificantBytes {
		fmt.Printf("%d ", n)
	}
	for _, n := range versionSignificantBytes {
		fmt.Printf("%d ", n)
	}
	fmt.Println("")

	return finalByteMask, idSignificantBytes, createEsSignificantBytes, versionSignificantBytes, true
}

func encodePartitionPeriodAndIdAndVersion(
	partitionPeriod int32,
	id int64,
	version int32,
	ctx *fasthttp.RequestCtx,
) (byte, []byte, []byte, []byte, bool) {
	partitionPeriodBytes, ok := getBytes(partitionPeriod, ctx)
	if !ok {
		return 0, nil, nil, false
	}

	idByteMask, idSignificantBytes, ok := encodeInt64(id, ctx)
	if !ok {
		return 0, nil, nil, false
	}

	versionByteMask, versionSignificantBytes, ok := encodeInt32(id, ctx)
	if !ok {
		return 0, nil, nil, false
	}

	finalByteMask := idByteMask<<2 + versionByteMask

	return finalByteMask, partitionPeriodBytes, idSignificantBytes, versionSignificantBytes, true
}

func encodeIdAndCreateEs(
	id int64,
	createEs int64,
	ctx *fasthttp.RequestCtx,
) (byte, []byte, []byte, bool) {
	idByteMask, idSignificantBytes, ok := encodeInt64(id, ctx)
	if !ok {
		return 0, nil, nil, false
	}

	esByteMask, createEsSignificantBytes, ok := encodeEpochSeconds(createEs, ctx)
	if !ok {
		return 0, nil, nil, false
	}

	finalByteMask := esByteMask<<3 + idByteMask
	/*
		fmt.Println("")
		fmt.Println("id:       %d", id)
		fmt.Println("createEs: %d", createEs)
		fmt.Printf("%d ", finalByteMask)
		for _, n := range idSignificantBytes {
			fmt.Printf("%d ", n)
		}
		for _, n := range createEsSignificantBytes {
			fmt.Printf("%d ", n)
		}
		fmt.Println("")
	*/

	return finalByteMask, idSignificantBytes, createEsSignificantBytes, true
}

func encodeEpochSeconds(
	epochSeconds int64,
	ctx *fasthttp.RequestCtx,
) (byte, []byte, bool) {
	esBytes, ok := getBytes(epochSeconds, ctx)
	if !ok {
		return 0, nil, false
	}

	var (
		byteMask           byte
		esSignificantBytes []byte
	)

	if epochSeconds < 4294967296 {
		esSignificantBytes = esBytes[0:4]
		byteMask = 0
	} else {
		esSignificantBytes = esBytes[0:5]
		byteMask = 1
	}

	return byteMask, esSignificantBytes, true
}

func encodeInt64(
	number int64,
	ctx *fasthttp.RequestCtx,
) (byte, []byte, bool) {
	bytes, ok := getBytes(number, ctx)
	if !ok {
		return 0, nil, false
	}

	var significantBytes []byte
	var byteMask byte

	if number < 256 {
		significantBytes = bytes[0:1]
		byteMask = 0
	} else if number < 65536 {
		significantBytes = bytes[0:2]
		byteMask = 1
	} else if number < 16777216 {
		significantBytes = bytes[0:3]
		byteMask = 2
	} else if number < 4294967296 {
		significantBytes = bytes[0:4]
		byteMask = 3
	} else if number < 1099511627776 {
		significantBytes = bytes[0:5]
		byteMask = 4
	} else if number < 281474976710656 {
		significantBytes = bytes[0:6]
		byteMask = 5
	} else if number < 72057594037927936 {
		significantBytes = bytes[0:7]
		byteMask = 6
	} else {
		significantBytes = bytes
		byteMask = 7
	}

	return byteMask, significantBytes, true
}

func encodeInt32(
	number int32,
	ctx *fasthttp.RequestCtx,
) (byte, []byte, bool) {
	bytes, ok := getBytes(number, ctx)
	if !ok {
		return 0, nil, false
	}

	var significantBytes []byte
	var byteMask byte

	if number < 256 {
		significantBytes = bytes[0:1]
		byteMask = 0
	} else if number < 65536 {
		significantBytes = bytes[0:2]
		byteMask = 1
	} else if number < 16777216 {
		significantBytes = bytes[0:3]
		byteMask = 2
	} else {
		significantBytes = bytes[0:4]
		byteMask = 3
	}

	return byteMask, significantBytes, true
}

func encodeUint16(
	number uint16,
	ctx *fasthttp.RequestCtx,
) (byte, []byte, bool) {
	bytes, ok := getBytes(number, ctx)
	if !ok {
		return 0, nil, false
	}

	var significantBytes []byte
	var byteMask byte

	if number < 256 {
		significantBytes = bytes[0:1]
		byteMask = 0
	} else {
		significantBytes = bytes[0:2]
		byteMask = 1
	}

	return byteMask, significantBytes, true
}

func getBytes(
	data interface{},
	ctx *fasthttp.RequestCtx,
) ([]byte, bool) {
	idBuffer := new(bytes.Buffer)
	err := binary.Write(idBuffer, binary.LittleEndian, data)
	if err != nil {
		log.Print("binary.Write failed:")
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return nil, false
	}

	return idBuffer.Bytes(), true
}
