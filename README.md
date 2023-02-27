# Telegram group manager bot in Golang

## Dependencies

* [Telego](https://github.com/mymmrac/telego)

## Commands

```
🕹 Available commands:
* /me: Gets the user presentation.
* /kickinactives: Removes all inactive users from group in a time period. Usage: /kickinactives <days>
* /setwelcome: Saves a new welcome message. Usage: /setwelcome <text>
* /setlang: Sets the language. Usage: /setlang <lang> (en: english, es: spanish)
* /start: Shows command usage.
* /welcome: Shows the welcome message.
* /stats: Shows user stats
* /inactives: Returns the list of inactive users in days period. Usage: /inactives <days>
* /setme: Sets the user presentation. Usage: /setme <text> adding a photo
* /help: Shows command usage.
```

## Environment variables

* `TELEGRAM_API_TOKEN`: Required. Telegram API token provided by BotFather.
* `SENTRY_DSN`: Optional. Sentry notification endpoint to catch exceptions.
* `DEBUG_JSON`: Optional. JSON memory repository to load.

## Execution

### Locally from code
```
$ TELEGRAM_API_TOKEN=xxxx SENTRY_DSN=xxxx go run main.go
```

### With Docker (recommended)

The `.env` file needs to contain the mentioned environment variables.

```
$ docker run --env-file .env ariel17/golang-telegram-group-manager:latest
```

## Deployment

Using terraform:

```
$ cd deployments
$ terraform apply
```
