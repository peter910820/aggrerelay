package model

type Bot struct {
	PlatForm string
	BotID    string
	Config   PlatFormConfig
}

type PlatFormConfig interface {
	InitPlatFormConfig() error
}

type DiscordConfig struct {
	Token string
	AppID string
}

func (d *DiscordConfig) InitPlatFormConfig() error {
	return nil
}

type LineConfig struct {
	Secret  string
	Channel string
}

func (l *LineConfig) InitPlatFormConfig() error {
	return nil
}

type BotMessage struct {
	BotID    string
	PlatForm string
	Message  string
}
