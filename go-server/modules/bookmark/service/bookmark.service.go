package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	redis3 "github.com/garyburd/redigo/redis"
	"pettyfox.top/bookmark/modules/bookmark"
	bkSsearch "pettyfox.top/bookmark/modules/bookmark/bkSearch"
	"pettyfox.top/bookmark/modules/redis"
	"pettyfox.top/bookmark/modules/word"
)

var KEY = "bookmark::user1"
var userId = "1"

func Save(bookmark bookmark.Bookmark) {
	rc := redis.RedisClient.Get()
	defer rc.Close()
	has := md5.Sum([]byte(bookmark.Url))

	bookmark.Id = hex.EncodeToString(has[:])
	data, _ := json.Marshal(&bookmark)
	rc.Do("hset", KEY, bookmark.Id, data)

	saveCoverIndex(bookmark)

}
func saveIndex(bookmark bookmark.Bookmark) {
	bkSsearch.SetDocIndex(userId, bookmark)
}
func saveCoverIndex(bookmark bookmark.Bookmark) {
	bkSsearch.SetDocIndex(userId, bookmark)
}
func Remove(params bookmark.IdsParams) {
	rc := redis.RedisClient.Get()
	defer rc.Close()
	for i := 0; i < len(params.Ids); i++ {
		bkSsearch.RemoveDocIndex(params.Ids[i])
		rc.Do("hdel", KEY, params.Ids[i])
	}
}
func List() []bookmark.Bookmark {
	rc := redis.RedisClient.Get()
	defer rc.Close()
	list, err := redis3.Strings(rc.Do("hvals", KEY))
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	bookmarkList := make([]bookmark.Bookmark, 0)
	for i := 0; i < len(list); i++ {
		item := bookmark.Bookmark{}
		json.Unmarshal([]byte(list[i]), &item)
		bookmarkList = append(bookmarkList, item)
	}
	return bookmarkList
}
func GetIndex(id string) map[string][]string {
	rc := redis.RedisClient.Get()
	defer rc.Close()
	item := bookmark.Bookmark{}
	str, _ := redis3.String(rc.Do("hget", KEY, id))
	if len(str) > 0 {
		json.Unmarshal([]byte(str), &item)
	}
	if len(item.Id) == 0 {
		return nil
	}
	result := map[string][]string{}
	result["url"] = word.Word2Index(item.Url)
	result["desc"] = word.Word2Index(item.Desc)
	result["name"] = word.Word2Index(item.Name)
	return result
}
func Search(k string) []bookmark.Bookmark {
	ids := bkSsearch.Search(userId, k, 0, 20)
	if len(ids) == 0 {
		suggest := SearchSuggest(k)
		if len(suggest) > 0 {
			ids = bkSsearch.Search(userId, suggest[0], 0, 20)
		}
	}
	//fmt.Printf("query:%v", ids)
	bookmarkList := make([]bookmark.Bookmark, 0)
	if len(ids) > 0 {
		ids := uniqueArr(ids)
		rc := redis.RedisClient.Get()
		defer rc.Close()
		for _, id := range ids {
			item := bookmark.Bookmark{}
			str, err := redis3.String(rc.Do("hget", KEY, id))
			if err != nil {
				println("err", err)
			}
			if len(str) > 0 {
				json.Unmarshal([]byte(str), &item)
			}
			if len(item.Id) > 0 {
				bookmarkList = append(bookmarkList, item)
			}
		}
	}
	return bookmarkList
}
func ResetIndex() {
	bkSsearch.RestIndex()
	bookmarks := List()
	for _, b := range bookmarks {
		saveIndex(b)
	}
}
func SearchSuggest(keyword string) []string {
	return bkSsearch.Suggest(userId, keyword, 0, 20)
}
func uniqueArr(m []string) []string {
	d := make([]string, 0)
	tempMap := make(map[string]bool, len(m))
	for _, v := range m { // 以值作为键名
		if tempMap[v] == false {
			tempMap[v] = true
			d = append(d, v)
		}
	}
	return d
}
func mergeArr(a, b []string) []string {
	var arr []string
	for _, i := range a {
		if len(i) == 0 {
			continue
		}
		arr = append(arr, i)
	}
	for _, j := range b {
		if len(j) == 0 {
			continue
		}
		arr = append(arr, j)
	}
	return arr
}
