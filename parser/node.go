package parser

const NodeTypeIdent = "ident"
const NodeTypeNumber = "number"
const NodeTypeString = "string"
const NodeTypeUnary = "unary"
const NodeTypeBinary = "binary"
const NodeTypeTernary = "ternary"
const NodeTypeCall = "call"
const NodeTypeIndex = "index"
const NodeTypeSelector = "selector"
const NodeTypeParen = "paren"
const NodeTypeArray = "array"
const NodeTypePair = "pair"
const NodeTypeObject = "object"

func NewIdentNode(token *Token) *Node {
	return &Node{type_: NodeTypeIdent, token: token}
}

func NewNumberNode(token *Token) *Node {
	return &Node{type_: NodeTypeNumber, token: token}
}

func NewStringNode(token *Token) *Node {
	return &Node{type_: NodeTypeString, token: token}
}

func NewUnaryNode(op *Token, target *Node) *Node {
	return &Node{type_: NodeTypeUnary, op: op, x: target}
}

func NewBinaryNode(op *Token, lhs, rhs *Node) *Node {
	return &Node{type_: NodeTypeBinary, op: op, x: lhs, y: rhs}
}

func NewTernaryNode(condition, lhs, rhs *Node) *Node {
	return &Node{type_: NodeTypeTernary, x: condition, y: lhs, z: rhs}
}

func NewCallNode(callee *Node, args []*Node) *Node {
	return &Node{type_: NodeTypeCall, x: callee, s: args}
}

func NewIndexNode(target *Node, key *Node) *Node {
	return &Node{type_: NodeTypeIndex, x: target, y: key}
}

func NewSelectorNode(target *Node, key *Token) *Node {
	return &Node{type_: NodeTypeSelector, x: target, token: key}
}

func NewParenNode(n *Node) *Node {
	return &Node{type_: NodeTypeParen, x: n}
}

func NewArrayNode(items []*Node) *Node {
	return &Node{type_: NodeTypeArray, s: items}
}

func NewPairNode(k *Token, v *Node) *Node {
	return &Node{type_: NodeTypePair, token: k, x: v}
}

func NewObjectNode(items []*Node) *Node {
	return &Node{type_: NodeTypeObject, s: items}
}

type Node struct {
	type_ string
	token *Token // ident, number, string
	op    *Token // unary, binary
	x     *Node  // unary, binary lhs, call callee, ternary condition
	y     *Node  // binary rhs, ternary lhs
	z     *Node  // ternary rhs
	s     []*Node
}

func (n *Node) UnaryTarget() *Node {
	return n.x
}

func (n *Node) BinaryLhs() *Node {
	return n.x
}

func (n *Node) BinaryRhs() *Node {
	return n.y
}

func (n *Node) Callee() *Node {
	return n.x
}

func (n *Node) Args() []*Node {
	return n.s
}

func (n *Node) IndexTarget() *Node {
	return n.x
}

func (n *Node) IndexKey() *Node {
	return n.y
}

func (n *Node) SelectorTarget() *Node {
	return n.x
}

func (n *Node) SelectorKey() string {
	return n.token.Text
}

func (n *Node) ParenTarget() *Node {
	return n.x
}

func (n *Node) ArrayItems() []*Node {
	return n.s
}

func (n *Node) ObjectItems() []*Node {
	return n.s
}

func (n *Node) PairKeyValue() (*Token, *Node) {
	return n.token, n.x
}

func (n *Node) TernaryCondition() *Node {
	return n.x
}

func (n *Node) TernaryLhs() *Node {
	return n.y
}

func (n *Node) TernaryRhs() *Node {
	return n.z
}

func (n *Node) Type() string {
	return n.type_
}

func (n *Node) Ident() string {
	return n.token.Text
}

func (n *Node) Number() string {
	return n.token.Text
}

func (n *Node) String() string {
	if n.token == nil {
		return ""
	}
	return n.token.Text
}

func (n *Node) Op() string {
	return n.op.Text
}
