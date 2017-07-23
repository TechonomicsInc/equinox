package adapters

import (
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/equinox/caches"
    "code.lukas.moe/x/equinox"
)

func IsPrivateMessage(msg *discordgo.Message) bool {
    ch, err := caches.GetChannel(msg.ChannelID)
    if err != nil {
        return false
    }

    return ch.Type == discordgo.ChannelTypeDM || ch.Type == discordgo.ChannelTypeGroupDM
}

func IgnorePrivateMessages(msg *discordgo.Message) equinox.AdapterEvent {
    if IsPrivateMessage(msg) {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}

func IgnorePublicMessages(msg *discordgo.Message) equinox.AdapterEvent {
    if !IsPrivateMessage(msg) {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}

func IgnoreOtherBots(msg *discordgo.Message) equinox.AdapterEvent {
    if msg.Author.Bot {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}

func IgnoreGroupMentions(msg *discordgo.Message) equinox.AdapterEvent {
    if msg.MentionEveryone {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
}