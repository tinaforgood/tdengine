package main

import (
	"database/sql"
	"fmt"

	"github.com/libgo/logx"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

// https://docs.taosdata.com/develop/

func main() {
	var taosDSN = "root:tt0819@http(localhost:6041)/"
	taos, err := sql.Open("taosRestful", taosDSN)
	if err != nil {
		logx.Errorf("failed to connect TDengine, err:", err)
		return
	}
	fmt.Println("Connected")
	createStable(taos)

	defer taos.Close()
}

func createStable(taos *sql.DB) {
	_, err := taos.Exec("CREATE DATABASE safeguarding keep 14 DURATION 10 BUFFER 16 WAL_LEVEL 1")
	if err != nil {
		logx.Errorf("failed to create database, err:", err)
	}
	_, err = taos.Exec("CREATE STABLE safeguarding.temperature (ts TIMESTAMP, current FLOAT) TAGS (seniorId BINARY(32))")
	if err != nil {
		logx.Errorf("failed to create stable, err:%v", err)
	}
}
