package sonicCli

import (
	"fmt"
	"github.com/expectedsh/go-sonic/sonic"
	"pettyfox.top/bookmark/conf"
	"pettyfox.top/bookmark/modules/word"
	"strconv"
	"strings"
)

var Ingester sonic.Ingestable
var Search sonic.Searchable

func InitSonicCli() {
	port, err := strconv.Atoi(conf.SonicConf["port"])
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	ingester, err := sonic.NewIngester(conf.SonicConf["host"], port, conf.SonicConf["password"])
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	Ingester = ingester

	search, err := sonic.NewSearch(conf.SonicConf["host"], port, conf.SonicConf["password"])
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	Search = search

}
func checkSearch() {
	err := Search.Ping()
	if err == nil {
		return

	}
	fmt.Printf("search ping err %v", err)
	port, err := strconv.Atoi(conf.SonicConf["port"])
	search, err := sonic.NewSearch(conf.SonicConf["host"], port, conf.SonicConf["password"])
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	Search = search
}
func checkIngester() {
	err := Ingester.Ping()
	if err == nil {
		return

	}
	fmt.Printf("Ingester ping err %v", err)
	port, err := strconv.Atoi(conf.SonicConf["port"])
	ingester, err := sonic.NewIngester(conf.SonicConf["host"], port, conf.SonicConf["password"])
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	Ingester = ingester
}
func Push(collection, bucket, object, text string) {
	checkIngester()
	list := word.Word2Index(text)
	err := Ingester.Push(collection, bucket, object, strings.Join(list, " "), "")
	if err != nil {
		fmt.Printf("err:%v", err)
	}
}
func ResetIndex(collection, bucket, obj string) {
	checkIngester()
	if len(obj) > 0 {
		Ingester.FlushObject(collection, bucket, obj)
	} else if len(bucket) > 0 {
		Ingester.FlushBucket(collection, bucket)
	} else if len(collection) > 0 {
		Ingester.FlushCollection(collection)
	}

}
func QUERY(collection, bucket, term string) []string {
	checkSearch()
	list := make([]string, 0)
	rs, err := Search.Query(collection, bucket, term, 20, 0, "")
	if err != nil {
		fmt.Printf("query err %v", err)
		return nil
	}
	for i := 0; i < len(rs); i++ {
		if len(rs[i]) > 0 {
			list = append(list, rs[i])
		}
	}
	return list
}
func SearchSuggest(collection, bucket, term string, limit int) []string {
	checkSearch()
	rs, err := Search.Suggest(collection, bucket, term, limit)
	if err != nil {
		fmt.Printf("query suggest err %v %v", err, term)
	}
	return rs
}
