package main

import (
	"fmt"
	controllers "goapp1/handlers"

	jsoniter "github.com/json-iterator/go"
)

type AppInfo struct {
	Name string
}

func main() {
	info := AppInfo{
		Name: "GoApp",
	}
	jsonString, _ := jsoniter.Marshal(&info)
	fmt.Println(string(jsonString))
	controllers.ShowInfo()
}
