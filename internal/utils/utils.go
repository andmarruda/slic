package utils

import (
	"strings"
)

func changeExtension(filename, newExt string) string {
    if dot := strings.LastIndex(filename, "."); dot != -1 {
        return filename[:dot] + "." + newExt
    }
    return filename + "." + newExt
}
