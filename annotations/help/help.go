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

func ShowAliases(annotation *annotations.Annotation, handler equinox.Handler, router *equinox.Router) {
    CreateMappingIfNeeded(handler)
    SanitizeHelpAnnotation(annotation)
    helpMapping[handler].ShowAliases = true
}
