package caches

import (
    "github.com/bwmarrin/discordgo"
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
