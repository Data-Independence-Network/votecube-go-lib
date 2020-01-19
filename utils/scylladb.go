package utils

import (
	"github.com/scylladb/gocqlx"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func Insert(
	preparedInsert *gocqlx.Queryx,
	v interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	return exec("INSERT", preparedInsert, v, ctx)
}

func Select(
	preparedSelect *gocqlx.Queryx,
	dest interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	if err := preparedSelect.Select(dest); err != nil {
		log.Printf("Error during SELECT of %s", ctx.UserValue("recordType"))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return false
	}

	return true
}

func SelectCount(
	preparedSelect *gocqlx.Queryx,
	dest *int,
	ctx *fasthttp.RequestCtx,
) bool {
	if err := preparedSelect.Select(dest); err != nil {
		log.Printf("Error during SELECT of %s", ctx.UserValue("recordType"))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return false
	}

	return true
}

func Update(
	preparedInsert *gocqlx.Queryx,
	v interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	return exec("UPDATE", preparedInsert, v, ctx)
}

func exec(
	operationType string,
	preparedInsert *gocqlx.Queryx,
	v interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	statement := preparedInsert.BindStruct(v)

	if err := statement.Exec(); err != nil {
		log.Printf("Error during %s of %s", operationType, ctx.UserValue("recordType"))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return false
	}

	return true
}
