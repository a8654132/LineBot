package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	// defer cancel()
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			imageURL := "https://cdn.free.com.tw/blog/wp-content/uploads/2014/08/Placekitten480-g.jpg"
			template := linebot.NewButtonsTemplate(
					imageURL, "哈囉你好!", "我相信這次會成功的",
					linebot.NewURITemplateAction("來看看卡莉", "https://www.pixiv.net/member_illust.php?mode=medium&illust_id=62861397"),
					linebot.NewMessageTemplateAction("Say hello!", "你好"),
			)

			if _, err := bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewTemplateMessage("TEST", template)).Do(); err != nil {
				log.Println(err)
			}
		}
	}
}
