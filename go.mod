module bitbucket.org/votecube/votecube-go-lib

go 1.13

require (
	github.com/fasthttp/router v0.5.3 // indirect
	github.com/klauspost/compress v1.9.7 // indirect
	github.com/lib/pq v1.3.0
	github.com/scylladb/gocqlx v1.3.1 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.3.1
