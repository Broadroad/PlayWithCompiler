package craft

/**
 * a simple Token。
 * only include type and test two types。
 */
type Token interface {
	GetType() TokenType
}
