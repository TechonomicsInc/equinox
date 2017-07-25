> #### NOTE: This project is still HEAVILY work-in-progress.
> #### DO NOT USE IT IN YOUR PROJECTS UNTIL THIS WARNING IS REMOVED.
> #### THINGS **WILL** CHANGE RAPIDLY, FREQUENT AND WITHOUT ANNOUNCEMENTS.

# Equinox

Code that makes you think "oh that's neat"

## What's all this fuzz about?

Equinox is a framework that helps you keep your code lean and clean while creating chatbots with discordgo.
This is achieved by moving the hard stuff like caching, message parsing, event handling and friends into the background.
All that's left are you, your modules and your imagination.

## Docs

...are coming soon. (tm).

Until then you're left with some example code below.

## Example

When the discord api is ready create a new router and configure it:

```go
// Create router
router := equinox.NewRouter()

// Ignore DMs, other bots and indirect mentions
// (The adapters package ships with equinox)
router.RegisterAdapter(equinox.MESSAGE_PRE_ANALYZE, adapters.IgnoreOtherBots)
router.RegisterAdapter(equinox.MESSAGE_PRE_ANALYZE, adapters.IgnorePrivateMessages)
router.RegisterAdapter(equinox.MESSAGE_PRE_ANALYZE, adapters.IgnoreGroupMentions)

// Simulate typing when executing commands
router.RegisterAdapter(equinox.HANDLER_PRE_EXECUTE, adapters.StartTyping)
router.RegisterAdapter(equinox.LAST_RESORT_PRE_EXECUTE, adapters.StartTyping)

// Register routes (aka initialize plugins)
router.AddRoutes(plugins.PluginList)
```

Then pass messages to equinox instead of doing stuff yourself

```go

func OnMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
    router.Handle(message.Message)
}
```

![](https://media.giphy.com/media/12NUbkX6p4xOO4/giphy.gif)
