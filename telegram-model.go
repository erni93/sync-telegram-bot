package main

import "fmt"

type Update struct {
	Id      int      `json:"update_id"`
	Message *Message `json:"message"`
}

func (u Update) String() string {
	return fmt.Sprintf("(id: %d, message: %s)", u.Id, u.Message)
}

type Message struct {
	Id    int      `json:"message_id"`
	Text  string   `json:"text"`
	Chat  Chat     `json:"chat"`
	Reply *Message `json:"reply_to_message"`
}

func (m Message) String() string {
	return fmt.Sprintf("(id: %d, text: %s, chat: %s,  reply: %s)", m.Id, m.Text, m.Chat, m.Reply)
}

type Chat struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Title string `json:"title"`
}

func (c Chat) isGroup() bool {
	return c.Type == "group" || c.Type == "supergroup"
}

func (c Chat) String() string {
	return fmt.Sprintf("(id: %d,  type:  %s, title:  %s)", c.Id, c.Type, c.Title)
}

type ForwardMessage struct {
	TargetChatId int `json:"chat_id"`
	FromChatId   int `json:"from_chat_id"`
	MessageId    int `json:"message_id"`
}

func (f ForwardMessage) String() string {
	return fmt.Sprintf("(TargetChatId: %d, FromChatId: %d,  MessageId: %d,)", f.TargetChatId, f.FromChatId, f.MessageId)
}
