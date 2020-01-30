package utils

import (
	"bitbucket.org/votecube/votecube-go-lib/model/scylladb"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"sync"
)

var (
	selectUserSession *gocqlx.Queryx
	updateUserSession *gocqlx.Queryx
)

type UserContext struct {
	ctx                    *fasthttp.RequestCtx
	okUserSession          bool
	sessionPartitionPeriod int32
	sessionId              string
	userSessionRows        []scylladb.UserSession
	UserId                 int64
	waitGroup              sync.WaitGroup
	parallel               bool
}

func NewParallelUserContext(
	ctx *fasthttp.RequestCtx,
	userId int64,
	waitGroup sync.WaitGroup,
) *UserContext {
	sessionPartitionPeriod, ok := ParseInt32Param(
		"sessionPartitionPeriod", ctx)
	if !ok {
		return nil
	}
	sessionId := ctx.UserValue("sessionId").(string)

	return &UserContext{
		ctx:                    ctx,
		sessionPartitionPeriod: sessionPartitionPeriod,
		sessionId:              sessionId,
		UserId:                 userId,
		waitGroup:              waitGroup,
		parallel:               true,
	}
}

func NewUserContext(
	ctx *fasthttp.RequestCtx,
	userId int64,
) *UserContext {
	sessionPartitionPeriod, ok := ParseInt32Param(
		"sessionPartitionPeriod", ctx)
	if !ok {
		return nil
	}
	sessionId := ctx.UserValue("sessionId").(string)

	return &UserContext{
		ctx:                    ctx,
		sessionPartitionPeriod: sessionPartitionPeriod,
		sessionId:              sessionId,
		UserId:                 userId,
	}
}

func IsValidSession(
	ctx *fasthttp.RequestCtx,
	userId int64,
) bool {
	userContext := NewUserContext(ctx, userId)
	if userContext == nil {
		return false
	}
	GetUserSession(userContext)

	if !CheckSession(userContext) {
		return false
	}
	return true
}

func GetUserSession(
	userContext *UserContext,
) {
	if userContext.parallel {
		defer userContext.waitGroup.Done()
	}

	selectUserSessionQuery := selectUserSession.BindMap(qb.M{
		"partition_period": userContext.sessionPartitionPeriod,
		"session_id":       userContext.sessionId,
	})
	userContext.okUserSession =
		Select(selectUserSessionQuery, &userContext.userSessionRows, userContext.ctx)
}

func CheckSession(
	userContext *UserContext,
) bool {
	if !userContext.okUserSession {
		log.Printf("Error in session lookup by partition_period: %d, session_id: %s\n", userContext.sessionPartitionPeriod, userContext.sessionId)
		userContext.ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return false
	}
	if len(userContext.userSessionRows) != 1 {
		log.Printf("Did not find user_credentials with user_id: %d\n", userContext.UserId)
		userContext.ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return false
	}

	userSession := userContext.userSessionRows[0]

	if userSession.UserId != userContext.UserId {
		log.Printf("Session user_id: %d does not match provided user_id: %d\n", userSession.UserId, userContext.UserId)
		userContext.ctx.Error("Internal Server Error", http.StatusInternalServerError)
		return false
	}

	if userSession.KeepSignedIn == 0 {
		lastValidActionEs := GetCurrentEs() - (15*60 + 15)
		if userSession.LastActionEs < lastValidActionEs {
			userContext.ctx.Error("Session timed out", http.StatusUnauthorized)
			return false
		}

		updateUserSessionQuery := updateUserSession.BindMap(qb.M{
			"partition_period": userContext.sessionPartitionPeriod,
			"session_id":       userContext.sessionId,
		})

		newLastActionEs := scylladb.UserSession{
			LastActionEs: GetCurrentEs(),
		}
		Update(updateUserSessionQuery, newLastActionEs, userContext.ctx)
	}

	return true
}

func SetupAuthQueries(
	session *gocql.Session,
) {
	stmt, names := qb.Select("user_sessions").Columns(
		"user_id",
		"last_action_es",
		"keep_signed_in",
	).Where(
		qb.Eq("partition_period"),
		qb.Eq("session_id"),
	).BypassCache().ToCql()
	selectUserSession = gocqlx.Query(session.Query(stmt), names)

	stmt, names = qb.Update("user_sessions").Set(
		"last_action_es",
	).Where(
		qb.Eq("partition_period"),
		qb.Eq("session_id"),
	).ToCql()
	updateUserSession = gocqlx.Query(session.Query(stmt), names)
}
