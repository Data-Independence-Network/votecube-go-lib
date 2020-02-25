package vespa

import (
	"bitbucket.org/votecube/votecube-go-lib/model/data"
	"bitbucket.org/votecube/votecube-go-lib/model/vespa"
	"bitbucket.org/votecube/votecube-go-lib/utils"
	"github.com/valyala/fasthttp"
	"strconv"
)

func AddFeedback(
	urlPrefix string,
	dataFeedback data.Feedback,
	ctx *fasthttp.RequestCtx,
) bool {
	modResponse := vespa.ModResponse{}
	feedbackAdd := vespa.FeedbackAdd{
		Fields: vespa.Feedback{
			AgeSuitability: dataFeedback.AgeSuitability,
			Contents:       dataFeedback.Contents,
			CreateDtb:      dataFeedback.CreateDtb,
			FeedbackTypeId: dataFeedback.FeedbackTypeId,
			Title:          dataFeedback.Title,
			UserId:         dataFeedback.UserId,
		},
	}
	if !utils.PostObject(
		urlPrefix+"/document/v1/feedback/feedback/docid/"+strconv.FormatInt(dataFeedback.Id, 10),
		feedbackAdd,
		modResponse,
		ctx,
	) {
		return false
	}

	return checkModResponse(modResponse, ctx)
}

func AddFeedbackComment(
	urlPrefix string,
	dataFeedbackComment data.FeedbackComment,
	ctx *fasthttp.RequestCtx,
) bool {
	modResponse := vespa.ModResponse{}

	opinionAdd := vespa.FeedbackCommentAdd{
		Fields: vespa.FeedbackComment{
			AgeSuitability: dataFeedbackComment.AgeSuitability,
			CreateDtb:      dataFeedbackComment.CreateDtb,
			FeedbackTypeId: dataFeedbackComment.FeedbackTypeId,
			FeedbackId:     dataFeedbackComment.FeedbackId,
			Text:           dataFeedbackComment.Text,
			UserId:         dataFeedbackComment.UserId,
		},
	}

	if !utils.PostObject(
		urlPrefix+"/document/v1/feedback_comment/feedback_comment/docid/"+strconv.FormatInt(dataFeedbackComment.Id, 10),
		opinionAdd,
		modResponse,
		ctx,
	) {
		return false
	}

	return checkModResponse(modResponse, ctx)
}

func GetFeedbackById(
	urlPrefix string,
	feedbackId int64,
	ctx *fasthttp.RequestCtx,
) (*vespa.Feedback, bool) {
	getFeedbackResponse := vespa.GetFeedbackResponse{}

	if !utils.GetObject(
		urlPrefix+"/document/v1/feedback_comment/feedback_comment/docid/"+strconv.FormatInt(feedbackId, 10),
		getFeedbackResponse,
		ctx,
	) {
		return nil, false
	}

	return &getFeedbackResponse.Record, true
}
