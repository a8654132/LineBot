
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"io/ioutil"

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
	cmd := exec.Command("wget", "-N", "http://140.115.153.185/file/test.txt")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	tmp, err:= ioutil.ReadFile("test.txt")
	content := string(tmp)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("8769")).Do(); err != nil {
					log.Print(err)
			}
		}
	}

// 	for _, event := range events {
// 		if event.Type == linebot.EventTypeMessage {
// 			switch message := event.Message.(type) {
// 			case *linebot.TextMessage:
// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+"")).Do(); err != nil {
// 					log.Print(err)
// 				}
// 			}
// 		}
// 	}
}
