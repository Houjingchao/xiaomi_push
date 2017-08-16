package xiaomi_push

import (
	"testing"
	"github.com/Houjingchao/xiaomi_push"
	"fmt"
)

func TestXiaomiPush_TopicRegIdSubscrib(t *testing.T) {
	XiaomiPush := xiaomi_push.NewXiaomiPush("3bd6yo1EzmTPQG0/tNNGbA==", []string{"cc.komiko.mengxiaozhuapp"}, "")
	b, err := XiaomiPush.TopicRegIdSubscrib("1", "oISnLd1oHb71yw439fzv/DrCNLNwHrLCdiXdMMT4mFw=")
	fmt.Println(b)
	fmt.Println(err)
}

func TestXiaomiPush_TopicRegIdUnsubscirbe(t *testing.T) {
	XiaomiPush := xiaomi_push.NewXiaomiPush("3bd6yo1EzmTPQG0/tNNGbA==", []string{"cc.komiko.mengxiaozhuapp"}, "")
	b, err := XiaomiPush.TopicRegIdUnsubscirbe("1", "oISnLd1oHb71yw439fzv/DrCNLNwHrLCdiXdMMT4mFw=")
	fmt.Println(b)
	fmt.Println(err)
}

func TestXiaomiPush_TopicAllByRegid(t *testing.T) {
	XiaomiPush := xiaomi_push.NewXiaomiPush("3bd6yo1EzmTPQG0/tNNGbA==", []string{"cc.komiko.mengxiaozhuapp"}, "")
	b, err := XiaomiPush.TopicAllByRegid("oISnLd1oHb71yw439fzv/DrCNLNwHrLCdiXdMMT4mFw=")
	fmt.Println(b)
	fmt.Println(err)
}