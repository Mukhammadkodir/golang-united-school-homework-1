package solution

import (
	"testing"
)

func TestGetMessage(T *testing.T){

	res := GetMessage()
	if res != "Hello 🗺️ !" {
		T.Fail()
	}

}