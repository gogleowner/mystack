package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	//example_1_use_unicode1()
	example_2_use_utf8()
}

/*
rune :
- rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.
- 유니코드(UTF-8)를 표현할 때 사용하는 형식 rune을 제공하고 있다.
- java로 비교하자면 char 이라고 생각하면 될듯.
*/
func example_1_use_unicode1() {
	var r1 rune = '한'
	fmt.Println(unicode.Is(unicode.Hangul, r1)) // 한글이면 true
	fmt.Println(unicode.Is(unicode.Latin, r1))

	fmt.Println(unicode.In(r1, unicode.Latin, unicode.Hangul))

	fmt.Println(unicode.IsGraphic('1')) // 화면에 표시되는 문자
	fmt.Println(unicode.IsGraphic('a'))
	fmt.Println("unicode.IsGraphic(1) => ", unicode.IsGraphic(1)) // 유니코드가 아니라서 안나오는것 같다.
	fmt.Println("unicode.IsGraphic('\ud55c') => ", unicode.IsGraphic('\ud55c'))
	fmt.Println(unicode.IsGraphic('\n'))

	fmt.Println(unicode.IsLetter('a')) // a는 문자이므로 true
	fmt.Println(unicode.IsLetter(1))   // 1은 숫자이므로 false

	fmt.Println(unicode.IsDigit('1'))     // 1은 숫자이므로 true
	fmt.Println(unicode.IsControl('\n'))  // \n은 제어문자이므로 true
	fmt.Println(unicode.IsMark('\u17c9')) // \u17c9은 마크이므로 true

	fmt.Println(unicode.IsPrint('1')) // 1은 Go언어에서 표시할 수 있으므로 true
	fmt.Println(unicode.IsPunct('.')) // .은 문장 부호(punctuation)이므로 true
}

func example_2_use_utf8() {
	var s string = "한"
	var r rune = '한'

	fmt.Println("s ->", len(s), "/ r ->", utf8.RuneLen(r)) // 한글 글자의 바이트수는 3바이트

	var hello string = "안녕하세요"
	fmt.Println(utf8.RuneCountInString(hello))

	b := []byte("안녕하세요")
	rb, size := utf8.DecodeRune(b)
	fmt.Printf("%c %d\n", rb, size)

	r2, size2 := utf8.DecodeRune(b[3:])
	fmt.Printf("%c %d\n", r2, size2)

	r3, size3 := utf8.DecodeLastRune(b)
	fmt.Printf("%c %d\n", r3, size3)

	r4, size4 := utf8.DecodeLastRune(b[:len(b)-3])
	fmt.Printf("%c %d\n", r4, size4)

	// iteration
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)
		b = b[size:]
	}

	fmt.Println(utf8.Valid(b)) // 안녕하세요 는 utf8이 맞으므로 true
	fmt.Println(utf8.Valid([]byte{0xff, 0xf1, 0xc1})) // 0xff, 0xf1, 0xc1은 utf8이 아니므로 false

	var han rune = '한'
	fmt.Println(utf8.ValidRune(han))
	var han2 rune = 0x11111111
	fmt.Println(utf8.ValidRune(han2)) // 0x11111111은 utf8이 아니므로 false
}

