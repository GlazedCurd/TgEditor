# TgEditor

This bot created to make fast previews for telegramm messages. Telegramm messages are supposed to be stored local on your computer as a files with markup. 

Bot starts in directory with files. 

```
./TgEditor --help
Usage of ./TgEditor:
      --debug          enable debug
      --mode string    telegram parse mode (default "HTML")
      --path string    working folder for the bot (default ".")
      --token string   telegramm token
      --users string   users who are allowed to use this bot separated by semicolumns
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
➜  TgEditor git:(main) ✗ ../TgEditor --token={YOUR_TELEGRAM_TOKEN} --mode=HTML --users={YOUR_TELEGRAM_LOGIN} --path=./example
```

After start, write anything to your bot and, if all is Ok, you'll get menu with a list of all files in bot working directory.

![file select menu](/docs/images/file_select_menu.png?raw=true)

If you click on ```example.html``` (because of ```--mode=HTML```) and you'll get rendered content of chosen file, with menu.   

![file menu](/docs/images/file_menu.png?raw=true)

```Reload``` button signals bot to reload file and update message. 
```Switch to another file``` button returns you to the list of files in working directory. 
