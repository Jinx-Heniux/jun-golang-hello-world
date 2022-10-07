package struct_

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `json:"name" doc:"我的名字"`
}

func findDoc(stru interface{}) map[string]string {
	t := reflect.TypeOf(stru).Elem()
	doc := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
	}

	return doc

}

func StructExample2() {
	var stru resume
	doc := findDoc(&stru)
	fmt.Printf("name字段为：%s\n", doc["name"])
}

/*
name字段为：我的名字
*/
