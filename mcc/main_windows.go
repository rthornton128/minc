package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func execLinker(lnk, flags, name string) {
	sysPath := string(filepath.Separator) + filepath.Join("Windows", "System32")
	if flags == "" {
		flags = fmt.Sprintf("-o %s.exe -e main -L %s %s.o -lkernel32", name,
			sysPath, name)
	}
	execProg(lnk, flags)
	os.Remove(name + ".o")
}
