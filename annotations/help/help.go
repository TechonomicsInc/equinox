package help

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/equinox/annotations"
)

func Name(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Name = annotation.Value[0]
}

func Category(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Category = annotation.Value[0]
}

func Description(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Description = annotation.Value[0]
}

func Usage(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Usage = annotation.Value[0]
}

func Example(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Example = annotation.Value[0]
}

func Incubating(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Incubating = true
}

func Warning(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Warning = annotation.Value[0]
}

func Note(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Note = annotation.Value[0]
}

func Aliases(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Aliases = annotation.Value
}