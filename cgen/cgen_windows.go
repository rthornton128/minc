package cgen

func (c *CGen) emitExit(code int) {
	c.emit("push %d\n", code)     // status code
	c.emit("call _ExitProcess\n") // syscall
}
