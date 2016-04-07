package cgen

func (c *CGen) emitExit(code int) {
	c.emit("mov rax, 0x3c\n")     // exit code
	c.emit("mov rdi, %d\n", code) // status code
	c.emit("syscall\n")           // syscall
}
