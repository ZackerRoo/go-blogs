package main

import "go-blogs/common/initialize"

func main() {
	print("hello")
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Router()

}
