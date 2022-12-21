package solution

import (
	"testing"
)

func Test_GetMessage(T *testing.T){

	res := GetMessage()
	if res != "Hello 🗺️ !" {
		T.Fail()
	}

}