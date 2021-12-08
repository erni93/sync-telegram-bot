package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	telegramWebHook, err := buildTelegramWebHook()
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", telegramWebHook.Handler).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":"+telegramWebHook.ListenPort, router))
}

func buildTelegramWebHook() (*TelegramWebHook, error) {
	port := flag.String("port", "8222", "Application port")
	botName := flag.String("botName", "", "Bot name, example @SyncEngineerBot")
	groupsWhiteList := flag.String("whiteList", "", "Groups white list separated by comma, example -1001338476919,-1001548985922")
	token := flag.String("token", "", "Telegram token, for more info visit https://core.telegram.org/bots#6-botfather")
	flag.Parse()
	if err := existsDefaultParameters(*botName, *groupsWhiteList, *token); err != nil {
		return nil, err
	}
	groups, err := splitWhiteList(*groupsWhiteList)
	if err != nil {
		return nil, err
	}
	return &TelegramWebHook{
		ListenPort:      *port,
		GroupsWhiteList: groups,
		BotName:         *botName,
		FordwardApiUrl:  "https://api.telegram.org/bot" + *token + "/forwardMessage",
	}, nil
}

func existsDefaultParameters(parameters ...string) error {
	for _, parameter := range parameters {
		if parameter == "" {
			return errors.New("Some parameters are empty, Add --help for more info")
		}
	}
	return nil
}

func splitWhiteList(groupsWhiteList string) ([]int, error) {
	var ids []int
	for _, item := range strings.Split(groupsWhiteList, ",") {
		id, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			return nil, err
		}
		ids = append(ids, int(id))
	}
	if len(ids) > 0 {
		return ids, nil
	}
	return nil, errors.New("Invalid whiteList format")
}
