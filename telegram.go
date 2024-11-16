package tgram

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewErrorReporter(fromAppName, apiToken, chatID string) (func(err error) error, error) {
	messenger, err := NewMessenger(apiToken, chatID)
	if err != nil {
		return nil, err
	}
	return func(err error) error {
		return messenger(fmt.Sprintf("Error from %s: %v", fromAppName, err))
	}, nil
}

func NewMessenger(apiToken, chatID string) (func(msg string) error, error) {
	_chatID, err := strconv.Atoi(chatID)
	if err != nil {
		return nil, fmt.Errorf("could not convert YOUR_CHAT_ID var to int. %w", err)
	}
	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return nil, fmt.Errorf("unable to create bot instance. %w", err)
	}
	return func(message string) error {
		// Create a new message
		msg := tgbotapi.NewMessage(int64(_chatID), message)

		// Send the message
		_, err = bot.Send(msg)
		if err != nil {
			return fmt.Errorf("failed to send telgram message. %w", err)
		}
		return nil
	}, nil
}

func ReportErr(apperr error) error {
	// Replace with your bot's API token
	YOUR_BOT_API_TOKEN := os.Getenv("YOUR_BOT_API_TOKEN")
	if YOUR_BOT_API_TOKEN == "" {
		return errors.New("YOUR_BOT_API_TOKEN env var is empty")
	}
	botToken := YOUR_BOT_API_TOKEN
	// Replace with the chat ID you want to send a message to
	YOUR_CHAT_ID := os.Getenv("YOUR_CHAT_ID")
	if YOUR_CHAT_ID == "" {
		return errors.New("YOUR_CHAT_ID env var is empty")
	}

	_chatID, err := strconv.Atoi(YOUR_CHAT_ID)
	if err != nil {
		return fmt.Errorf("could not convert YOUR_CHAT_ID var to int. %w", err)
	}

	chatID := int64(_chatID)

	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return fmt.Errorf("unable to create bot instance. %w", err)
	}

	// Create a new message
	msg := tgbotapi.NewMessage(chatID, apperr.Error())

	// Send the message
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send telgram message. %w", err)
	}

	return nil
}
