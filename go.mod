module bitbucket.org/votecube/votecube-go-lib

go 1.13

require (
	github.com/gocql/gocql v0.0.0-20200103014340-68f928edb90a
	github.com/json-iterator/go v1.1.9
	github.com/klauspost/compress v1.9.7
	github.com/scylladb/gocqlx v1.3.3
	github.com/valyala/fasthttp v1.7.0
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.3.1
