package utils

import (
	"bitbucket.org/votecube/votecube-go-lib/sequence"
	"database/sql"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func SetupDb(
	path string,
) *sql.DB {
	DB, err := sql.Open("postgres", "postgresql://"+path+"/votecube?sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	return DB
}

func GetSeq(
	sequence sequence.Sequence,
	ctx *fasthttp.RequestCtx,
) (int64, bool) {
	idCursor, err := sequence.GetCursor(1)
	if err != nil {
		log.Printf("AddOpinion: Unable to access %s sequence\n", sequence.Name)
		log.Print(err)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return 0, false
	}

	return idCursor.Next(), true
}
