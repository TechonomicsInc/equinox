// Code generated by "stringer -type=Context"; DO NOT EDIT.

package annotations

import "fmt"

const _Context_name = "CONTEXT_ROOTCONTEXT_ANNOTATIONCONTEXT_PARAMSCONTEXT_STRINGCONTEXT_COMMENT_BLOCKCONTEXT_COMMENT_LINE"

var _Context_index = [...]uint8{0, 12, 30, 44, 58, 79, 99}

func (i Context) String() string {
	if i < 0 || i >= Context(len(_Context_index)-1) {
		return fmt.Sprintf("Context(%d)", i)
	}
	return _Context_name[_Context_index[i]:_Context_index[i+1]]
}
