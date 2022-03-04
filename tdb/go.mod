module github.com/evolbioinf/neighbors/tdb

go 1.17

require (
	github.com/evolbioinf/neighbors/util v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.14.12
)

require github.com/evolbioinf/clio v0.0.0-20210309091639-82cb91a31b0c // indirect

replace github.com/evolbioinf/neighbors/util => ../util/
