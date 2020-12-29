package handlers

import (
	"chitchat/config"
	"chitchat/models"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var logger *log.Logger
var localizer *i18n.Localizer

func init() {
	// 初始化 Localizer 以便被所有处理器方法使用
	localizer = i18n.NewLocalizer(config.ViperConfig.LocaleBundle, config.ViperConfig.App.Language)

	// 日志初始化
	file, err := os.OpenFile(config.ViperConfig.App.Log+"/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	/**
	func New(out io.Writer, prefix string, flag int) *Logger
	New创建一个Logger。参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。
	*/
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

// 检查用户是否已登录并拥有一个会话，如果不是err则不是nil
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 生成 HTML 模板 generateHTML(writer, threads, "layout", "auth.navbar", "index")
// template  包使用介绍 TODO
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s/%s.html", config.ViperConfig.App.Language, file))
	}
	//  generateHTML 方法中将这个函数通过 template.FuncMap 组装后再通过 Funcs 方法应用到视图模板中
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// 日期格式化
func formatDate(t time.Time) string {
	datetime := "2006-01-02 15:04:05"
	return t.Format(datetime)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

//记录日志信息
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

//异常处理统一重定向道错误页面
func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}
