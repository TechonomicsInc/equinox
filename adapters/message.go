/*
 * Copyright (C) 2017 Subliminal Apps
 *
 * Licensed under the EUPL, Version 1.1 only (the "Licence");
 *
 * You may not use this work except in compliance with the Licence.
 * You may obtain a copy of the Licence at:
 * <https://joinup.ec.europa.eu/software/page/eupl>
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Licence is distributed on an "AS IS" basis,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the Licence for the specific language governing permissions and limitations
 * under the Licence.
 */

package adapters

import (
    "code.lukas.moe/x/discordgo"
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