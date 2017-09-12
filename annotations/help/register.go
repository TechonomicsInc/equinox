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
    "strings"
)

var annotationMap = map[string]equinox.AnnotationHandler{
    "Incubating":  Incubating,
    "Warning":     Warning,
    "Warn":        Warning,
    "Note":        Note,
    "Hint":        Note,
    "Name":        Name,
    "Category":    Category,
    "Description": Description,
    "Descr":       Description,
    "Summary":     Description,
    "Usage":       Usage,
    "Example":     Example,
    "Aliases":     ShowAliases,
    "ShowAliases": ShowAliases,
}

func GetListenersForHandler(handler equinox.Handler) []string {
    buf := []string{}

    for listener, h := range router.Routes {
        if h == handler {
            buf = append(buf, listener)
        }
    }

    return buf
}

func InjectInto(r *equinox.Router) {
    for k, v := range annotationMap {
        r.RegisterAnnotationHandler(k, v)
    }

    router = r
}

func GetForRoute(route string) *Help {
    for _, kind := range []string{"", "{p}", "{@}"} {
        if l, ok := router.Routes[kind+route]; ok {
            return GetForHandler(l)
        }
    }

    return nil
}

func GetForHandler(handler equinox.Handler) *Help {
    if help, ok := helpMapping[handler]; ok {
        return help
    }

    return nil
}

func GetForName(name string) *Help {
    name = strings.ToLower(name)

    for _, help := range helpMapping {
        if strings.ToLower(help.Name) == name {
            return help
        }
    }

    return nil
}
