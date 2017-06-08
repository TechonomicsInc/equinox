package discordAdapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/bwmarrin/discordgo"
)

func IgnorePrivateMessages(args ...interface{}) equinox.AdapterEvent {
    input := args[0].(*discordgo.MessageCreate)
    session := args[1].(*discordgo.Session)

    channel, err := session.Channel(input.ChannelID)
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
