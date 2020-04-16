package mysql

import "database/sql"

// Database model
type Database struct {
	DB  *sql.DB
	TBL string
}

// MySQL database connection
var MySQL Database
