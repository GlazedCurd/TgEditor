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
- [MarkdownV2](https://core.telegram.org/bots/api#markdownv2-style) <- markdown requires escape 
- [HTML](https://core.telegram.org/bots/api#html-style)
- [markdown](https://core.telegram.org/bots/api#markdown-style)

```--token``` - token from [bot father](https://core.telegram.org/bots/features#botfather) used to authorise TgEditor in telegram.

Usage example: 

Start bot in example folder from this repo:
```
➜  TgEditor git:(main) ✗ cd example
➜  example git:(main) ✗ ../TgEditor --token={YOUR_TELEGRAM_TOKEN} --mode=HTML
```

After start, write anything to your bot and, if all is Ok, you'll get menu with a list of all files in bot working directory.

![file select menu](/docs/images/file_select_menu.png?raw=true)

If you click on ```example.html``` (because of ```--mode=HTML```) and you'll get rendered content of chosen file, with menu.   

![file menu](/docs/images/file_menu.png?raw=true)

```Reload``` button signals bot to reload file and update message. 
```Switch to another file``` button returns you to the list of files in working directory. 

> Warning. Current version don't have any access restrictions. 