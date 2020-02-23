package playscript

type PlayObject interface {
	ToString() string
	GetFields() map[*Variable]interface{}
	GetValue(variable *Variable) interface{}
	SetValue(variable *Variable, value interface{})
}

type BasePlayObject struct {
	fields map[*Variable]interface{}
}

func (this *BasePlayObject) PutAllFields(fields map[*Variable]interface{}) {
	for v, fieldValue := range fields {
		this.SetValue(v, fieldValue)
	}
}

func (this *BasePlayObject) SetValue(variable *Variable, value interface{}) {
	this.fields[variable] = value
}

func (this *BasePlayObject) GetFields() map[*Variable]interface{} {
	return this.fields
}

func (this *BasePlayObject) ToString() string {
	return ""
}

func (this *BasePlayObject) GetValue(variable *Variable) interface{} {
	object := this.fields[variable]
	if object == nil {
		object = GetNullObject()
	}
	return object
}

func NewBasePlayObject() *BasePlayObject {
	return &BasePlayObject{fields: make(map[*Variable]interface{})}
}
