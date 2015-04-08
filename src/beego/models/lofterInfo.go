package models

import (
	"encoding/json"
	"fmt"
)

type Lofter struct {
	Name string
	Rss  string
}

var (
	LofterList []Lofter
)

func init() {

	lofter := Lofter{"jiannanren", "jiannanren.lofter.com/rss"}

	LofterList = []Lofter{lofter}
	// LofterList[0].Name = "jiannanren"
	// LofterList[0].Rss = "jiannanren.lofter.com/rss"
}

func GetAll() string {

	b, err := json.Marshal(LofterList)
	if err != nil {
		fmt.Println("json err:", err)
	}
	return string(b)
}
