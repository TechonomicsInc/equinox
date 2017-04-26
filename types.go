package equinox

import "code.lukas.moe/x/wormhole"

// POGOFunc is a void func without parameters
type POGOFunc func()

// POGOFuncW1 is a void func with one parameter
type POGOFuncW1 func(arg *wormhole.Wormhole)

// POGOFuncWV is a void func with n parameters
type POGOFuncWV func(args ...*wormhole.Wormhole)

// A NOOP that satisfies POGOFunc
func NOOP() {}

// A NOOP that satisfies POGOFuncW1
func NOOPW1(arg *wormhole.Wormhole) {}

// A NOOP that satiesfies POGOFuncWV
func NOOPWV(args ...*wormhole.Wormhole) {}
