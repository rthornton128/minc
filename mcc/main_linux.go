package main

import (
	"fmt"
	"os"
)

func execLinker(lnk, flags, name string) {
	if flags == "" {
		flags := fmt.Sprintf("-o %s -e main %s.o", name, name)
	}
	execProg(lnk, flags)
	os.Remove(name)
}
