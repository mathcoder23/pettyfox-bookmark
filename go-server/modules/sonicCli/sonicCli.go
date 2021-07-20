package sonicCli

import (
	"github.com/expectedsh/go-sonic/sonic"
	"pettyfox.top/bookmark/conf"
	"pettyfox.top/bookmark/modules/word"
	"strconv"
)

var Ingester sonic.Ingestable
var Search sonic.Searchable

func InitSonicCli() {
	port, err := strconv.Atoi(conf.SonicConf["port"])
	if err != nil {
		println(err)
	}
	ingester, err := sonic.NewIngester(conf.SonicConf["host"], port, conf.SonicConf["password"])
	if err != nil {
		println(err)
	}
	Ingester = ingester

	search, err := sonic.NewSearch(conf.SonicConf["host"], port, conf.SonicConf["password"])
	if err != nil {
		println(err)
	}
	Search = search

}
func Push(collection, bucket, object, text string) {
	list := word.Word2Index(text)
	for i := range list {
		err := Ingester.Push(collection, bucket, object, list[i], "")
		if err != nil {
			println(err)
		}
	}
	err := Ingester.Push(collection, bucket, object, text, "")
	if err != nil {
		println(err)
	}
}

func QUERY(collection, bucket, term string) []string {
	list := make([]string, 0)
	rs, err := Search.Query(collection, bucket, term, 20, 0, "")
	if err != nil {
		println(err)
		return nil
	}
	for i := 0; i < len(rs); i++ {
		if len(rs[i]) > 0 {
			list = append(list, rs[i])
		}
	}
	return list
}
