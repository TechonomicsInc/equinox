package help

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/equinox/annotations"
    "regexp"
)

type HelpMapping = map[equinox.Handler]*Help

type Help struct {
    Name        string
    Category    string
    Description string
    Usage       string
    Example     string
}

var (
    router       *equinox.Router
    helpMapping  = HelpMapping{}
    spaceTrimmer = regexp.MustCompile(`\n(\ +)`)
)

func GetOverview() map[string][]string {
    // Category -> []PluginName
    tmp := map[string][]string{}

    for _, mapping := range helpMapping {
        tmp[mapping.Category] = append(tmp[mapping.Category], mapping.Name)
    }

    return tmp
}

func CreateMappingIfNeeded(handler equinox.Handler) {
    if _, ok := helpMapping[handler]; !ok {
        helpMapping[handler] = &Help{}
    }
}

func SanitizeHelpAnnotation(annotation *annotations.Annotation) {
    annotation.Key = spaceTrimmer.ReplaceAllString(annotation.Key, "\n")

    for i := range annotation.Value {
        annotation.Value[i] = spaceTrimmer.ReplaceAllString(annotation.Value[i], "\n")
    }
}
