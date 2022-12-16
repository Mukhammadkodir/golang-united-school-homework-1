package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {

	m := emoji.Sprint("Hello :world_map:!")
	return m

}

func Out(){
	fmt.Println(GetMessage())
}