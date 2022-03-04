module github.com/evolbioinf/neighbors/index

go 1.17

replace github.com/evolbioinf/neighbors/util => ../util

replace github.com/evolbioinf/neighbors/tdb => ../tdb/

replace github.com/evolbioinf/neighbors/tax => ../tax/

require (
	github.com/evolbioinf/clio v0.0.0-20210309091639-82cb91a31b0c
	github.com/evolbioinf/neighbors/tax v0.0.0-00010101000000-000000000000
	github.com/evolbioinf/neighbors/util v0.0.0-00010101000000-000000000000
)
