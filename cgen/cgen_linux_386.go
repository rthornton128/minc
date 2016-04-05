package cgen

const (
	ExitSuccess = 0
	ExitFailure = 1
)

func (c *CGen) emitExit(code int) {
	c.emit("mov eax, 1\n")        // exit code
	c.emit("mov ebx, %d\n", code) // status code
	c.emit("int 0x80\n")          // syscall
}
