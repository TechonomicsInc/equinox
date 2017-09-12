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

package equinox

import (
    "reflect"
    "github.com/bwmarrin/discordgo"
    "code.lukas.moe/x/equinox/annotations"
)

type AnnotationHandler func(annotation *annotations.Annotation, handler Handler, router *Router)

type RuntimeAdapter func(handler Handler, msg *discordgo.Message, router *Router) AdapterEvent

// POGOFunc is a void func without parameters
type POGOFunc func()

// POGOFuncW1 is a void func with one parameter
type POGOFuncW1 func(msg *discordgo.Message)

// POGOFuncWV is a void func with n parameters
type POGOFuncWV func(msg *discordgo.Message, args ...interface{})

// A NOOP that satisfies POGOFunc
func NOOP() {}

// A NOOP that satisfies POGOFuncW1
func NOOPW1(msg *discordgo.Message) {}

// A NOOP that satiesfies POGOFuncWV
func NOOPWV(msg *discordgo.Message, args ...interface{}) {}

func TypeOf(v interface{}) string {
    t := reflect.TypeOf(v)

    if t.Kind() == reflect.Ptr {
        return "*" + t.Elem().Name()
    }

    return t.Name()
}
