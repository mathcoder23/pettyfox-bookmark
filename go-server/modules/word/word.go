package word

import (
	"github.com/mozillazg/go-pinyin"
	"github.com/wangbin/jiebago"
	"strings"
)

var seg jiebago.Segmenter

func init() {
	seg.LoadDictionary("dict.txt")
}
func Word2Index(word string) []string {
	rs := make([]string, 0)
	if len(word) == 0 {
		return rs
	}
	results := seg.CutForSearch(word, true)
	for item := range results {
		if len(strings.TrimSpace(item)) == 0 {
			continue
		}
		//println("text:", item)
		rs = append(rs, item)

		pinyinFullArgs := pinyin.NewArgs()
		pinyinFullArgs.Style = pinyin.Normal
		pinyinFirstArgs := pinyin.NewArgs()
		pinyinFirstArgs.Style = pinyin.FirstLetter

		firsts := pinyin.LazyPinyin(item, pinyinFirstArgs)
		if len(firsts) > 0 {
			rs = append(rs, strings.Join(firsts, ""))
		}
		fulls := pinyin.LazyPinyin(item, pinyinFullArgs)
		if len(fulls) > 0 {
			rs = append(rs, strings.Join(fulls, ""))
		}
	}
	return rs
}
