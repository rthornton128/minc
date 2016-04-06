package main

import (
	"fmt"
	"path/filepath"
)

func execLinker(lnk, flags, name string) {
	sysPath := string(filepath.Separator) + filepath.Join("Windows", "System32")
	if flags == "" {
		flags = fmt.Sprintf("-o %[1]s.exe -e main -L %[2]s %[1]s.o -lkernel32",
			name, sysPath)
	}
	execProg(lnk, flags, name+".o")
}
