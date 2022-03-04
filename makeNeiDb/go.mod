module github.com/evolbioinf/neighbors/makeNeiDb

go 1.17

replace github.com/evolbioinf/neighbors/tdb => ../tdb/

replace github.com/evolbioinf/neighbors/util => ../util

require (
	github.com/evolbioinf/clio v0.0.0-20210309091639-82cb91a31b0c
	github.com/evolbioinf/neighbors/tdb v0.0.0-00010101000000-000000000000
	github.com/evolbioinf/neighbors/util v0.0.0-00010101000000-000000000000
)

require github.com/mattn/go-sqlite3 v1.14.12 // indirect
