package relay

import (
	"github.com/bwmarrin/discordgo"

	"aggrerelay/model"
)

func DiscordRelay(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	allowedChannels := map[string]bool{
		"123456789012345678": true, // 頻道 ID
		"987654321098765432": true,
	}

	if !allowedChannels[m.ChannelID] {
		return
	}

	// send msg to channel
	model.RelayChan <- msg
}
