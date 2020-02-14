package vespa

import "bitbucket.org/votecube/votecube-go-lib/model/data"

type Error struct {
	Description string `json:"description"`
	Id          int    `json:"id"`
}

type ModResponse struct {
	Errors []Error `json:"errors"`
	Id     string  `json:"id"`
	Path   string  `json:"path"`
}

type OpinionAdd struct {
	Fields data.Opinion `json:"fields"`
}

type StringFieldUpdate struct {
	Assign string `json:"assign"`
}

type OpinionUpdateFields struct {
	Text StringFieldUpdate `json:"text"`
}

type OpinionUpdate struct {
	Fields OpinionUpdateFields `json:"fields"`
}

type PollAdd struct {
	Fields data.Poll `json:"fields"`
}
