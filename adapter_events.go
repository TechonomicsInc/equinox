//go:generate stringer -type AdapterEvent
package equinox

// AdapterEvent is an enum that tells the router to continue/stop execution
type AdapterEvent int

const (
    CONTINUE_EXECUTION     AdapterEvent = iota
    STOP_EXECUTION
    PANIC
    NO_HANDLERS_REGISTERED
)

func (a AdapterEvent) ShouldAbort() bool {
    return a == STOP_EXECUTION || a == NO_HANDLERS_REGISTERED
}

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
