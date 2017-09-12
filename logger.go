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
    "fmt"
    "time"
)

func logf(format string, a ...interface{}) {
    if debugMode {
        log(fmt.Sprintf(format, a...))
    }
}

func log(msg string) {
    if debugMode {
        fmt.Printf(
            "[%s] (%s) %s\n",
            time.Now().Format("15:04:05 02-01-2006"),
            "EQUINOX",
            msg,
        )
    }
}
