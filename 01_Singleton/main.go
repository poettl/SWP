package main

import (
	DBHelper "swp/singleton/dbhelper"
)

func main() {
	tryDynamic()
}

func tryDynamic() {
	db := DBHelper.GetDB()
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	db = DBHelper.GetDB()
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	DBHelper.CloseDB()
}
