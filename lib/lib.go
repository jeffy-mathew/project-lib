package lib

import (
	"embed"
	_ "embed"
	"log"
)

//go:embed test.json
var f embed.FS

var data []byte

func init() {
	var err error
	data, err = f.ReadFile("test.json")
	if err != nil {
		log.Println(err)
	}
}

func GetData() string {
	return string(data)
}
