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
			imageURL := "https://i.pximg.net/c/600x600/img-master/img/2017/05/05/20/44/05/62754624_p0_master1200.jpg"
			template := linebot.NewButtonsTemplate(
					imageURL, "A", "B",
					linebot.NewURITemplateAction("來看看卡莉", "https://i.pximg.net/c/600x600/img-master/img/2017/05/05/20/44/05/62754624_p0_master1200.jpg"),
					linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
			)

			if _, err := bot.ReplyMessage(
				event.ReplyToken,
				linebot.NewTemplateMessage("TEST", template)).Do(); err != nil {
				log.Println(err)
			}
		}
	}
}
