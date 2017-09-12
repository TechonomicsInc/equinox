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
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/equinox/caches"
)

func StartTyping(msg *discordgo.Message) equinox.AdapterEvent {
    caches.Session().ChannelTyping(msg.ChannelID)

    return equinox.CONTINUE_EXECUTION
}