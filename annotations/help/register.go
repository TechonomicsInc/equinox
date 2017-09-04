package help

import (
    "code.lukas.moe/x/equinox"
)

func InjectInto(r *equinox.Router) {
    r.RegisterAnnotationHandler("Name", Name)

    r.RegisterAnnotationHandler("Description", Description)
    r.RegisterAnnotationHandler("Descr", Descr)
    r.RegisterAnnotationHandler("Summary", Summary)

    r.RegisterAnnotationHandler("Usage", Usage)
    r.RegisterAnnotationHandler("Example", Example)

    router = r
}

func GetForRoute(route string) *Help {
    if l, ok := router.Routes[route]; ok {
        return GetForHandler(l)
    }

    return nil
}

func GetForHandler(handler equinox.Handler) *Help {
    if help, ok := helpMapping[handler]; ok {
        return help
    }

    return nil
}
