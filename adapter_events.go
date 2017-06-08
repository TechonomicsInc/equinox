package equinox

//go:generate stringer -type AdapterEvent

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
