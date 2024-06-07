module github.com/rrgmc/litsql/sq/reflectxargs

go 1.22.0

require (
	github.com/iancoleman/strcase v0.3.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/rrgmc/litsql v0.8.1
	gotest.tools/v3 v3.5.1
)

require github.com/google/go-cmp v0.5.9 // indirect

replace github.com/rrgmc/litsql => ../..
