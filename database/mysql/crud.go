package mysql

import (
	"fmt"
)

// Insert data to database
func (d *Database) Insert(i interface{}) error {
	data := getColumnsData(i)
	columns := ""
	for i := 0; i < len(data); i++ {
		columns += "?"
		if i != len(data)-1 {
			columns += ","
		}
	}

	q, err := d.DB.Query("INSERT INTO "+d.TBL+" VALUES ("+columns+")", data...)
	if err != nil {
		return err
	}
	defer q.Close()

	return nil
}

// Select data from database
func (d *Database) Select(i interface{}, columns ...string) error {
	c := ""
	for range columns {
		c += "?"
		if i != len(columns)-1 {
			c += ","
		}
	}
	fmt.Println(c)
	// q, err := d.DB.Query("SELECT ? FROM "+d.TBL, columns...)
	// if err != nil {
	// 	return err
	// }
	// defer q.Close()

	return nil
}
