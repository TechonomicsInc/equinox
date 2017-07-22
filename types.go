package equinox

import "reflect"

// POGOFunc is a void func without parameters
type POGOFunc func()

// POGOFuncW1 is a void func with one parameter
type POGOFuncW1 func(arg interface{})

// POGOFuncWV is a void func with n parameters
type POGOFuncWV func(args ...interface{})

// A NOOP that satisfies POGOFunc
func NOOP() {}

// A NOOP that satisfies POGOFuncW1
func NOOPW1(arg interface{}) {}

// A NOOP that satiesfies POGOFuncWV
func NOOPWV(args ...interface{}) {}

func TypeOf(v interface{}) string {
    t := reflect.TypeOf(v)

    if t.Kind() == reflect.Ptr {
        return "*" + t.Elem().Name()
    }

    return t.Name()
}