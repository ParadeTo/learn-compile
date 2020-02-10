package playscript

var breakObject *BreakObject

type BreakObject struct {
}

func (breakObject *BreakObject) ToString() string {
	return "BreakObject"
}

func GetBreakObject() *BreakObject {
	if breakObject == nil {
		breakObject = &BreakObject{}
	}
	return breakObject
}
