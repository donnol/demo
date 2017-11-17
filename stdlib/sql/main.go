package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	drivers := sql.Drivers()
	fmt.Println(drivers)

	ct := &sql.ColumnType{}
	dtn := ct.DatabaseTypeName()
	fmt.Println(dtn)
}
