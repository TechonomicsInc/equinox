// Code generated by "stringer -type AdapterEvent adapter_events.go"; DO NOT EDIT.

package equinox

import "fmt"

const _AdapterEvent_name = "NO_HANDLERS_REGISTEREDCONTINUE_EXECUTIONSTOP_EXECUTIONPANIC"

var _AdapterEvent_index = [...]uint8{0, 22, 40, 54, 59}

func (i AdapterEvent) String() string {
    if i < 0 || i >= AdapterEvent(len(_AdapterEvent_index)-1) {
        return fmt.Sprintf("AdapterEvent(%d)", i)
    }
    return _AdapterEvent_name[_AdapterEvent_index[i]:_AdapterEvent_index[i+1]]
}
