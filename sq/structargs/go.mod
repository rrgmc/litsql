module github.com/rrgmc/litsql/sq/structargs

go 1.22.0

require (
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rrgmc/litsql v0.7.1
	gotest.tools/v3 v3.5.1
)

require github.com/google/go-cmp v0.5.9 // indirect

replace github.com/rrgmc/litsql => ../..
