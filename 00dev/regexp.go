package main

import (
	"fmt"
	"regexp"
)

func Testmain9() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("mira"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("mira eka edo eco eto eyo eki", 10)
	fmt.Println(search)

}
