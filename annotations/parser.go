package annotations

import (
    "unicode"
)

//go:generate stringer -type=Context
type Context int

const (
    ANNOTATION_CREATE = '@'
    PARAMS_OPEN       = '('
    PARAMS_CLOSE      = ')'
    NEW_ARG           = ','
    STRING_BOUND      = '"'

    COMMENT_BOUND = '#'
    COMMENT_END   = '\n'

    CONTEXT_ROOT       Context = iota
    CONTEXT_ANNOTATION
    CONTEXT_PARAMS
    CONTEXT_STRING
    CONTEXT_COMMENT
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

    // Loop through the input creating annotations as we move
    for _, char := range []rune(input) {
        switch context {
        case CONTEXT_ROOT:
            // In root context search for annotations to create
            switch char {
            case COMMENT_BOUND:
                commentedContext = context
                context = CONTEXT_COMMENT

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
                    panic("Unexpected character '" + string(char) + "' at root level!")
                }
            }

        case CONTEXT_ANNOTATION:
            // In annotation context search for params.
            // Not having params is ok too and results in an empty param array.
            // You can not leave out the () since ) resets the context.
            switch char {
            case COMMENT_BOUND:
                commentedContext = context
                context = CONTEXT_COMMENT

            case PARAMS_OPEN:
                // A ( appeared thus signaling that params follow.
                context = CONTEXT_PARAMS

            default:
                // We're in Annotation context but no special chars appeared.
                // We can assume that char is part of the annotation's name.
                if unicode.IsLetter(char) || unicode.IsNumber(char) {
                    buf.Key += string(char)
                } else {
                    panic("Unexpected character '" + string(char) + "' in annotation name!")
                }
            }

        case CONTEXT_PARAMS:
            switch char {
            case COMMENT_BOUND:
                commentedContext = context
                context = CONTEXT_COMMENT

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
                    panic("Unexpected character '" + string(char) + "' parameter literal!")
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

        case CONTEXT_COMMENT:
            switch char {
            case COMMENT_BOUND, COMMENT_END:
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
