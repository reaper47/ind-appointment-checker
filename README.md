# IND Appointment Checker

Do you spend too much time refreshing IND biometrics and/or residence sticker 
appointments instead of watching cat videos on TikTok? Have you ever wished 
someone did this heavy work for you? 

Well, my friend, look no further because this program is exactly what you need!
Get notified when an earlier IND appointment pops up.

## Telegram

You need a Telegram bot to receive notifications. Otherwise, you are out of 
luck, and better off finding another solution. 

Follow these [steps](https://tutorial.cytron.io/2021/09/01/how-to-create-a-telegram-bot-get-the-api-key-and-chat-id/) to set up a bot
on Telegram.

Paste the Telegram bot API key and chat_id in the .env file.

## Installation

1. Download the latest [release](https://github.com/reaper47/ind-appointment-checker/releases/tag/v1.1.1) and unzip it.
1. Open the .env file to adjust the variables:

    - `TELEGRAM_CHATID`: the chat_id number from the Telegram section above
    - `TELEGRAM_BOTID`: the bot API key from the Telegram section above
    - `IND_START_DATE`: the start date for the search. For example, if start date is 15/05/2022, the program will notify you for dates after 15/05/2022. The date will be automatically set to today's if empty.
    - `IND_CURRENT_APPOINTMENT_BIOMETRICS`: your current biometrics appointment date in the form dd/mm/yyyy
    - `IND_CURRENT_APPOINTMENT_RESIDENCE_STICKER`: your current residence sticker appointment date in the form dd/mm/yyyy

3. Start the program by double-clicking on it or executing `./ind`

## Build from source

You do not trust the release binaries? No problem. We've got you covered. 
Follow the following steps:

1. Clone the repository

```bash
git clone https://www.github.com/reaper47/ind-appointment-checker.git
```

2. Fill the correct values in the ./bin/.env file
```
nano ./bin/.env
```

3. Build the program
```bash
make build
```

4. Start the program
```bash
./bin/ind
```

## Feedback

If you have any feedback, please reach out to me at macpoule@gmail.com

## Support

For support, email macpoule@gmail.com, open an issue or start a discussion on GitHub.
