package caches

const (
    SESSION = "session"
    CHANNEL = "channel"
    SERVER  = "server"
    USER    = "user"
)

func Key(p string, s string) string {
    return p + "::" + s
}
