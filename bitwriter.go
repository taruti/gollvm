package llvm

/*
#define __STDC_LIMIT_MACROS
#define __STDC_CONSTANT_MACROS
#include <llvm-c/BitWriter.h>
#include <stdlib.h>
*/
import "C"
import "errors"
import "unsafe"

var writeBitcodeToFileErr = errors.New("Failed to write bitcode to file")

func WriteBitcodeToFile(m Module, filename string) error {
	cfilename := C.CString(filename)
	result := C.LLVMWriteBitcodeToFile(m.C, cfilename)
	C.free(unsafe.Pointer(cfilename))
	if result == 0 {
		return nil
	}
	return writeBitcodeToFileErr
}

// TODO(nsf): Figure out way how to make it work with io.Writer
