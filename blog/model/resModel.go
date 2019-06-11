package model

import (
	"encoding/json"
	"fmt"
)

type BaseModel interface {
	String() string
}

type myModel struct {
	BaseModel
	msg   string
	obj   interface{}
	errNo int
}

func (m *myModel) String() string {
	if m.obj == nil {
		return fmt.Sprintf(`{"msg":"%s", "errno":"%d"}`, m.msg, m.errNo)
	}
	o, _ := json.Marshal(m.obj)
	return fmt.Sprintf(`{"errno":"%d", %s}`, m.errNo, o)
}

func SuccessModel(a ...interface{}) BaseModel {
	if len(a) != 1 {
		panic("invalid argument")
		return nil
	}

	m, ok := a[0].(string)
	if ok {
		return &myModel{
			BaseModel: nil,
			msg:       m,
			obj:       nil,
			errNo:     0,
		}
	} else {
		return &myModel{
			BaseModel: nil,
			msg:       "",
			obj:       a[0],
			errNo:     0,
		}
	}

}

func ErrorModel(a ...interface{}) BaseModel {
	if len(a) == 1 {
		m, ok := a[0].(string)
		if ok {
			return &myModel{
				BaseModel: nil,
				msg:       m,
				obj:       nil,
				errNo:     -1,
			}
		}
	} else if len(a) == 2 {
		o := a[0]
		m, ok := a[1].(string)
		if ok {
			return &myModel{
				BaseModel: nil,
				msg:       m,
				obj:       o,
				errNo:     -1,
			}
		}
	}

	panic("invalid argument")
	return nil
}
