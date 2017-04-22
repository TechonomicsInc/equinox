/*

Equinox - A servcice-agnostic chatbot framework

Intro

Equinox is a framework that helps you creating expressive, easy to use and fun to code chatbots.
This is done by moving the hard stuff like proper caching and message parsing into the background.
Everything that's left is you and your modules.

The Router

The router is the heart of equinox.
It works like a normal event-dispatcher but doesn't require it's own coroutines or management functions.
It tries to follow KISS where possible and only runs code when explicitly told to.

Understanding Wormholes

Wormholes are a combination of mutexes and empty interfaces that allow a simple kind of generic programming in go
without code generation.

Take a look at https://godoc.org/code.lukas.moe/x/wormhole for more information.

Understanding Adapters

Adapters are a central concept of equinox to allow the service-agnosticism.
Equinox itself only uses stdlib functions and relays Wormholes to specialised adapters.

Example: You're using discord and want to ignore private messages

    import (
        "code.lukas.moe/x/equinox"
        "code.lukas.moe/x/equinox/discordAdapter"
    )

    r := equinox.NewRouter()
    r.RegisterAdapter(equinox.MESSAGE_ANALYZE, discordAdapter.IgnorePrivateMessages)

RegisterAdapter() takes the name of an equinox event on the left and the adapter function on the right.
As soon as this event is dispatched all listeners are called sequentially.

Adapters are forced to return one of the three possible status codes:

    const (
        CONTINUE_EXECUTION AdapterEvent = iota
        STOP_EXECUTION
        PANIC
    )

These codes should be pretty self-explaining.
The first adapter that returns a STOP_EXECUTION code breaks the event chain and aborts all further operations.

Message Prefixing

Usually chatbots require their own prefix to trigger executions.
Example:

    !echo hello world

The prefix parsing is done by equinox using a special "Prefix Handler".
If your bot only support a single static prefix, creating this adapter is rather easy:

    router.SetPrefixHandler(
        equinox.NewStaticPrefix("!"),
    )

If your bot needs to support dynamic prefixes (eg configurable on a per-server level) you can easily implement your own
prefix handler. Just create a function that resolves the prefix and returns a wormhole to it.

Example for discord:

    router.SetPrefixHandler(func(args ...*wormhole.Wormhole)(*wormhole.Wormhole){
        var prefix string

        // retrieve prefix from database //

        return wormhole.ToString(prefix)
    })

Note that the prefix handler is called for every incoming message.
You should implement some kind of prefix caching if you don't want degrading performance and/or overloaded database
servers.
 */
package equinox
