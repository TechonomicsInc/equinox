package equinox

import (
    "reflect"
    "github.com/bwmarrin/discordgo"
)

// POGOFunc is a void func without parameters
type POGOFunc func()

// POGOFuncW1 is a void func with one parameter
type POGOFuncW1 func(msg *discordgo.Message)

// POGOFuncWV is a void func with n parameters
type POGOFuncWV func(msg *discordgo.Message, args ...interface{})

// A NOOP that satisfies POGOFunc
func NOOP() {}

// A NOOP that satisfies POGOFuncW1
func NOOPW1(msg *discordgo.Message) {}

// A NOOP that satiesfies POGOFuncWV
func NOOPWV(msg *discordgo.Message, args ...interface{}) {}

func TypeOf(v interface{}) string {
    t := reflect.TypeOf(v)

    if t.Kind() == reflect.Ptr {
        return "*" + t.Elem().Name()
    }

    return t.Name()
}