> #### NOTE: This project is still HEAVILY work-in-progress.
> #### DO NOT USE IT IN YOUR PROJECTS UNTIL THIS WARNING IS REMOVED.
> #### THINGS **WILL** CHANGE RAPIDLY, FREQUENT AND WITHOUT ANNOUNCEMENTS.

# Equinox

Code that makes you think "oh that's neat"

## What's all this fuzz about?

Equinox is a framework that helps you keep your code lean and clean while creating chatbots with discordgo.
This is achieved by moving the hard stuff like caching, message parsing, event handling and friends into the background.
All that's left are you, your modules, beautiful code and your imagination.

## Awesome Features:

## Simplicity

Hate your million-lines long code base that handles the `MESSAGE_CREATE` event?

Here's what it looks like with equinox:

```go

func OnMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
    router.Handle(message.Message)
}
```

## Adapters

Equinox is built on, and deals with, adapters.
Adapters are little functions that can intercept the message parsing proccess.
This allows you to completely take over the parser stages or use the shipped stuff.

Example to ignore messages from other bots:
```go
router.RegisterAdapter(equinox.MESSAGE_PRE_ANALYZE, func(msg *discordgo.Message) equinox.AdapterEvent {
    if msg.Author.Bot {
        return equinox.STOP_EXECUTION
    }

    return equinox.CONTINUE_EXECUTION
})
```

or just use the preset

```go
router.RegisterAdapter(equinox.MESSAGE_PRE_ANALYZE, adapters.IgnoreOtherBots)
```

It's as simple and flexible as:
- Input gets in
- Do your stuff (whatever you want. equinox doesn't care.)
- Return your OK or abort the entire event chain

## Annotations

Did you know that Go doesn't support annotations?<br>
Equinox uses them anyway by shipping a little parser.

Just implement a `Meta()` function on your Plugin and get the most beautiful and flexible code ever:

```go
type Ping struct{}

func (p *Ping) Meta() string {
    return `
        @Name("Ping")
        @Purpose("Pings stuff")
        @Params("none")
        @Access("EVERYONE")
        @PrefixListeners(
            "ping",
            "p",
            "pong"
        )
    `;
}
```

Did you notice how we used annotations that are not present in this repo?<br>
That's a little demonstration of the way annotations work.

Only `@Listeners`, `@PrefixListeners` and `@MentionListeners` are catched by equinox.<br>
The rest is implemented through adapters (the things you read about some lines ago).

They are called `AnnotationHandler` and `RuntimeAdapter` here but essentially work the same way.

## Docs

...are coming soon. (tm).

The API is not stable so right now writing them wouldn't be a smart thing to do.

## HOW DOES THIS WORK?

![](https://media.giphy.com/media/12NUbkX6p4xOO4/giphy.gif)
