# TgEditor

This bot created to make fast previews for telegramm messages. Telegramm messages are supposed to be stored local on your computer as a files with markup. 

Bot starts in directory with files. 

```
./TgEditor --help
    Usage of ./TgEditor:
      --mode string    telegram parse mode
      --token string   telegramm token
```

```--mode``` - value of the [parse_mode](https://core.telegram.org/bots/api#sendmessage) field that defines the way that telegram uses to render a message.  

Telegram supports several markup languages:
- [MarkdownV2](https://core.telegram.org/bots/api#markdownv2-style)
- [HTML](https://core.telegram.org/bots/api#html-style)
- [markdown](https://core.telegram.org/bots/api#markdown-style)

```--token``` - token from [bot father](https://core.telegram.org/bots/features#botfather) used to authorise TgEditor in telegram.

Usage example: 

Start bot in example folder from this repo:
```
example git:(main) ✗ ../TgEditor --token={YOUR_TELEGRAM_TOKEN} --mode=HTML
```


