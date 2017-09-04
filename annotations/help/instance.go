package help

import "code.lukas.moe/x/equinox"

type HelpMapping = map[equinox.Handler]*Help

type Help struct {
    Name        string
    Category    string
    Description string
    Usage       string
    Example     string
}

var (
    helpMapping = HelpMapping{}
    router      *equinox.Router
)

func GetOverview() map[string][]string {
    // Category -> []CommandName
    tmp := map[string][]string{}

    for route, handler := range router.Routes {
        mapping := helpMapping[handler]
        tmp[mapping.Category] = append(tmp[mapping.Category], route)
    }

    return tmp
}

func CreateMappingIfNeeded(handler equinox.Handler) {
    if _, ok := helpMapping[handler]; !ok {
        helpMapping[handler] = &Help{}
    }
}
