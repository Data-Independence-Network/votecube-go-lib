# votecube-go-lib

Common Go libraries


## Install Instructions

Install SqlBoiler

set GOPATH=

go get -u -t github.com/volatiletech/sqlboiler

go get -u github.com/glerchundi/sqlboiler-crdb
NOT go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

sqlboiler -o model\crdb\models\ -p crdb crdb
