package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	channelSecret := "66fdd3246b5f37baae63be6676823ada"
	channelAccessToken := "Ui2NhUdEXK78K+AZcVhLRDrJplBlfLWBSKH+lAM075AdI2JBx1ZCp8bb9TtDR1X7xojouh88eAGJkKOPz7SXExqg4dJU2xJRNazF3MPy9maGiGCG3nXtQ6F7Crs7JJRYvNPY4rhdeaKjHClHMH6HygdB04t89/1O/w1cDnyilFU="

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
			if event.Source.Type == linebot.EventSourceTypeGroup {
				groupID := event.Source.GroupID
				log.Println("Group ID:", groupID)
	
				// ส่งข้อความกลับไปยังกลุ่ม (ถ้าต้องการ)
				_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Group ID: "+groupID)).Do()
				if err != nil {
					log.Println("Error replying message:", err)
				}
			}
		}
	})
	

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
