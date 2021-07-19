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
	sonicCli.Ingester.Push(KEY, "url", bookmark.Id, bookmark.Url, "")
	sonicCli.Ingester.FlushBucket(KEY, "url")

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
func Search(k string) {
	println(k)
	rs, err := sonicCli.Search.Suggest(KEY, "url", k,20)
	if err != nil {
	}
	println("rs", rs)
	for i := 0; i < len(rs); i++ {
		println("rs:", "a", len(rs[i]))
	}
}
