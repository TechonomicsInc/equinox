package equinox

//go:generate stringer -type Event

// Event is a simple enum that is used to register and dispatch events
type Event int

const (
    // Called before any message parsing is done
    MESSAGE_PRE_ANALYZE  Event = iota

    // Called after the message has been identified as "for us" but before any advanced parsing
    MESSAGE_ANALYZE

    // Called after the parsing is finished
    MESSAGE_POST_ANALYZE

    // Called to wake adapters that check if any @mentions are present.
    MESSAGE_CHECK_MENTIONS

    // Called to wake adapters adapters that check if any @mentions for us are present
    MESSAGE_CHECK_OUR_MENTIONS

    // Called before the matched handler is executed
    HANDLER_PRE_EXECUTE

    // Called after the matched handler was executed
    HANDLER_POST_EXECUTE

    // Called before executing the last resort
    LAST_RESORT_PRE_EXECUTE

    // Called after executing the last resort
    LAST_RESORT_POST_EXECUTE
)
