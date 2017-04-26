package equinox

//go:generate stringer -type AdapterEvent

// AdapterEvent is an enum that tells the router to continue/stop execution
// You should always try to ignore "direct" AdapterEvent values and use the bound functions instead.
// This way the error-handling and Event-Parsing stays in a central place.
//
// The easiest usage form is to defer the corresponding panic-handler and call AdapterEvent#Act():
//
//     defer AdapterPanicHandler()
//     r.Dispatch(MY_COOL_EVENT, arg1, arg2, arg...).Act()
//
// If you want to have some more control about the panics/errors there is AdapterEvent#ShouldAbort().
// Just wrap the Dispatch call into an if-clause and return as needed.
//
//     if r.Dispatch(MY_COOL_EVENT, arg1, arg2, arg...).ShouldAbort() {
//         return
//     }
//
// And last but not least you can always throw the Event at a switch clause.
//
//     switch r.Dispatch(MY_COOL_EVENT, arg1, arg2, arg...) {
//     case CONTINUE_EXECUTION, NO_HANDLERS_REGISTERED:
//         /* Continue execution */
//         break
//
//     case STOP_EXECUTION:
//         /* Kill execution */
//         return
//
//     case PANIC:
//         panic("wew i did not see that coming")
//
//     }
//
// AdapterEvents might be added/removed/changed until equinox reaches a stable release so make sure to add a "default"
// case that handles unknown Events.
//
type AdapterEvent int

const (
    // Indicates that no handlers were registered.
    // It's recommended to act like the return value was CONTINUE_EXECUTION if this happens.
    NO_HANDLERS_REGISTERED AdapterEvent = iota

    // Indicates that the outer routine should continue to work normally
    CONTINUE_EXECUTION

    // Indicates that the outer routine should return as soon as possible
    STOP_EXECUTION

    // Indicates that the outer routine should throw a panic
    PANIC
)

// ShouldAbort is a convenience function that checks if a state was encountered that
// requires the outer routine to exit.
func (a AdapterEvent) ShouldAbort() bool {
    return a == STOP_EXECUTION || a == PANIC
}

// Act is a convenience function checks all possible event values and acts accordingly for you.
// That means:
// - CONTINUE_EXECUTION and NO_HANDLERS_REGISTERED is a NOOP.
// - Everything else panics with a proper (and parsable) error message.
func (a AdapterEvent) Act() {
    switch a {
    case
        CONTINUE_EXECUTION,
        NO_HANDLERS_REGISTERED:
        return

    case
        STOP_EXECUTION:
        panic(AdapterPanic{a, "NORMAL_EVENT_TERMINATION"})

    case PANIC:
        panic(AdapterPanic{a, "EXPECTED_ABNORMAL_EVENT_TERMINATION"})

    default:
        panic(AdapterPanic{a, "UNEXPECTED_ABNORMAL_EVENT_TERMINATION"})
    }
}
