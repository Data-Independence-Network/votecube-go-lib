package vespa

import (
	"bitbucket.org/votecube/votecube-go-lib/model/data"
	"bitbucket.org/votecube/votecube-go-lib/model/vespa"
	"bitbucket.org/votecube/votecube-go-lib/utils"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"strconv"
)

func AddOpinion(
	urlPrefix string,
	opinion data.Opinion,
	ctx *fasthttp.RequestCtx,
) bool {
	modResponse := vespa.ModResponse{}
	opinionAdd := vespa.OpinionAdd{
		Fields: opinion,
	}
	if !utils.PostObject(
		urlPrefix+"/document/v1/opinion/opinion/docid/"+strconv.FormatInt(opinion.Id, 10),
		opinionAdd,
		modResponse,
		ctx,
	) {
		return false
	}

	return checkModResponse(modResponse, ctx)
}

func AddPoll(
	urlPrefix string,
	poll data.Poll,
	ctx *fasthttp.RequestCtx,
) bool {
	modResponse := vespa.ModResponse{}
	pollAdd := vespa.PollAdd{
		Fields: poll,
	}
	if !utils.PostObject(
		urlPrefix+"/document/v1/poll/poll/docid/"+strconv.FormatInt(poll.Id, 10),
		pollAdd,
		modResponse,
		ctx,
	) {
		return false
	}

	return checkModResponse(modResponse, ctx)
}

func UpdateOpinion(
	urlPrefix string,
	opinion data.Opinion,
	ctx *fasthttp.RequestCtx,
) bool {
	opinionUpdate := vespa.OpinionUpdate{
		Fields: vespa.OpinionUpdateFields{
			Text: vespa.StringFieldUpdate{
				Assign: opinion.Text,
			},
		},
	}
	modResponse := vespa.ModResponse{}
	if !utils.PutObject(
		urlPrefix+"/document/v1/opinion/opinion/docid/"+strconv.FormatInt(opinion.Id, 10),
		opinionUpdate,
		modResponse,
		ctx,
	) {
		return false
	}

	return checkModResponse(modResponse, ctx)
}
func checkModResponse(
	response vespa.ModResponse,
	ctx *fasthttp.RequestCtx,
) bool {

	if len(response.Errors) > 0 {
		log.Print(response.Errors[0].Description)
		log.Print(response.Errors[0].Id)
		ctx.Error("Internal Server Error", http.StatusInternalServerError)

		return false
	}

	return true
}
