package utils

import (
	"github.com/scylladb/gocqlx"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"reflect"
)

func Insert(
	preparedInsert *gocqlx.Queryx,
	r interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	return exec("INSERT", preparedInsert, r, ctx)
}

func Select(
	preparedSelect *gocqlx.Queryx,
	dest interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	if err := preparedSelect.Select(dest); err != nil {
		log.Printf("Error during SELECT of \"%s\"\n", reflect.TypeOf(dest))
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
		log.Printf("Error during SELECT COUNT of %s\n", ctx.UserValue("recordType"))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return false
	}

	return true
}

func Update(
	preparedInsert *gocqlx.Queryx,
	r interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	return exec("UPDATE", preparedInsert, r, ctx)
}

func ExecAux(
	preparedModification *gocqlx.Queryx,
	r interface{},
	errorMessage string,
) bool {
	statement := preparedModification.BindStruct(r)

	if err := statement.Exec(); err != nil {
		log.Println(errorMessage)
		log.Print(err)
		return false
	}

	return true
}

func exec(
	operationType string,
	preparedInsert *gocqlx.Queryx,
	r interface{},
	ctx *fasthttp.RequestCtx,
) bool {
	statement := preparedInsert.BindStruct(r)

	if err := statement.Exec(); err != nil {
		log.Printf("Error during %s of \"%s\"\n", operationType, reflect.TypeOf(r))
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return false
	}

	return true
}
