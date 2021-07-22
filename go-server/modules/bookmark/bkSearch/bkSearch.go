package bkSsearch

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"log"
	"pettyfox.top/bookmark/conf"
	"pettyfox.top/bookmark/modules/bookmark"
	"pettyfox.top/bookmark/modules/word"
	"strings"
)

var rc *redisearch.Client
var ac *redisearch.Autocompleter
var sc *redisearch.Schema

func Init() {
	rc = redisearch.NewClient(conf.RedisSearchConf["address"], "bk_index_1")
	ac = redisearch.NewAutocompleter(conf.RedisSearchConf["address"], "bk_ac_1")
	// Create a schema
	sc = redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("id")).
		AddField(redisearch.NewTextField("desc")).
		AddField(redisearch.NewTextFieldOptions("name", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewTextField("url"))
	a, _ := ac.Length()
	log.Println("suggest len:", a)
	// Drop an existing index. If the index does not exist an error is returned
	info, err := rc.Info()
	if err != nil {
		// Create the index with the given schema
		if err2 := rc.CreateIndex(sc); err2 != nil {
			log.Println("create index err ", err2)
		} else {
			log.Println("create index")
		}

	} else {
		log.Println("index exist", info)
	}

}
func RestIndex() {
	rc.Drop()
	if err2 := rc.CreateIndex(sc); err2 != nil {
		log.Println("recreate index err ", err2)
	} else {
		log.Println("recreate index")
	}
	ac = redisearch.NewAutocompleter(conf.RedisSearchConf["address"], "bk_ac_1")
}
func RemoveDocIndex(bookmarkId string) {
	rc.DeleteDocument("bk_doc_"+bookmarkId)
	//TODO 清理SUGGEST，根据推入的关键词计数器来清理，如果计数器为0，删除对应的建议，采用增强版的布隆过滤器
}
func SetDocIndex(userId string, bookmark bookmark.Bookmark) {
	if len(bookmark.Id) == 0 || len(bookmark.Url) == 0 {
		return
	}
	RemoveDocIndex(bookmark.Id)
	nameIndex := ""
	descIndex := ""
	urlIndex := ""
	if len(bookmark.Name) > 0 {
		nameSplit := word.Word2Index(bookmark.Name)
		nameIndex = strings.Join(nameSplit, " ")
		println(nameSplit, nameIndex)
		//添加建议
		for _, s := range nameSplit {
			if len(s) <= 1 {
				continue
			}
			err := ac.AddTerms(redisearch.Suggestion{Term: s, Score: 1})
			if err != nil {
				log.Println("suggest name err", err)
			} else {
				log.Println("suggest name ok", s)
			}
		}
	}
	if len(bookmark.Desc) > 0 {
		descSplit := word.Word2Index(bookmark.Desc)
		descIndex = strings.Join(descSplit, " ")

		//添加建议
		for _, s := range descSplit {
			if len(s) <= 1 {
				continue
			}
			ac.AddTerms(redisearch.Suggestion{Term: s, Score: 1})
		}
	}
	if len(bookmark.Url) > 0 {
		urlSplit := word.Word2Index(bookmark.Url)
		urlIndex = strings.Join(urlSplit, " ")

		//添加建议
		for _, s := range urlSplit {
			if len(s) <= 1 {
				continue
			}
			ac.AddTerms(redisearch.Suggestion{Term: s, Score: 1})

		}
	}

	doc := redisearch.NewDocument("bk_doc_"+bookmark.Id, 1.0)
	doc.Set("id", bookmark.Id)
	//建立索引
	if len(nameIndex) > 0 {
		doc.Set("name", nameIndex)
	}
	if len(urlIndex) > 0 {
		doc.Set("url", urlIndex)
	}
	if len(descIndex) > 0 {
		doc.Set("desc", urlIndex)
	}
	if err := rc.Index(doc); err != nil {
		log.Fatal(err)
	}

}
func Search(userId, word string, offset, limit int) []string {
	results := make([]string, 0)
	if len(word) == 0 {
		return results
	}
	docs, _, err := rc.Search(redisearch.NewQuery(word).
		Limit(offset, limit).
		SetReturnFields("id"))
	if err != nil {
		log.Println(err)
	}
	if len(docs) > 0 {
		for _, doc := range docs {
			results = append(results, fmt.Sprintf("%v", doc.Properties["id"]))
		}
	}
	//fmt.Printf("search totoal:%v rp:%v,rs:%v", total, docs, results)
	return results
}
func Suggest(userId, word string, offset, limit int) []string {
	if len(word) == 0 {
		return nil
	}
	ss, err := ac.SuggestOpts(word, redisearch.DefaultSuggestOptions)
	if err != nil {
		log.Println(err)
	}
	results := make([]string, 0)
	if len(ss) > 0 {
		for _, s := range ss {
			results = append(results, s.Term)
		}
	}

	//fmt.Printf("search  suggest:%v", ss)
	return results
}
