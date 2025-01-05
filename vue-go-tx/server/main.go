package main

import (
	"server/mysql"
	"server/router"
)

func main() {
	mysql.Init()
	router.Init()
}
