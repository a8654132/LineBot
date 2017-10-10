package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var mess = &Messenger{}

func main() {
	port := os.Getenv("PORT")
	log.Println("Server start in port:", port)
	mess.VerifyToken = os.Getenv("TOKEN")
	mess.AccessToken = os.Getenv("TOKEN")
	log.Println("Bot start in token:", mess.VerifyToken)
	mess.MessageReceived = MessageReceived
	http.HandleFunc("/webhook", mess.Handler)

	binary, _ := Redis_Get(mac)
	user := new(USER_MAC)
	json.Unmarshal(binary,&user)
	onlyonecontent := user.USER[1].CONTENT

	for i:=0 ; i<5 ; i++ {
  	mess.SendSimpleMessage("1460870680701162", fmt.Sprintf("第%d次主動傳送訊息",i+1))
	}

	mess.SendSimpleMessage("1460870680701162", fmt.Sprintf(onlyonecontent))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//MessageReceived :Callback to handle when message received.
func MessageReceived(event Event, opts MessageOpts, msg ReceivedMessage) {
	// log.Println("event:", event, " opt:", opts, " msg:", msg)
	resp, err := mess.SendSimpleMessage(opts.Sender.ID, fmt.Sprintf("你好,現在是被動的,%s,%s", opts.Sender.ID ,msg.Text))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}
