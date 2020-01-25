package scylladb

type Opinion struct {
	PollId          int64
	PartitionPeriod int32
	AgeSuitability  int64
	OpinionId       int64
	ThemeId         int64
	LocationId      int32
	IngestBatchId   int32
	Version         int16
	RootOpinionId   int64
	ParentOpinionId int64
	CreateEs        int64
	UserId          int64
	Data            []byte
	InsertProcessed bool
}

type OpinionUpdate struct {
	PollId          int64
	PartitionPeriod int32
	IngestBatchId   int32
	OpinionId       int64
	Version         int16
	UpdateProcessed bool
}

type Poll struct {
	PollId          int64
	ThemeId         int64
	LocationId      int32
	IngestBatchId   int32
	CreateEs        int64
	UserId          int64
	PartitionPeriod int32
	AgeSuitability  int64
	Data            []byte
	InsertProcessed bool
}

type UserSession struct {
	PartitionPeriod int32
	SessionId       string
	LastActionEs    int64
	KeepSignedIn    int8
	UserId          int64
	Data            []byte
}
