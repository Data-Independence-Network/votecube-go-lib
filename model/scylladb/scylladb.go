package scylladb

type Opinion struct {
	OpinionId       uint64
	RootOpinionId   uint64
	PollId          uint64
	ParentId        uint64
	Position        string
	CreateHour      string
	UserId          uint64
	CreateEs        int64
	UpdateEs        int64
	Version         uint16
	Data            []byte
	InsertProcessed bool
}

type OpinionUpdate struct {
	OpinionId       uint64
	RootOpinionId   uint64
	PollId          uint64
	UpdateHour      string
	UserId          uint64
	UpdateEs        int64
	Data            []byte
	Version         uint16
	UpdateProcessed bool
}

type Poll struct {
	PollId     uint64
	ThemeId    uint64
	LocationId uint32
	UserId     uint64
	CreateHour string
	CreateEs   int64
	Data       []byte
	BatchId    int
}

type Thread struct {
	PollId   uint64
	UserId   uint64
	CreateEs int64
	Data     []byte
}
