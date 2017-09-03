package annotations

import (
    "unicode"
)

const (
    ANNOTATION_CREATE = '@'
    PARAMS_OPEN       = '('
    PARAMS_CLOSE      = ')'
    NEW_ARG           = ','
    STRING_BOUND      = '"'

    CONTEXT_ROOT       int = iota
    CONTEXT_ANNOTATION
    CONTEXT_PARAMS
    CONTEXT_STRING
)

type Annotation struct {
    Key   string
    Value []string
}

func Parse(input string) []*Annotation {
    // Container for results
    parsed := []*Annotation{}

    // Loop through the input creating annotations as we move
    context := CONTEXT_ROOT
    var buf *Annotation

    for _, char := range []rune(input) {
        switch context {
        case CONTEXT_ROOT:
            // In root context search for annotations to create
            switch char {
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
                    panic("Parse error at " + string(char))
                }
            }

        case CONTEXT_ANNOTATION:
            // In annotation context search for params.
            // Not having params is ok too and results in an empty param array.
            // You can not leave out the () since ) resets the context.
            switch char {
            case PARAMS_OPEN:
                // A ( appeared thus signaling that params follow.
                context = CONTEXT_PARAMS
            default:
                // We're in Annotation context but no special chars appeared.
                // We can assume that char is part of the annotation's name.
                buf.Key += string(char)
            }

        case CONTEXT_PARAMS:
            switch char {
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