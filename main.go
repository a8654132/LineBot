package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "context"
	// "time"
	"encoding/json"

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
		if event.Type == linebot.EventTypeFollow {
			prof := event.Source.UserID

			if _, err := bot.PushMessage("Uecc089487f1487a78637be4e2fe3dca9", linebot.NewTextMessage("Hello, world\n"+prof)).Do(); err != nil {
					log.Print(err)
			}
		}

		if event.Type == linebot.EventTypeMessage {
			// fmt.Println(prof)
			binary, _ := Redis_Get(mac)
			user := new(USER_MAC)
			json.Unmarshal(binary,&user)
			var allcontent string
			for i:=0;i < len(user.USER) ; i++{
				allcontent = allcontent+user.USER[i].CONTENT
			}
			if _, err := bot.PushMessage("U496f42701ef55ea7f5e250508ba642f0", linebot.NewTextMessage(allcontent)).Do(); err != nil {
					log.Print(err)
			}
		}
	}





}
