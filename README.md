> #### NOTE: This project is still HEAVILY work-in-progress.
> #### DO NOT USE IT IN YOUR PROJECTS UNTIL THIS WARNING IS REMOVED.

# Equinox

A little discord wrapper with some "oh that's neat" stuff

## Intro

Equinox is a framework that helps you creating expressive, yet easy to use and fun to code chatbots.
This is done by moving the hard stuff like proper caching and message parsing into the background.
Everything that's left is you and your modules.

Equinox is also heavily event-based and allows you to hook into almost any program stage.<br>
Whether you want high-level abstractions or just need a "drop-in router" to shorten your command-parsing code.<br>
The only limit is your imagination.

## The Router

The router is the heart of equinox.
It works like a normal event-dispatcher but doesn't require it's own coroutines or management functions.
It tries to follow KISS where possible and only runs code when explicitly told to.
