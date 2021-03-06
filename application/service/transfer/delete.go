package transfer

func (c *Transfer) Delete(identity string) (err error) {
	if c.Pipes.Empty(identity) {
		return
	}
	c.Pipes.Remove(identity)
	return c.Schema.Delete(identity)
}
