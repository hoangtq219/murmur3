package test

import (
	"fmt"
	"testing"
)

func SwitchTest()  {
	switch 4 {
	case 15:
		fmt.Println(15)
	case 14:
		fmt.Println(14)
	case 13:
		fmt.Println(15)
	case 12:
		fmt.Println(15)
	case 11:
		fmt.Println(15)
	case 10:
		fmt.Println(15)
	case 9:
		fmt.Println(15)
	case 8:
		fmt.Println(15)
	case int(7):
		fmt.Println(15)
	case 6:
		fmt.Println(15)
	case 5:
		fmt.Println(15)
	case int(4):
		fmt.Println(15)
	case 3:
		fmt.Println(15)
	case 2:
		fmt.Println(15)
	case 1:
		fmt.Println(15)
	}
}

func testPutByte(text string)  {

	fmt.Println("Len(data) =", len(text))
	fmt.Println([]byte(text))

	data := make([]byte, 23)

	for i := 0; i < len(text); i++ {
		index := (i%8) * 2
		data[index] = text[i]
		data[index + 1] = text[i] >> 8
		if i == 7 || i == 15  {
			fmt.Println(data)
		}
	}

	fmt.Println(data)
	fmt.Println(len(data))
}

func TestAddInt(t *testing.T) {
	//value := m.bb.remaining()
	value  := 10
	switch value {
	case 15:
		fmt.Println(15)
	case 14:
		fmt.Println(14)
	case 13:
		fmt.Println(13)
	case 12:
		fmt.Println(12)
	case 11:
		fmt.Println(11)
	case 10:
		fmt.Println(10)
	case 9:
		fmt.Println(9)
	case 8:
		fmt.Println(8)
	case 7:
		fmt.Println(7)
	case 6:
		fmt.Println(6)
	case 5:
		fmt.Println(5)
	case 4:
		fmt.Println(4)
	case 3:
		fmt.Println(3)
	case 2:
		fmt.Println(2)
	case 1:
		fmt.Println(1)
	}
}


//func main() {
//	//TestAddInt()
//	//testPutByte("681236075540516864")
//}
