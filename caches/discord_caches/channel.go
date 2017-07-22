package discord_caches

import (
    "code.lukas.moe/x/equinox/caches"
    "github.com/bwmarrin/discordgo"
)

func GetChannel(id string) (*discordgo.Channel, error) {
    ch := caches.Get(id)
    if ch != nil {
        return ch.(*discordgo.Channel), nil
    }

    ch, err := caches.Session().(*discordgo.Session).Channel(id)
    if err != nil {
        caches.Set(id, caches.NewItem(ch))
        return ch.(*discordgo.Channel), nil
    }

    return nil, err
}

func GetGuild(id string) (*discordgo.Guild, error) {
    ch := caches.Get(id)
    if ch != nil {
        return ch.(*discordgo.Guild), nil
    }

    ch, err := caches.Session().(*discordgo.Session).Channel(id)
    if err != nil {
        caches.Set(id, caches.NewItem(ch))
        return ch.(*discordgo.Guild), nil
    }

    return nil, err
}

func GetUser(id string) (*discordgo.User, error) {
    ch := caches.Get(id)
    if ch != nil {
        return ch.(*discordgo.User), nil
    }

    ch, err := caches.Session().(*discordgo.Session).User(id)
    if err != nil {
        caches.Set(id, caches.NewItem(ch))
        return ch.(*discordgo.User), nil
    }

    return nil, err
}