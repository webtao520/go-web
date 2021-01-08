package article

import (
	"github.com/gocolly/colly"
)

type Douban struct {

}

// 方法
func (douban Douban) Get(url string) map[string]interface{}{
  // https://book.douban.com/subject/35229199/?icn=index-topchart-subject
  // https://book.douban.comcomments/
  var sli []map[string]string 
 // subject:=make(map[string]string)
 c := colly.NewCollector()
  c.OnHTML("#time .comment-item", func(e *colly.HTMLElement) {
	s := make(map[string]string)
	s["content"] = e.ChildText(".comment-content")
	s["username"] = e.ChildText(".comment-info a")
	s["commenttime"] = e.ChildText(".comment-info .comment-time")
	sli = append(sli, s)
  })  
  c.Visit(url)
  ret:=make(map[string]interface{})
  ret["list"] = sli
  return ret
}
