package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
)

type TelegramWebHook struct {
	ListenPort      string
	GroupsWhiteList []int
	BotName         string
	FordwardApiUrl  string
}

func (t *TelegramWebHook) isWhiteListed(id int) bool {
	for _, v := range t.GroupsWhiteList {
		if v == id {
			return true
		}
	}
	return false
}

func (t *TelegramWebHook) isBotMentioned(message string) bool {
	return strings.Contains(message, t.BotName)
}

func (t *TelegramWebHook) Handler(w http.ResponseWriter, r *http.Request) {
	var update, err = parseInput(r)
	if err != nil {
		log.Printf("error parsing input data: %s", err.Error())
		return
	}
	log.Printf("%d - Input: %s", update.Id, update)
	if update.Message == nil {
		log.Printf("%d - message is empty", update.Id)
		return
	}
	if !update.Message.Chat.isGroup() {
		log.Printf("%d - message is not coming from a group", update.Id)
		return
	}
	if !t.isWhiteListed(update.Message.Chat.Id) {
		log.Printf("%d - message from an unknown group: %d", update.Id, update.Message.Chat.Id)
		return
	}
	if !t.isBotMentioned(update.Message.Text) {
		log.Printf("%d - the Bot is not mentioned", update.Id)
		return
	}
	t.forwardMessage(update)
}

func parseInput(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		return nil, err
	}
	return &update, nil
}

func (t *TelegramWebHook) forwardMessage(update *Update) {
	var wg sync.WaitGroup
	for i := range t.GroupsWhiteList {
		if t.GroupsWhiteList[i] != update.Message.Chat.Id {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				var messageId int
				if update.Message.Reply == nil {
					messageId = update.Message.Id
				} else {
					messageId = update.Message.Reply.Id
				}

				forward := &ForwardMessage{
					TargetChatId: t.GroupsWhiteList[i],
					FromChatId:   update.Message.Chat.Id,
					MessageId:    messageId,
				}
				log.Printf("%d - calling forward message: %s", update.Id, forward)
				input, err := json.Marshal(forward)
				if err != nil {
					log.Printf("%d - error creating json input: %s", update.Id, err)
					return
				}
				resp, err := http.Post(t.FordwardApiUrl, "application/json", bytes.NewBuffer(input))
				if err != nil {
					log.Printf("%d - error calling telegram api: %s", update.Id, err)
					return
				}
				if resp.StatusCode == http.StatusOK {
					log.Printf("%d - telegram api called successfully: %s", update.Id, forward)
				} else {
					log.Printf("%d - telegram api response not valid: %d", update.Id, resp.StatusCode)
				}
			}(i)
		}
	}
	wg.Wait()
}
