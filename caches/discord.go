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

package caches

import (
    "code.lukas.moe/x/discordgo"
)

func GetChannel(id string) (*discordgo.Channel, error) {
    cid := "discord::channel::" + id

    ch := Get(cid)
    if ch != nil {
        return ch.(*discordgo.Channel), nil
    }

    ch, err := Session().Channel(id)
    if err == nil {
        Set(cid, NewItem(ch))
        return ch.(*discordgo.Channel), nil
    }

    return nil, err
}

func GetGuild(id string) (*discordgo.Guild, error) {
    cid := "discord::guild::" + id

    ch := Get(cid)
    if ch != nil {
        return ch.(*discordgo.Guild), nil
    }

    ch, err := Session().Guild(id)
    if err == nil {
        Set(cid, NewItem(ch))
        return ch.(*discordgo.Guild), nil
    }

    return nil, err
}

func GetUser(id string) (*discordgo.User, error) {
    cid := "discord::user::" + id

    ch := Get(cid)
    if ch != nil {
        return ch.(*discordgo.User), nil
    }

    ch, err := Session().User(id)
    if err == nil {
        Set(cid, NewItem(ch))
        return ch.(*discordgo.User), nil
    }

    return nil, err
}
