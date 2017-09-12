/*
 * Copyright (C) 2017 Subliminal Apps
 *
 * Licensed under the EUPL, Version 1.1 only (the "Licence");
 *
 * You may not use this work except in compliance with the Licence.
 * You may obtain a copy of the Licence at:
 * <https://joinup.ec.europa.eu/software/page/eupl>
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Licence is distributed on an "AS IS" basis,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the Licence for the specific language governing permissions and limitations
 * under the Licence.
 */

package help

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/equinox/annotations"
    "regexp"
    "strings"
)

type HelpMapping = map[equinox.Handler]*Help

type Help struct {
    HRef equinox.Handler

    Name string

    Category    string
    Description string
    Usage       string
    Example     string

    ShowAliases bool
    Incubating  bool
    Warning     string
    Note        string
}

var (
    router       *equinox.Router
    helpMapping  = HelpMapping{}
    spaceTrimmer = regexp.MustCompile(`\n(\ +)`)

    abbreviations = map[string]string{
        "[^]":     "`",
        "[code]":  "`",
        "[/code]": "`",
    }
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
        helpMapping[handler] = &Help{HRef: handler}
    }
}

func SanitizeHelpAnnotation(annotation *annotations.Annotation) {
    annotation.Key = spaceTrimmer.ReplaceAllString(annotation.Key, "\n")

    for i := range annotation.Value {
        annotation.Value[i] = spaceTrimmer.ReplaceAllString(annotation.Value[i], "\n")

        for b, a := range abbreviations {
            annotation.Value[i] = strings.Replace(annotation.Value[i], b, a, -1)
        }
    }
}
