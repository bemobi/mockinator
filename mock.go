package mockinator

import (
	"reflect"
	"runtime"
	"sync/atomic"
)

type Mockinator struct {
	Count      map[string]*int64
	CustomFunc map[string]func()
	Error      map[string]error
	Return     map[string]interface{}
}

func (m *Mockinator) MustInit() {
	m.Error = make(map[string]error)
	m.Count = make(map[string]*int64)
	m.CustomFunc = make(map[string]func())
	m.Return = make(map[string]interface{})
}

func (m *Mockinator) Execute(funcInterface interface{}) (interface{}, error) {
	if m.Count[m.getNameByInterface(funcInterface)] == nil {
		m.Count[m.getNameByInterface(funcInterface)] = new(int64)
	}
	atomic.AddInt64(m.Count[m.getNameByInterface(funcInterface)], 1) // increment

	if m.CustomFunc[m.getNameByInterface(funcInterface)] != nil {
		m.CustomFunc[m.getNameByInterface(funcInterface)]() // Call customFunc if exist
	}

	return m.Return[m.getNameByInterface(funcInterface)], m.Error[m.getNameByInterface(funcInterface)] // return
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

func (m *Mockinator) getNameByInterface(funcInterface interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(funcInterface).Pointer()).Name()
}
