module github.com/evolbioinf/neighbors

go 1.18

require (
	github.com/evolbioinf/clio v0.0.0-20230620105705-02d07225d27e
	github.com/evolbioinf/nwk v0.0.0-20220622085205-07dd42c9fcab
	github.com/mattn/go-sqlite3 v1.14.16
)

replace github.com/evolbioinf/neighbors/tdb => ../tdb

replace github.com/evolbioinf/neighbors/tax => ../tax
