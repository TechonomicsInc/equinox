package adapters

import (
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/equinox/caches"
)

func StartTyping(msg *discordgo.Message) equinox.AdapterEvent {
    caches.Session().ChannelTyping(msg.ChannelID)

    return equinox.CONTINUE_EXECUTION
}