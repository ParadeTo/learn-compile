package playscript

var nullObject *NullObject

type NullObject struct {
	*ClassObject
}

func (*NullObject) ToString() string {
	return "Null"
}

func GetNullObject() *NullObject {
	if nullObject == nil {
		nullObject = &NullObject{}
	}
	return nullObject
}
