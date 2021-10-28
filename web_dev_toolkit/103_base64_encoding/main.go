package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := `To be, or not to be, that is the question:
	Whether 'tis nobler in the mind to suffer
	The slings and arrows of outrageous fortune,
	Or to take arms against a sea of troubles
	And by opposing end them. To die—to sleep,
	No more; and by a sleep to say we end
	The heart-ache and the thousand natural shocks
	That flesh is heir to: 'tis a consummation
	Devoutly to be wish'd. To die, to sleep;
	To sleep, perchance to dream—ay, there's the rub:
	For in that sleep of death what dreams may come`
	// custom encoding
	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)

	// default encoding

	s64 = base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("here is the standard encoding provided by default")
	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}
