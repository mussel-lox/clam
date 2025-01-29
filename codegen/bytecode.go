package codegen

const (
	Constant OperationCode = iota
	Nil
	True
	False
	Fun
	Negate
	Not
	Add
	Subtract
	Multiply
	Divide
	Equal
	Greater
	Less
	GetGlobal
	SetGlobal
	GetLocal
	SetLocal
	Pop
	Closure
	Capture
	GetUpvalue
	SetUpvalue
	JumpIfFalse
	Jump
	Call
	Invoke
	Return
	Print
	Impossible
)

type OperationCode uint8

type ConstantIndex uint16
type GlobalIndex uint8
type LocalOffset uint8
type JumpOffset int16
type CallPosition uint16
