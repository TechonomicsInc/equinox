package help

import (
    "code.lukas.moe/x/equinox"
    "strings"
)

func InjectInto(r *equinox.Router) {
    r.RegisterAnnotationHandler("Name", Name)
    r.RegisterAnnotationHandler("Category", Category)

    r.RegisterAnnotationHandler("Description", Description)
    r.RegisterAnnotationHandler("Descr", Descr)
    r.RegisterAnnotationHandler("Summary", Summary)

    r.RegisterAnnotationHandler("Usage", Usage)
    r.RegisterAnnotationHandler("Example", Example)

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
