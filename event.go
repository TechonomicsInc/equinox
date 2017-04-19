package equinox

type Event int

const (
    MESSAGE_PRE_ANALYZE  Event = iota
    MESSAGE_ANALYZE
    MESSAGE_POST_ANALYZE
)
