package _randomString

import (
	"fmt"
	"testing"
)

func TestRandomString_Get(t *testing.T) {

	fmt.Println(New().Get())
	fmt.Println(New().Get())
	fmt.Println(New().Get())

	fmt.Println(New().Char("0123456789").Get())
	fmt.Println(New().Char("0123456789").Get())
	fmt.Println(New().Char("0123456789").Get())

	fmt.Println(New().Char("qwertyuioplkjhgfdsazxcvbnm").Get())
	fmt.Println(New().Char("qwertyuioplkjhgfdsazxcvbnm").Get())
	fmt.Println(New().Char("qwertyuioplkjhgfdsazxcvbnm").Get())

	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNM").Get())
	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNM").Get())
	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNM").Get())

	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNMqwertyuioplkjhgfdsazxcvbnm0123456789").Get())
	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNMqwertyuioplkjhgfdsazxcvbnm0123456789").Get())
	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNMqwertyuioplkjhgfdsazxcvbnm0123456789").Get())

	fmt.Println(New().Char("0123456789").Size(4).Get())
	fmt.Println(New().Char("0123456789").Size(4).Get())
	fmt.Println(New().Char("0123456789").Size(4).Get())

	fmt.Println(New().Char("qwertyuioplkjhgfdsazxcvbnm").Size(4).Get())
	fmt.Println(New().Char("qwertyuioplkjhgfdsazxcvbnm").Size(4).Get())
	fmt.Println(New().Char("qwertyuioplkjhgfdsazxcvbnm").Size(4).Get())

	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNM").Size(4).Get())
	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNM").Size(4).Get())
	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNM").Size(4).Get())

	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNMqwertyuioplkjhgfdsazxcvbnm0123456789").Size(4).Get())
	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNMqwertyuioplkjhgfdsazxcvbnm0123456789").Size(4).Get())
	fmt.Println(New().Char("QWERTYUIOPLKJHGFDSAZXCVBNMqwertyuioplkjhgfdsazxcvbnm0123456789").Size(4).Get())

}
