package craft

package craft
/**
 * Token的类型
 */
type TokenType int

const (
	Plus TokenType = iota, // +
		Minus, // -
		Star, // *
		Slash, // /

		GE, // >=
		GT, // >
		EQ, // ==
		LE, // <=
		LT, // <

		SemiColon, // ;
		LeftParen, // (
		RightParen, // )

		Assignment, // =

		If,
		Else,

		Int,

		Identifier, //标识符

		IntLiteral, //整型字面量
		StringLiteral //字符串字面量
)

func (d TokenType) String() string {
	return [...]string{"Minus", "Star", "Slash", "GE", "GT", "EQ", "LE", "LT", "SemiColon", "LeftParen", "RightParen", "Assignment", "If", "Else", "Int", "Identifier", "IntLiteral", "StringLiteral"}[d]
}
