# Sync-telegram-bot
Telegram bot that allows you to forward messages between different channels

# How it works
The application has different parameters

- **port**: Application port
- **botname**: Bot name, example @SyncEngineerBot, name used to detect if the bot was mentioned, this parameter and the bot name should be the same.
- **whitelist**: Groups white list separated by comma, example -1001338476919,-1001548985922, the bot will only be listening in those groups.
- **token**: Telegram token, for more info visit https://core.telegram.org/bots#6-botfather

The bot will forward the message to the different groups defined on the whitelist.

To trigger the bot, it has to be mentioned. If the message has a reply the bot will forward the first message instead.

Example, Group A

![demo2](https://user-images.githubusercontent.com/16189689/147395288-0f36cee4-35af-4aa8-a522-7b82287c73f1.png)

Group B

![demo1](https://user-images.githubusercontent.com/16189689/147395287-b80e036b-44c5-4319-be0f-d93924010a15.png)

# Docker image

There is a docker image available, https://hub.docker.com/repository/docker/erni93/sync-telegram-bot

Example

```
docker run -p 4040:4040 sync-telegram-bot -botname "@SyncEngineerBot" -port 4040 -token "XXXX Your telegram token XXXX" -whitelist "-693863188,-704904670
```

It will print out logs about any action and error

```
2021/12/25 23:07:36 Application listening on port 8222
2021/12/25 23:09:10 479362786 - Input: (id: 479362786, message: (id: 872202, text: el bot ahora reenvia el mensaje en el que se le cita si no esta respondiendo a otro @SyncEngineerBot , pd: feliz año nuevo chino, chat: (id: -1001338476919,  type:  supergroup, title:  Offtopic de Programación 🐙),  reply: <nil>))
2021/12/25 23:09:10 479362786 - calling forward message: (TargetChatId: -1001548985922, FromChatId: -1001338476919,  MessageId: 872202,)
2021/12/25 23:09:10 479362786 - telegram api response not valid: 400
```
