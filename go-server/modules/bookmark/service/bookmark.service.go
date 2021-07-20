package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	redis3 "github.com/garyburd/redigo/redis"
	"pettyfox.top/bookmark/modules/bookmark"
	"pettyfox.top/bookmark/modules/redis"
	"pettyfox.top/bookmark/modules/sonicCli"
)

var KEY = "bookmark::user1"

func Save(bookmark bookmark.Bookmark) {
	rc := redis.RedisClient.Get()
	defer rc.Close()
	has := md5.Sum([]byte(bookmark.Url))

	bookmark.Id = hex.EncodeToString(has[:])
	data, _ := json.Marshal(&bookmark)
	rc.Do("hset", KEY, bookmark.Id, data)

	if len(bookmark.Url) > 0 {
		sonicCli.Push(KEY, "url", bookmark.Id, bookmark.Url)
	}
	if len(bookmark.Name) > 0 {
		sonicCli.Push(KEY, "name", bookmark.Id, bookmark.Name)
	}
	if len(bookmark.Desc) > 0 {
		sonicCli.Push(KEY, "url", bookmark.Id, bookmark.Desc)
	}

}
func Remove(params bookmark.IdsParams) {
	rc := redis.RedisClient.Get()
	defer rc.Close()
	for i := 0; i < len(params.Ids); i++ {
		rc.Do("hdel", KEY, params.Ids[i])
	}
}
func List() []bookmark.Bookmark {
	rc := redis.RedisClient.Get()
	defer rc.Close()
	list, err := redis3.Strings(rc.Do("hvals", KEY))
	if err != nil {

	}
	bookmarkList := make([]bookmark.Bookmark, 0)
	for i := 0; i < len(list); i++ {
		item := bookmark.Bookmark{}
		json.Unmarshal([]byte(list[i]), &item)
		bookmarkList = append(bookmarkList, item)
	}
	return bookmarkList
}
func Search(k string) []bookmark.Bookmark {
	a1 := sonicCli.QUERY(KEY, "url", k)
	a2 := sonicCli.QUERY(KEY, "desc", k)
	a3 := sonicCli.QUERY(KEY, "name", k)
	ids := mergeArr(a1, mergeArr(a2, a3))
	bookmarkList := make([]bookmark.Bookmark, 0)
	if len(ids) > 0 {
		ids := uniqueArr(ids)
		rc := redis.RedisClient.Get()
		defer rc.Close()
		for _, id := range ids {
			item := bookmark.Bookmark{}
			str, _ := redis3.String(rc.Do("hget", KEY, id))
			if len(str) > 0 {
				json.Unmarshal([]byte(str), &item)
			}
			bookmarkList = append(bookmarkList, item)
		}
	}
	return bookmarkList
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
		arr = append(arr, i)
	}
	for _, j := range b {
		arr = append(arr, j)
	}
	return arr
}
