# gocommand


	package main

	import (
		"github.com/lizongshen/gocommand"
		"log"
	)

	func main() {
		_, out, err := gocommand.NewCommand().Exec("ls /")
		if err != nil {
			log.Panic(err)
		}

		log.Println(out)
	}
