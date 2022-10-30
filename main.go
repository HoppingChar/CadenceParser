package main

// import (
// 	"cadence-parser/parser"
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	path := "./flow_token.cdc"
// 	code, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	res_json := parser.Parse2Json(code)
// 	fmt.Printf("%s\n", res_json)
// }

import (
	"cadence-parser/rest"
)

func main() {
	r := rest.SetupRouter()
	r.Run(":8080")
}
