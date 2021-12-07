package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	telegramWebHook := &TelegramWebHook{
		ProcessedUpdates: NewProcessedUpdates(),
		GroupsWhiteList: []int{
			-1001338476919, // Offtopic de Programación
			-1001548985922, // Offtopic sin censura
		},
		BotName:        "@SyncEngineerBot",
		FordwardApiUrl: "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/forwardMessage",
	}
	router := mux.NewRouter()
	router.HandleFunc("/", telegramWebHook.Handler).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8222", router))
}
