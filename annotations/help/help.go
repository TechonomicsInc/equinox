package help

import (
    "code.lukas.moe/x/equinox"
    "code.lukas.moe/x/equinox/annotations"
)

// @Name("CoolPlugin")
func Name(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Name = annotation.Value[0]
}

// @Category("Funny Stuff")
func Category(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Category = annotation.Value[0]
}

// @Description("Cool Description Here")
func Description(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Description = annotation.Value[0]
}

// @Descr("Cool Description Here")
// Alias for @Description
func Descr(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    Description(annotation, handler, router)
}

// @Summary("Cool Description Here")
// Alias for @Description
func Summary(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    Description(annotation, handler, router)
}

// @Usage("How to do stuff")
func Usage(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Usage = annotation.Value[0]
}

// @Example("This is how to do stuff")
func Example(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].Example = annotation.Value[0]
}
