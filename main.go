package main

import (
	"bufio"
	"cars_telegram_bot/Admin"
	api "cars_telegram_bot/handleAPI"
	"cars_telegram_bot/user"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(api.GetApiToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	updates := bot.GetUpdatesChan(u)

	go receiveUpdates(ctx, updates)

	// Wait for a newline symbol, then cancel handling updates
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	cancel()
}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {

	adminStatus := Admin.CheckForAdmin(update.SentFrom().ID)
	switch {
	// Handle messages
	case update.Message != nil && adminStatus:
		Admin.HandleAdminMessage(update.Message)
		break

	case update.Message != nil && !adminStatus:
		user.HandleMessage(update.Message)
		break

	// Handle button clicks
	case update.CallbackQuery != nil && adminStatus:
		Admin.HandleAdminQuery(update.CallbackQuery)
		break

	case update.CallbackQuery != nil && !adminStatus:
		user.HandleButton(update.CallbackQuery)
		break
	}

}
