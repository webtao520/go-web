// 只有一个消息提示文本，那就是访问的群组不存在时返回的 Cannot read thread
// 通过 go-i18n 自动生成翻译文件

package main

import "github.com/nicksnyder/go-i18n/v2/i18n"

var messages = []i18n.Message{
	i18n.Message{
		ID:          "thread_not_found",        // ID 是消息文本的唯一标识
		Description: "Thread not exists in db", //
		Other:       "Cannot read thread",      // Other 则是对应的翻译字符串（默认是英文）
	},
}
