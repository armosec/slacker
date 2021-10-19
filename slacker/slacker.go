package slacker

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackBot struct {
	DefaultChannelID string
	token            string
	api              *slack.Client
}

func SlackBotInit(channelID, token string, isDebugMode bool) (*SlackBot, error) {
	if len(token) < 57 {
		return nil, fmt.Errorf("invalid token length")
	}

	bot := &SlackBot{
		DefaultChannelID: channelID,
		token:            token,
	}
	bot.api = slack.New(bot.token, slack.OptionDebug(isDebugMode))
	if channelID == "" {
		return bot, nil
	}
	if err := bot.JoinChannel(bot.DefaultChannelID); err != nil {
		return nil, fmt.Errorf("unable to join default channel %s due to %v", channelID, err.Error())
	}
	return bot, nil
}

func (bot *SlackBot) JoinChannel(channelID string) error {
	if channelID == "" {
		channelID = bot.DefaultChannelID
	}

	_, _, _, err := bot.api.JoinConversation(channelID)
	return err
}

func (bot *SlackBot) SendMessage(level, channelID, message string) error {
	if channelID == "" {
		channelID = bot.DefaultChannelID
	} else {
		bot.JoinChannel(channelID)
	}
	_, _, err := bot.api.PostMessage(channelID, slack.MsgOptionText(fmt.Sprintf("[%s]: %s", level, message), false))
	return err
}

func (bot *SlackBot) SendErrorMessage(channelID, message string) error {
	return bot.SendMessage("ERROR", channelID, message)
}

func (bot *SlackBot) SendCriticalMessage(channelID, message string) error {
	return bot.SendMessage("CRITICAL", channelID, message)
}

func (bot *SlackBot) SendINFOMessage(channelID, message string) error {
	return bot.SendMessage("INFO", channelID, message)
}
