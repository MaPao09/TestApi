package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	channelSecret := "2ba2a0fa740b0e221247e2de5b7a1e18"
	channelAccessToken := "bwNTU+jdpfCtMwhIjlLnFQtGoBPG9AKu/9zIy61Vnj0ULeT/qr3VsvAyv5KP3mo6in2efiQbGjfJqdHs6aGXmMikE79y94n5n5AxRCAmG4mOXrubNmNxB//uN3Gaaep0vjCbqFgb1bJMy+/QZkoPdwdB04t89/1O/w1cDnyilFU="

	bot, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				http.Error(w, "Invalid Signature", http.StatusBadRequest)
			} else {
				http.Error(w, "Bad Request", http.StatusBadRequest)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("คุณพิมพ์ว่า: "+message.Text)).Do()
					if err != nil {
						log.Println("Error replying message:", err)
					}
				}
			}
		}
	})

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
