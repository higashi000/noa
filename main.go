package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/higashi000/noa/router"
)

func main() {
	r := router.NewRouter()

	pprof.Register(r)
	r.Run(":5000")
}
