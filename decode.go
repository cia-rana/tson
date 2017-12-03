package tson

import (
	"encoding/json"
	"reflect"
)

func Unmarshal(jsonBytes []byte, v interface{}) error {
	rt, err := NewStruct(v)
	if err != nil {
		return err
	}

	vi := reflect.New(rt).Interface()

	if err := json.Unmarshal(jsonBytes, &vi); err != nil {
		return err
	}

	data, err := json.Marshal(vi)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}

func NewStruct(v interface{}) (rt reflect.Type, err error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		err = &json.InvalidUnmarshalError{reflect.TypeOf(v)}
		return
	}

	rt = newStruct(rv.Elem().Type())
	return
}

func newStruct(rt reflect.Type) reflect.Type {
	rs := make([]reflect.StructField, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)

		if f.Type.String() == "*time.Time" {
			f.Type = rtt
		} else {
			switch f.Type.Kind() {
			case reflect.Array:
				f.Type = reflect.ArrayOf(f.Type.Len(), newStruct(f.Type.Elem()))
			case reflect.Chan:
				f.Type = reflect.ChanOf(f.Type.ChanDir(), newStruct(f.Type.Elem()))
			case reflect.Func:
				ins := make([]reflect.Type, f.Type.NumIn())
				for i := 0; i < f.Type.NumIn(); i++ {
					ins[i] = newStruct(f.Type.In(i))
				}
				outs := make([]reflect.Type, f.Type.NumOut())
				for i := 0; i < f.Type.NumOut(); i++ {
					outs[i] = newStruct(f.Type.Out(i))
				}
				f.Type = reflect.FuncOf(ins, outs, f.Type.IsVariadic())
			case reflect.Interface:
				// TODO
			case reflect.Map:
				f.Type = reflect.MapOf(newStruct(f.Type.Key()), newStruct(f.Type.Elem()))
			case reflect.Ptr:
				f.Type = reflect.PtrTo(newStruct(f.Type.Elem()))
			case reflect.Slice:
				f.Type = reflect.SliceOf(newStruct(f.Type.Elem()))
			case reflect.Struct:
				f.Type = newStruct(f.Type)
			case reflect.UnsafePointer:
				// TODO
			}
		}

		rs[i] = f
	}
	return reflect.StructOf(rs)
}
