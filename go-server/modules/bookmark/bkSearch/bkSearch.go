package bkSsearch

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	. "pettyfox.top/bookmark/conf"
	"pettyfox.top/bookmark/modules/bookmark"
	"pettyfox.top/bookmark/modules/word"
	"strings"
	"time"
)

var RedisClient *redis.Pool

func Init() {
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,                 //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			c, err := redis.Dial(RedisSearchConf["type"], RedisSearchConf["address"])

			if err != nil {
				fmt.Printf("redis err:%v", err)
				println("redis search error", err)
				return nil, err
			}
			if RedisSearchConf["auth"] != "" {
				if _, err := c.Do("AUTH", RedisSearchConf["auth"]); err != nil {
					c.Close()
					fmt.Println("redis password error", err)
					return nil, err
				}
			}
			c.Do("select", RedisSearchConf["db"])

			fmt.Println("redis search ok")
			return c, nil
		},
	}
	initIndex()
}
func initIndex() {
	println("init redis search")
	rc := RedisClient.Get()
	defer rc.Close()
	rp, err := rc.Do("FT.CREATE",
		"bk_index_1", "ON", "HASH", "PREFIX", "1", "bk_doc", "SCHEMA",
		"name", "TEXT",
		"url", "TEXT",
		"desc", "TEXT",
	)
	if err != nil {
		fmt.Printf("redis search err:%v", err)
	}
	fmt.Printf("search rp:%v", rp)

}

func SetIndex(userId string, bookmark bookmark.Bookmark) {
	if len(bookmark.Id) == 0 || len(bookmark.Url) == 0 {
		return
	}
	rc := RedisClient.Get()
	defer rc.Close()
	nameIndex := ""
	descIndex := ""
	urlIndex := ""
	if len(bookmark.Name) > 0 {
		nameSplit := word.Word2Index(bookmark.Name)
		nameIndex = strings.Join(nameSplit, " ")
		println(nameSplit, nameIndex)
		//添加建议
		for _, s := range nameSplit {
			rc.Do("FT.SUGADD", "ac:"+userId, s, "1")
		}
	}
	if len(bookmark.Desc) > 0 {
		descSplit := word.Word2Index(bookmark.Desc)
		descIndex = strings.Join(descSplit, " ")

		//添加建议
		for _, s := range descSplit {
			rc.Do("FT.SUGADD", "ac:"+userId, s, "1")
		}
	}
	if len(bookmark.Url) > 0 {
		urlSplit := word.Word2Index(bookmark.Url)
		urlIndex = strings.Join(urlSplit, " ")

		//添加建议
		for _, s := range urlSplit {
			rc.Do("FT.SUGADD", "ac:"+userId, s, "1")
		}
	}
	//建立索引
	if len(nameIndex) > 0 {
		rp, err := rc.Do("HSET", "bk_doc_"+bookmark.Id, "name", nameIndex)
		fmt.Printf("set index:%v,%v", rp, err)
	}
	if len(urlIndex) > 0 {
		rc.Do("HSET", "bk_doc_"+bookmark.Id, "url", urlIndex)
	}
	if len(descIndex) > 0 {
		rc.Do("HSET", "bk_doc_"+bookmark.Id, "desc", descIndex)
	}

}
func Search(userId, word string, offset, limit int) {
	if len(word) == 0 {
		return
	}
	rc := RedisClient.Get()
	defer rc.Close()
	rp, err := redis.(rc.Do("FT.SEARCH", "bk_index_"+userId, word, "LIMIT", offset, limit))
	if err != nil {
		fmt.Printf("redis search err:%v", err)
	}

	fmt.Printf("search rp:%v", rp)
}
