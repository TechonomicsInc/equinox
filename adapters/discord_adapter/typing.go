package discord_adapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/equinox/caches"
)

func StartTyping(args ...interface{}) equinox.AdapterEvent {
    session := caches.Get(caches.SESSION).(*discordgo.Session)

    session.ChannelTyping(args[0].(string))

    return equinox.CONTINUE_EXECUTION
}
