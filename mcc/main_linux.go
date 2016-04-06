package main

import "fmt"

func execLinker(lnk, flags, name string) {
	if flags == "" {
		flags = fmt.Sprintf("-o %[1]s -e main %[1]s.o", name)
	}
	execProg(lnk, flags, name+".o")
}
