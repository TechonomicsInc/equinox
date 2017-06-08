package discordAdapter

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/wormhole"
    "github.com/bwmarrin/discordgo"
)

func IgnoreOtherBots(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    input := args[0].AsBox().(*discordgo.MessageCreate)

    if input.Author.Bot {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}

func IgnoreGroupMentions(args ...*wormhole.Wormhole) equinox.AdapterEvent {
    input := args[0].AsBox().(*discordgo.MessageCreate)

    if input.MentionEveryone {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}
