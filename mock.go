package mockinator

import (
	"reflect"
	"sync/atomic"
)

type Mockinator struct {
	Count      map[interface{}]*int64
	CustomFunc map[interface{}]func()
	Error      map[interface{}]error
	Return     map[interface{}]interface{}
}

func (m *Mockinator) MustInit() {
	m.Error = make(map[interface{}]error)
	m.Count = make(map[interface{}]*int64)
	m.CustomFunc = make(map[interface{}]func())
	m.Return = make(map[interface{}]interface{})
}

func (m *Mockinator) Execute(funcInterface interface{}) (interface{}, error) {
	atomic.AddInt64(m.Count[m.getNameByInterface(funcInterface)], 1) // increment

	if m.CustomFunc[m.getNameByInterface(funcInterface)] != nil {
		m.CustomFunc[m.getNameByInterface(funcInterface)]() // Call customFunc if exist
	}

	return m.Return[m.getNameByInterface(funcInterface)], m.Error[m.getNameByInterface(funcInterface)] // return
}

func (m *Mockinator) getNameByInterface(funcInterface interface{}) string {
	return reflect.TypeOf(funcInterface).Name()
}

func (m *Mockinator) ExecuteAndReturnErrorOnly(funcInterface interface{}) error {
	_, err := m.Execute(m.getNameByInterface(funcInterface))
	return err
}

func (m *Mockinator) GetCount(funcInterface interface{}) int64 {
	return *m.Count[m.getNameByInterface(funcInterface)]
}

func (m *Mockinator) SetError(funcInterface interface{}, errorData error) {
	m.Error[m.getNameByInterface(funcInterface)] = errorData
}

func (m *Mockinator) SetReturn(funcInterface interface{}, returnData interface{}) {
	m.Return[m.getNameByInterface(funcInterface)] = returnData
}

func (m *Mockinator) SetCustomFunc(funcInterface interface{}, customFunc func()) {
	m.CustomFunc[m.getNameByInterface(funcInterface)] = customFunc
}
