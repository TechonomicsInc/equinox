package discord_adapter

import (
    "code.lukas.moe/x/equinox"
    "github.com/bwmarrin/discordgo"
)

func IgnoreOtherBots(args ...interface{}) equinox.AdapterEvent {
    input := args[0].(*discordgo.Message)

    if input.Author.Bot {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}

func IgnoreGroupMentions(args ...interface{}) equinox.AdapterEvent {
    input := args[0].(*discordgo.Message)

    if input.MentionEveryone {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}
