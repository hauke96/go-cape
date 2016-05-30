package main

import (
	"fmt"
	"github.com/hauke96/go-cape"
	"os"
)

// call this with the following arguments or leave them out to use the ones I set in the source code
// -a=123 --bbb=456 -c -d=true --eee=hello

// a (s one letter) is the short argument aaa (three letters) the long version

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-a=123")
		os.Args = append(os.Args, "--bbb=456")
		//		os.Args = append(os.Args, "-c") //you can also disable this
		os.Args = append(os.Args, "-d=true")
		os.Args = append(os.Args, "--eee=hello") //you can also disable this
	}

	parser := cape.NewParser()

	a := parser.RegisterArgument("aaa", "a", "A sAmple Argument.").Required().Int()
	b := parser.RegisterArgument("bbb", "b", "B is another sample argument.").Int()
	c := parser.RegisterArgument("ccc", "c", "See ... eh ... c? This can also be an argument.").Default("false").Required().Bool()
	d := parser.RegisterArgument("ddd", "d", "Dis is a boolean value.").Bool()
	e := parser.RegisterArgument("eee", "e", "Eehh ... some other argument i guess ...").Default("lool").String()

	parser.Parse()

	fmt.Println("aaa/a:", *a)
	fmt.Println("bbb/b:", *b)
	fmt.Println("ccc/c:", *c)
	fmt.Println("ddd/d:", *d)
	fmt.Println("eee/e:", *e)

	parser.ShowHelp()
}
