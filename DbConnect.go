package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"


type DbConnect struct {
    conn *sql.DB
}

func (d *DbConnect) connect() *sql.DB {
    d.conn, _ = sql.Open("mysql", "master:mastersecret@(11.11.11.11:3306)/busca")
    return d.conn
}

func (d DbConnect) isConnected() bool {
    if(d.conn == nil) {
        log.Fatal("Not Connected!")
    }

    if err := d.conn.Ping(); err != nil {
        log.Fatal(err)
    }

    return true
}

func (d DbConnect) getResult(rows *sql.Rows) []map[string]string {
    columns, _ := rows.Columns()
    IColumns := make([]interface{}, len(columns))
    IValues := make([]interface{}, len(columns))

    for i, _ := range columns {
        IColumns[i] = &IValues[i]
    }

    result := []map[string]string{}

    for rows.Next() {

        rows.Scan(IColumns...)

        m := make(map[string]string)

        for i, col := range columns {
            m[col] = convertToHuman(IValues[i])
        }
        result = append(result, m)
    }
    return result
}

func convertToHuman(data interface{}) string {
    val, _ := data.([]byte)
    return string(val)
}
