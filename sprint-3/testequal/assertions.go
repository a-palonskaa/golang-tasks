//go:build !solution

package testequal

import (
	"bytes"
	"fmt"
	"reflect"
)

func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()

	if checkIfEqual(expected, actual) {
		return true
	}

	t.Errorf(getMessage("not equal", expected, actual, msgAndArgs...))
	return false
}

func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()

	if !checkIfEqual(expected, actual) {
		return true
	}

	t.Errorf(getMessage("equal", expected, actual, msgAndArgs...))
	return false
}

func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()

	if !AssertEqual(t, expected, actual, msgAndArgs...) {
		t.FailNow()
	}
}

func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()

	if !AssertNotEqual(t, expected, actual, msgAndArgs...) {
		t.FailNow()
	}
}

func checkIfEqual(obj1, obj2 interface{}) bool {
	obj1Value := reflect.ValueOf(obj1)
	obj2Value := reflect.ValueOf(obj2)

	if obj1Value.Kind() != obj2Value.Kind() {
		return false
	}

	switch kind := obj1Value.Kind(); kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return obj1Value.Int() == obj2Value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return obj1Value.Uint() == obj2Value.Uint()
	case reflect.String:
		return obj1.(string) == obj2.(string)
	case reflect.Slice:
		{
			if obj1Value.IsNil() != obj2Value.IsNil() {
				return false
			}

			if obj1Value.Len() != obj2Value.Len() {
				return false
			}

			obj1ElemType := obj1Value.Type().Elem()
			obj2ElemType := obj2Value.Type().Elem()

			if obj1ElemType.Kind() != obj2ElemType.Kind() {
				return false
			}

			switch obj1ElemType.Kind() {
			case reflect.Uint8:
				return bytes.Equal(obj1Value.Bytes(), obj2Value.Bytes())
			case reflect.Int:
				for i := 0; i < obj1Value.Len(); i++ {
					if obj1Value.Index(i).Int() != obj2Value.Index(i).Int() {
						return false
					}
				}
				return true
			default:
				return false
			}
		}
	case reflect.Map:
		{
			if obj1Value.IsNil() != obj2Value.IsNil() {
				return false
			}

			obj1KeyType := obj1Value.Type().Key().Kind()
			obj2KeyType := obj2Value.Type().Key().Kind()

			if (obj1KeyType != obj2KeyType) || (obj1KeyType != reflect.String) {
				return false
			}

			obj1ElemType := obj1Value.Type().Elem().Kind()
			obj2ElemType := obj2Value.Type().Elem().Kind()

			if (obj1ElemType != obj2ElemType) || (obj1ElemType != reflect.String) {
				return false
			}

			if obj1Value.Len() != obj2Value.Len() {
				return false
			}

			for _, key := range obj1Value.MapKeys() {
				if obj2Value.MapIndex(key).String() != obj1Value.MapIndex(key).String() {
					return false
				}
			}
			return true
		}
	default:
		return false
	}
}

func getMessage(str string, expected, actual interface{}, msgAndArgs ...interface{}) string {
	message := fmt.Sprintf(str+": \nexpected: %v\nactual  : %v", expected, actual)
	if len(msgAndArgs) > 0 {
		if msg, ok := msgAndArgs[0].(string); ok {
			message += "\nmessage : " + fmt.Sprintf(msg, msgAndArgs[1:]...)
		}
	}
	return message
}
