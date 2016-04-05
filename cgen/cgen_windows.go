package cgen

const (
	ExitSuccess = 0
	ExitFailure = 1
)

func (c *CGen) emitExit(code int) {
	c.emit("push %d\n", code)     // status code
	c.emit("call _ExitProcess\n") // syscall
}
