> #### NOTE: This project is still HEAVILY work-in-progress.
> #### DO NOT USE IT IN YOUR PROJECTS UNTIL THIS WARNING IS REMOVED.

# Equinox

A service-agnostic chatbot framework

## Intro

Equinox is a framework that helps you creating expressive, yet easy to use and fun to code chatbots.
This is done by moving the hard stuff like proper caching and message parsing into the background.
Everything that's left is you and your modules.

Equinox is also heavily event-based and allows you to hook into almost any program stage.<br>
Whether you want high-level abstractions or just need a "drop-in router" to shorten your command-parsing code.<br>
The only limit is your imagination.

## Chat-Services

This framework is "service-agnostic" but also "dev-supportive".<br>
That means that the core is truly generic but you can always rely on the service-specific "Adapters" when you need them.

Currently the only supported adapters are for [Slack](https://slack.com/) and [Discord](https://discordapp.com/).<br>
More adapters might be added in the future.

## The Router

The router is the heart of equinox.
It works like a normal event-dispatcher but doesn't require it's own coroutines or management functions.
It tries to follow KISS where possible and only runs code when explicitly told to.

## Adapters

Adapters are a central concept of equinox to allow the service-agnosticism.
Equinox itself only uses stdlib functions and relays boxed values to specialised adapters.

Example: You're using discord and want to ignore private messages

```go
import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/equinox/discordAdapter"
)

r := equinox.NewRouter()
r.RegisterAdapter(equinox.MESSAGE_ANALYZE, discordAdapter.IgnorePrivateMessages)
```

`RegisterAdapter()` takes the name of an equinox event on the left and the adapter function on the right.
As soon as this event is dispatched all listeners are called sequentially.

Adapters are forced to return one of the three possible status codes:

```go
const (
    CONTINUE_EXECUTION AdapterEvent = iota
    STOP_EXECUTION
    PANIC
)
```

These codes should be pretty self-explaining.
The first adapter that returns a `STOP_EXECUTION` code breaks the event chain and aborts all further operations.

## Adapter Events

You'll only meet those guys when implementing your own adapters.

`AdapterEvent` is an enum that tells the router to continue/stop execution.
You should always try to ignore "direct" `AdapterEvent` values and use the bound functions instead.
This way the error-handling and Event-Parsing stays in a central place.

The easiest usage form is to defer the corresponding panic-handler and call `AdapterEvent#Act()`.
This automates the reaction to the returned event and you only have to take care of errors.

```go
defer AdapterPanicHandler()
r.Dispatch(MY_COOL_EVENT, arg1, arg2, arg...).Act()
```

If you want to have some more control about the panics/errors there is `AdapterEvent#ShouldAbort()`.
Just wrap the Dispatch call into an if-clause and return as needed.

```go
if r.Dispatch(MY_COOL_EVENT, arg1, arg2, arg...).ShouldAbort() {
    return
}

// else do something
```

And last but not least you can always throw the Event at a switch clause.

```go
switch r.Dispatch(MY_COOL_EVENT, arg1, arg2, arg...) {
case CONTINUE_EXECUTION, NO_HANDLERS_REGISTERED:
 /* Continue execution */
 break

case STOP_EXECUTION:
 /* Kill execution */
 return

case PANIC:
 panic("wew i did not see that coming")

}
```

AdapterEvents might be added/removed/changed until equinox reaches a stable release so make sure to add a "default"
case that handles unknown Events.

## Message Prefixing

Usually chatbots require their own prefix to trigger executions.
Example:

```
!echo hello world
```

The prefix parsing is done by equinox using a special "Prefix Handler".
If your bot only support a single static prefix, creating this adapter is rather easy:

```
router.SetPrefixHandler(
    equinox.NewStaticPrefix("!"),
)
```

If your bot needs to support dynamic prefixes (eg configurable on a per-server level) you can easily implement your own
prefix handler. Just create a function that resolves the prefix and returns a wormhole to it.

Example for discord:

```go
router.SetPrefixHandler(func(args ...interface{})(interface{}){
    var prefix string

    // retrieve prefix from database //

    return prefix
})
```

Note that the prefix handler is called for every incoming message.
You should implement some kind of prefix caching if you don't want degrading performance and/or overloaded database
servers.
