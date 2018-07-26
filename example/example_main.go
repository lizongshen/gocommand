package main

import (
	"log"

	"github.com/lizongshen/gocommand"
)

// 这是一个示例
func main() {
	_, out, err := gocommand.NewCommand().Exec("ls /")
	if err != nil {
		log.Panic(err)
	}

	log.Println(out)

}
