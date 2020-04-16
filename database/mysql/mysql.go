package mysql

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/MahdiRazaqi/nevees-backend/config"
	// load mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect to MySQL
func Connect() error {
	connection, err := sql.Open("mysql", config.CFG.MySQL.User+":"+config.CFG.MySQL.Password+"@tcp("+config.CFG.MySQL.Host+")/"+config.CFG.MySQL.DB)
	if err != nil {
		return err
	}

	// check connection
	if err := connection.Ping(); err != nil {
		return err
	}
	MySQL.DB = connection

	return nil
}

// Table create and select table
func (m *Database) Table(table string, i interface{}) *Database {
	if err := createTable(table, i); err != nil {
		return nil
	}
	m.TBL = table

	return m
}

func checkExistTable(table string) bool {
	_, status := MySQL.DB.Query("SELECT * FROM " + table)
	if status != nil {
		return false
	}
	return true
}

func createTable(table string, i interface{}) error {
	if !checkExistTable(table) {
		columns := strings.Join(readTags(i), ",")
		tbl, err := MySQL.DB.Query("CREATE TABLE " + table + " (" + columns + ")")
		if err != nil {
			return err
		}
		defer tbl.Close()
	}
	return nil
}

func readTags(i interface{}) []string {
	t := reflect.TypeOf(i)
	val := reflect.Indirect(reflect.ValueOf(i))

	tags := []string{}
	for count := 0; count < t.NumField(); count++ {
		tag := val.Type().Field(count).Tag.Get("mysql")
		tags = append(tags, tag)
	}

	return tags
}

func getColumnsData(i interface{}) []interface{} {
	t := reflect.ValueOf(i)

	values := []interface{}{}
	for count := 0; count < t.NumField(); count++ {
		values = append(values, t.Field(count).Interface())
	}

	return values
}
