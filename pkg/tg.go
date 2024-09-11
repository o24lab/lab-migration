package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"pkg/errorx"
)

func TgSend(msg1 string) errorx.Error {
	// 替换成你自己的 Bot Token
	botToken := "7494712089:AAEKf9cy57lB3Dr_iy2MKvyglOd3cbUbDgI"

	// 替换成你要发送消息的聊天 ID
	chatID := int64(6289722448) // 这是聊天 ID，通常是你的用户 ID 或群组 ID

	// 创建一个新的 Bot 实例
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return errorx.WithStack(err)
	}
	// 创建一个新的消息
	msg := tgbotapi.NewMessage(chatID, msg1)

	// 发送消息
	_, err = bot.Send(msg)
	if err != nil {
		return errorx.WithStack(err)
	}
	return nil
}
