package craft

type TokenReader interface {
	Read() Token
	Peek() Token
	UnRead()
	GetPosition() int
	SetPosition()
}
