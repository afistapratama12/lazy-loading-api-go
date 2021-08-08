package service

const (
	findAll    = "SELECT * FROM products"
	queryWhere = "%s WHERE title LIKE '%%%s%%' OR description LIKE '%%%s%%'"
	querySort  = "%s ORDER BY price %s"
	queryPage  = `%s LIMIT %d OFFSET %d`
)
