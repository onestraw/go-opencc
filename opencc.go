package opencc

// #cgo LDFLAGS: -lopencc
// #include "opencc/opencc.h"
import "C"
import "errors"

type OpenccHandle struct {
	handle C.opencc_t
}

func Error() error {
	err := C.opencc_error()
	if C.GoString(err) != "" {
		return errors.New(C.GoString(err))
	}
	return nil
}

func New(config string) (*OpenccHandle, error) {
	h := C.opencc_open(C.CString(config))
	if h == nil {
		return nil, errors.New("opencc_open fail")
	}
	return &OpenccHandle{handle: h}, Error()
}

func (h *OpenccHandle) Convert(input string) (string, error) {
	output := C.opencc_convert_utf8(h.handle, C.CString(input), C.ulong(len(input)))
	defer C.opencc_convert_utf8_free(output)
	return C.GoString(output), Error()
}

func (h *OpenccHandle) Free() error {
	C.opencc_close(h.handle)
	return Error()
}
