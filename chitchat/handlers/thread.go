package handlers

import (
	"chitchat/models"
	_ "chitchat/models"
	"net/http"
)

// GET / /threads/new
// 创建主题页面
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "new.thread")
	}
}

// POST  / /thread/create
// 执行主题创建逻辑
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET / thread/read
// 通过 ID 渲染指定全组页面
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, _ := models.ThreadByUUID(uuid)
	// TODO 本地化
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, &thread, "layout", "navbar", "thread")
	} else {
		generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
	}
}
