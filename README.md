> #### NOTE: This project is still HEAVILY work-in-progress.
> #### DO NOT USE IT IN YOUR PROJECTS UNTIL THIS WARNING IS REMOVED.
> #### THINGS **WILL** CHANGE RAPIDLY, FREQUENT AND WITHOUT ANNOUNCEMENTS.

# Equinox

Code that makes you think "oh that's neat"

## What's all this fuzz about?

Equinox is a framework that helps you to keep your code lean and clean while creating chatbots with discordgo.
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
This allows you to completely take over the parser stages or use the shipped stuff as a "drop-in core replacement".
It's all up to you.

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
Equinox uses them anyway by re-implementing them with a little parser.<br>

Equinox annotations look a bit like Java's but only map to vararg funcs.<br>
Arrays or Maps are not supported to keep the parser small.

To use them just implement a `Meta()` function on your Plugin and get the most beautiful and flexible code you've ever seen:

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
            "pong",
        )
    `;
}
```

Did you notice how we used annotations that are not present in this repo?<br>
That's a little demonstration of the way annotations work.

Only `@Listeners`, `@PrefixListeners` and `@MentionListeners` are catched by equinox.<br>
The rest is implemented through adapters (the things you read about some lines ago).

They are called `AnnotationHandler` and `RuntimeAdapter` here but essentially work the same way.

The parsing stage of annotation-processing is also highly customizable.
Once the parser encounters an unknown annotation it checks wether someone registered a handler for that.
If a handler is found, the function gets called with the matching paramters and *can* (but isn't required to) register event handlers that fire once this plugin handles a message. This is especially useful for modifying logic like `@Access(OWNER)` or `@RequiredLevel(12)` or something else.

## Automated Listeners w. Parameter Expansion

Ever needed plugins that work with `!name` and `@bot name`?<br>
Equinox has multiple ways to solve that!

Either create an annotation for each type:
```go
@PrefixListeners("name")
@MentionListeners("name")
```

Or create two generic listeners:
```go
@Listeners("{p}name", "{@}name")
```

Or use parameter expansion:
```go
@Listeners("{p,@}name")
```

The thing in `{}` is called "equinox expression" and supports more than just `@` and `p`.<br>
A documentation on that will follow later.<br>
Also, as you see, expressions support bash-like expansion with commas.

## The Last Resort

While Adapters and Annotations are extremely powerful and flexible tools that can solve *almost* all problems you'll ever face while writing a bot, there's one case where they can't help you: Cleverbot.

Or to be precise:<br>
"Doing stuff when the message looks like an `@mention` command but isn't".

This makes the last resort (which is technically nothing more than a fancy adapter) perfect to implement bot-apis like cleverbot that reply with messages when tagged instead of "no command".

Implementation Example:
```go
router.RegisterAdapter(equinox.LAST_RESORT_PRE_EXECUTE, func(msg *discordgo.Message) equinox.AdapterEvent {
    // Resolve names in received message
    m := msg.ContentWithMentionsReplaced()
    
    // Get a reply from your bot-api
    response := SendMessageToYourBotApi(m)
    
    // Send it back to discord
    caches.Session().ChannelMessageSend(
        msg.ChannelID,
        response,
    )

    return equinox.CONTINUE_EXECUTION
})
```

## Docs

...are coming soon. (tm).

The API is not stable so right now writing them wouldn't be a smart thing to do.

## HOW DOES THIS WORK?

![](https://media.giphy.com/media/12NUbkX6p4xOO4/giphy.gif)
