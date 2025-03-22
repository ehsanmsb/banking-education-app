package main

import (
	"github.com/ehsanmsb/banking-education-app/internal/db"
)

// I am going to implement Hexagonal Architecture (adaptors and ports)
func main() {
	db.Database()
	//app.StartServer()
}
