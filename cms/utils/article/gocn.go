package article

import (
	"strings"

	"github.com/gocolly/colly"
)

type Gocn struct {

}

// 抓取内容
func (gocn Gocn) Get(url string) map[string]interface{} {
	ret:=make(map[string]interface{})
	c:=colly.NewCollector()
	// 注册函数。函数将在每个HTML上执行
	c.OnHTML("#main .title", func(e *colly.HTMLElement) {
		ret["title"] = strings.TrimSpace(e.Text)
	})

	c.Visit(url)
	return ret
}
