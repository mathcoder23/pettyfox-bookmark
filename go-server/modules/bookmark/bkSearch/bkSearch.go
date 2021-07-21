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

func Init() {
	rc = redisearch.NewClient(conf.RedisSearchConf["address"], "bk_index_1")
	ac = redisearch.NewAutocompleter(conf.RedisSearchConf["address"], "bk_ac_1")
	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("id")).
		AddField(redisearch.NewTextField("desc")).
		AddField(redisearch.NewTextFieldOptions("name", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewTextField("url"))

	// Drop an existing index. If the index does not exist an error is returned
	rc.Drop()
	// Create the index with the given schema
	if err := rc.CreateIndex(sc); err != nil {
		log.Fatal(err)
	}
}
func removeDocIndex(bookmarkId string) {
	rc.DeleteDocument(bookmarkId)
	//TODO 清理SUGGEST，根据推入的关键词计数器来清理，如果计数器为0，删除对应的建议
}
func SetDocIndex(userId string, bookmark bookmark.Bookmark) {
	if len(bookmark.Id) == 0 || len(bookmark.Url) == 0 {
		return
	}
	nameIndex := ""
	descIndex := ""
	urlIndex := ""
	if len(bookmark.Name) > 0 {
		nameSplit := word.Word2Index(bookmark.Name)
		nameIndex = strings.Join(nameSplit, " ")
		println(nameSplit, nameIndex)
		//添加建议
		for _, s := range nameSplit {
			ac.AddTerms(redisearch.Suggestion{Term: s})
		}
	}
	if len(bookmark.Desc) > 0 {
		descSplit := word.Word2Index(bookmark.Desc)
		descIndex = strings.Join(descSplit, " ")

		//添加建议
		for _, s := range descSplit {
			ac.AddTerms(redisearch.Suggestion{Term: s})
		}
	}
	if len(bookmark.Url) > 0 {
		urlSplit := word.Word2Index(bookmark.Url)
		urlIndex = strings.Join(urlSplit, " ")

		//添加建议
		for _, s := range urlSplit {
			ac.AddTerms(redisearch.Suggestion{Term: s})
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
func Search(userId, word string, offset, limit int) {
	if len(word) == 0 {
		return
	}
	docs, total, err := rc.Search(redisearch.NewQuery(word).
		Limit(offset, limit).
		SetReturnFields("id"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("search totoal:%v rp:%v", total, docs)
}
func Suggest(userId, word string, offset, limit int) {
	if len(word) == 0 {
		return
	}
	ss, err := ac.SuggestOpts(word, redisearch.DefaultSuggestOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("search  suggest:%v", ss)
}
