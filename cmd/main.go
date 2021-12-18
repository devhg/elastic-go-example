package main

import (
	"log"

	"github.com/devhg/es/internal/conf"
)

// https://segmentfault.com/a/1190000024438897
func main() {
	config := conf.LoadAndInit("./config/config.yml")

	s, stop := initApp(config)
	defer stop()

	log.Fatal(s.ListenAndServe())
}
