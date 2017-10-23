package lua

func OpenBit(L *LState) int {
	mod := L.RegisterModule(BitLibName, bitFuncs).(*LTable)
	L.Push(mod)
	return 1
}

var bitFuncs = map[string]LGFunction{
	"band":   bitAnd,
	"bor":    bitBor,
	"lshift": bitLShift,
	"rshift": bitRShift,
}

func bitBor(L *LState) int {
	L.Push(LNumber(int32(L.CheckNumber(1)) | int32(L.CheckNumber(2))))
	return 1
}
func bitAnd(L *LState) int {
	L.Push(LNumber(int32(L.CheckNumber(1)) & int32(L.CheckNumber(2))))
	return 1
}

func bitLShift(L *LState) int {
	L.Push(LNumber(int32(L.CheckNumber(1)) << uint(L.CheckNumber(2))))
	return 1
}

func bitRShift(L *LState) int {
	L.Push(LNumber(int32(L.CheckNumber(1)) >> uint(L.CheckNumber(2))))
	return 1
}
