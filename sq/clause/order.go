package clause

// general
const (
	OrderRawQuery = -10
)

// select
const (
	OrderWith        = 50
	OrderSelectStart = 150 // the SELECT should be printed after this order
	OrderDistinct    = 170
	OrderColumns     = 200
	OrderFrom        = 400
	OrderJoin        = 800
	OrderWhere       = 1000
	OrderGroupBy     = 1500
	OrderHaving      = 1600
	OrderWindow      = 1800
	OrderUnion       = 1900
	OrderOrderBy     = 2000
	OrderOffset      = 3000
	OrderLimit       = 3001
)

// insert
const (
	// OrderWith = 50
	OrderInsertStart        = 150 // the INSERT should be printed after this order
	OrderTable              = 200
	OrderInsertOverriding   = 250
	OrderValues             = 1000
	OrderInsertConflict     = 1500
	OrderInsertDuplicateKey = 1501
	OrderReturning          = 2000
)

// update
const (
	// OrderWith = 50
	OrderUpdateStart = 150 // the UPDATE should be printed after this order
	OrderUpdateOnly  = 151
	// OrderTable = 200
	OrderSet = 300
	// OrderFrom = 400
	// OrderJoin = 800
	// OrderWhere = 1000
	// OrderReturning = 2000
)

// delete
const (
	// OrderWith = 50
	OrderDeleteStart = 150 // the DELETE should be printed after this order
	OrderDeleteOnly  = 151
	OrderDeleteFrom  = 300
	// OrderFrom = 400 // == USING
	// OrderJoin = 800
	// OrderWhere = 1000
	// OrderReturning = 2000
)
