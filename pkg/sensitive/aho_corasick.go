package sensitive

import (
	"MyIM/model/dao"
	"MyIM/model/db"
	"context"
	"github.com/BobuSumisu/aho-corasick"
	"regexp"
)

var ac *ahocorasick.Trie

func Init() {
	// 读取敏感词库
	LoadSensitiveWords()
}

func LoadSensitiveWords() {
	dictDao := dao.DictSensitiveWords{}
	list, err := dictDao.GetList(context.Background(), db.GetDB())
	if err != nil {
		panic(err)
	}
	ac = ahocorasick.NewTrieBuilder().AddStrings(list).Build()
}

func Match(ctx context.Context, content string) []*ahocorasick.Match {
	return ac.MatchString(content)
}

func FilterCharacters(input string) string {
	// 定义正则表达式匹配英文、数字和中文数字
	re := regexp.MustCompile("[a-zA-Z0-9一二三四五六七八九壹贰叁肆伍陆柒捌玖]+")

	// 使用正则表达式进行匹配和替换
	filtered := re.ReplaceAllStringFunc(input, func(s string) string {
		// 将匹配到的字符替换为*
		return "*" // 这里可以根据需要指定其他替换逻辑
	})

	return filtered
}
