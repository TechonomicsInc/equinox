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

package annotations

import (
    "unicode"
)

//go:generate stringer -type=Context
type Context int

type Compound = [2]rune

const (
    CONTEXT_ROOT       Context = iota
    CONTEXT_ANNOTATION
    CONTEXT_PARAMS
    CONTEXT_STRING

    CONTEXT_COMMENT_BLOCK
    CONTEXT_COMMENT_LINE
)

var (
    ANNOTATION_CREATE = '@'
    PARAMS_OPEN       = '('
    PARAMS_CLOSE      = ')'
    NEW_ARG           = ','
    STRING_BOUND      = '"'

    COMMENT_LINE        = Compound{'/', '/'}
    COMMENT_BLOCK_START = Compound{'/', '*'}
    COMMENT_BLOCK_END   = Compound{'*', '/'}
)

type Annotation struct {
    Key   string
    Value []string
}

func Parse(input string) []*Annotation {
    // Container for results
    parsed := []*Annotation{}

    // Holds the current context
    context := CONTEXT_ROOT

    // Holds the previous context when entering comments
    commentedContext := CONTEXT_ROOT

    // Buffers
    paramBuf := ""
    var buf *Annotation

    // Skipper
    // Used to skip n chars from within the loop
    skipper := 0

    // Loop through the input creating annotations as we move
    cinput := []rune(input)
    for pos, char := range cinput {
        if skipper > 0 {
            skipper--
            continue
        }

        switch context {
        case CONTEXT_ROOT:
            // In root context search for annotations to create
            switch char {
            case COMMENT_LINE[0], COMMENT_BLOCK_START[0]:
                if cinput[pos+1] == COMMENT_LINE[1] {
                    commentedContext = context
                    context = CONTEXT_COMMENT_LINE
                    skipper += 1
                }

                if cinput[pos+1] == COMMENT_BLOCK_START[1] {
                    commentedContext = context
                    context = CONTEXT_COMMENT_BLOCK
                    skipper += 1
                }

            case ANNOTATION_CREATE:
                // Creating annotations involves updating buf and setting a new ctx
                if buf != nil {
                    parsed = append(parsed, buf)
                }

                buf = &Annotation{}
                context = CONTEXT_ANNOTATION

            default:
                // Check if the char is spacing.
                // These are allowed endlessly to reduce the work of
                // trimming the string.
                if !unicode.IsSpace(char) {
                    // Other things are not allowed
                    printParseError(input, pos, context, "Unexpected character at root level!")
                }
            }

        case CONTEXT_ANNOTATION:
            // In annotation context search for params.
            // Not having params is ok too and results in an empty param array.
            // You can not leave out the () since ) resets the context.
            switch char {
            case COMMENT_LINE[0], COMMENT_BLOCK_START[0]:
                if cinput[pos+1] == COMMENT_LINE[1] {
                    commentedContext = context
                    context = CONTEXT_COMMENT_LINE
                    skipper += 1
                }

                if cinput[pos+1] == COMMENT_BLOCK_START[1] {
                    commentedContext = context
                    context = CONTEXT_COMMENT_BLOCK
                    skipper += 1
                }

            case PARAMS_OPEN:
                // A ( appeared thus signaling that params follow.
                context = CONTEXT_PARAMS

            default:
                // We're in Annotation context but no special chars appeared.
                // We can assume that char is part of the annotation's name.
                if unicode.IsLetter(char) || unicode.IsNumber(char) {
                    buf.Key += string(char)
                } else {
                    printParseError(input, pos, context, "Unexpected character in annotation name!")
                }
            }

        case CONTEXT_PARAMS:
            switch char {
            case COMMENT_LINE[0], COMMENT_BLOCK_START[0]:
                if cinput[pos+1] == COMMENT_LINE[1] {
                    commentedContext = context
                    context = CONTEXT_COMMENT_LINE
                    skipper += 1
                }

                if cinput[pos+1] == COMMENT_BLOCK_START[1] {
                    commentedContext = context
                    context = CONTEXT_COMMENT_BLOCK
                    skipper += 1
                }

            case PARAMS_CLOSE:
                // A ) appeared thus signaling that the params are over.
                // Since this is the last thing an annotation contains we can reset the context
                // to root and start over with the next annotation.
                context = CONTEXT_ROOT

            case STRING_BOUND:
                // A string just started.
                // Switch context
                context = CONTEXT_STRING

            case NEW_ARG:
                // A new argument begins.
                // This token is mostly unused until other parameter types are allowed.

            default:
                // Allow spacing between args but no other chars
                if !unicode.IsSpace(char) {
                    printParseError(input, pos, context, "Unexpected character in parameter literal!")
                }
            }

        case CONTEXT_STRING:
            switch char {
            case STRING_BOUND:
                // End of string is reached.
                // Reset context.
                context = CONTEXT_PARAMS

                // Also write the string into the annotation-struct
                buf.Value = append(buf.Value, paramBuf)
                paramBuf = ""

            default:
                // Char belongs to string.
                // Add to buffer
                paramBuf += string(char)
            }

        case CONTEXT_COMMENT_BLOCK:
            if char == COMMENT_BLOCK_END[0] && cinput[pos+1] == COMMENT_BLOCK_END[1] {
                context = commentedContext
                skipper += 1
            }

        case CONTEXT_COMMENT_LINE:
            if char == '\n' {
                context = commentedContext
            }
        }
    }

    // Add the remaining last annotation (if any)
    if buf != nil {
        parsed = append(parsed, buf)
    }

    // Profit!
    return parsed
}

func printParseError(input string, pos int, ctx Context, emsg string) {
    cinput := []rune(input)

    pre := cinput[:pos]
    post := cinput[pos+1:]

    err := "\n [=================================================]\n"

    err += string(pre)
    err += " ~~> " + string(cinput[pos]) + " <~~ "
    err += string(post)

    err += "\n [=================================================]\n"

    err += " Equinox encountered an annotation that it could not understand.\n"
    err += " The error is marked as ~~> E <~~ in the code above.\n"
    err += "\n"

    err += " Context: " + ctx.String() + "\n"
    err += " Error: " + emsg + "\n"
    err += "\n"
    err += " Try to check if the context makes sense in the code above.\n"
    err += " For example: Seeing a " + CONTEXT_PARAMS.String() + " in a string indicates unclosed string literals.\n"
    err += " Happy hacking!"

    err += "\n [=================================================]\n"

    panic(err)
}
