package discordAdapter

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/wormhole"
    "github.com/bwmarrin/discordgo"
)

func IgnorePrivateMessages(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    input := args[0].AsBox().(*discordgo.MessageCreate)
    session := args[1].AsBox().(*discordgo.Session)

    channel, err := session.Channel(input.ChannelID)
    if err != nil {
        return equinox.STOP_EXECUTION
    }

    if !channel.IsPrivate {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}

func IgnoreChannelMessages(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    ret := IgnorePrivateMessages(args...)

    if ret == equinox.STOP_EXECUTION {
        return equinox.CONTINUE_EXECUTION
    }

    return equinox.STOP_EXECUTION
}
