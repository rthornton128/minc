package cgen

const (
	ExitSuccess = 0
	ExitFailure = 1
)

func (c *CGen) emitExit(code int) {
	c.emit("mov rax, 0x3c\n")     // exit code
	c.emit("mov rdi, %d\n", code) // status code
	c.emit("syscall\n")           // syscall
}
