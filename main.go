package main

import (
	"github.com/higashi000/noa/router"
)

func main() {
	r := router.NewRouter()

	r.Run(":5000")
}
