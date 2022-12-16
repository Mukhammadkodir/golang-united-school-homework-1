package solution

import (
	"testing"
)

func TestMain(T *testing.T) {

	res := GetMessage()
	if res != "Hello 🗺️ !" {
		T.Fail()
	}

}
