package main

import (
	"fmt"
	"time"

	"github.com/cia-rana/tson"
)

type Person struct {
	Feature   Feature    `json:"feature"`
	CreatedAt *time.Time `json:"created_at"`
}

type Feature struct {
	Name  string     `json:"name"`
	Birth *time.Time `json:"birth"`
}

var jsonString = `
{
	"feature": {
		"name": "cia-rana",
		"birth": "1988-10-10"
	},
	"created_at": "2017-12-05"
}
`

func main() {
	var (
		person Person
		err    error
	)

	err = tson.Unmarshal([]byte(jsonString), &person)
	fmt.Println(person, err)

	tson.SetLayout(`2006-01-02`)
	err = tson.Unmarshal([]byte(jsonString), &person)
	fmt.Println(person, err)
}
