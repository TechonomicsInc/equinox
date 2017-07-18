package discord_adapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/equinox/caches"
)

func IgnorePrivateMessages(args ...interface{}) equinox.AdapterEvent {
    session := caches.Get(caches.SESSION).(*discordgo.Session)
    input := args[0].(*discordgo.MessageCreate)

    channel, err := session.State.Channel(input.ChannelID)
    if err != nil {
        return equinox.STOP_EXECUTION
    }

    if !channel.IsPrivate {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}

func IgnoreChannelMessages(args ...interface{}) equinox.AdapterEvent {
    ret := IgnorePrivateMessages(args...)

    if ret == equinox.STOP_EXECUTION {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}
