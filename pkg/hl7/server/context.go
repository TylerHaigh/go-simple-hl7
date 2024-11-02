package server

type Ctx interface {
	Next() error

	Request() *Req
	Response() *Res
}

type DefaultCtx struct {
	Req   *Req
	Res   *Res
	Error error

	handlerIndex int
	handlers     []HandlerFunction

	errorHandlerIndex int
	errorHandlers     []ErrorFunction
}

func (c *DefaultCtx) Request() *Req {
	return c.Req
}

func (c *DefaultCtx) Response() *Res {
	return c.Res
}

func (c *DefaultCtx) Next() error {

	// Increment Handler Index
	c.handlerIndex++

	// Check if next handler exists
	stackSize := len(c.handlers)
	if c.handlerIndex >= stackSize {
		return nil
	}

	// Get the Handler
	h := c.handlers[c.handlerIndex]

	// Execute the Handler
	err := h(c)

	if err != nil {
		c.Error = err
		return c.nextError()
	}

	return nil
}

func (c *DefaultCtx) nextError() error {

	// Check if next handler exists
	stackSize := len(c.errorHandlers)
	if c.errorHandlerIndex >= stackSize {
		return nil
	}

	// Get the Handler
	h := c.errorHandlers[c.errorHandlerIndex]

	// Increment Handler Index
	c.errorHandlerIndex++

	h(c, c.Error)

	c.nextError()
	return c.Error
}

type HandlerFunction func(Ctx) error
type ErrorFunction func(Ctx, error) error
