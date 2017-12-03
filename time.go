package tson

import (
	"reflect"
	"sync"
	"time"
)

var format = struct {
	layout string
	mutex  *sync.Mutex
}{
	time.RFC3339,
	new(sync.Mutex),
}

func SetLayout(layout string) {
	format.mutex.Lock()
	format.layout = layout
	format.mutex.Unlock()
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	tm, err := time.Parse(`"`+format.layout+`"`, string(data))
	t.Time = tm
	return err
}

var rtt = reflect.TypeOf(&Time{})
