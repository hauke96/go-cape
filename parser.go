package cape

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type parser struct {
	args           map[string]*argument
	longToShortArg map[string]string
	requiredArgs   []*argument
	KnownLongArgs  string
	KnownShortArgs string
	description    string
}

// NewParser creates an empty parser with no arguments.
func NewParser() *parser {
	p := parser{
		args:           make(map[string]*argument),
		longToShortArg: make(map[string]string),
		KnownLongArgs:  ":",
		KnownShortArgs: ":",
		description:    "",
	}
	return &p
}

// Description sets a new description for further information at the help-text.
func (p *parser) Description(newDescription string) {
	p.description = newDescription
}

// RegisterArgument creates a new argument based on the parameters given.
// Now allowed is an empty shortKey and existing long/shortKeys
func (p *parser) RegisterArgument(longKey, shortKey, help string) *argument {
	// the long key is allowed to be empty, the short key not
	if shortKey == "" {
		return nil
	}
	// when the long key and a fitting entry exists, we've a duplicate here
	if longKey != "" && p.longToShortArg[longKey] != "" {
		return nil
	}
	// when there's a short key entry already, we've a duplicate here
	if p.args[shortKey] != nil {
		return nil
	}

	a := newArgument()
	a.longKey = longKey
	a.shortKey = shortKey
	a.helpText = help

	p.KnownShortArgs += shortKey + ":"
	p.KnownLongArgs += longKey + ":"
	p.longToShortArg[longKey] = shortKey

	p.args[shortKey] = a
	return a
}

// Parse goes through the arguments (from 1 to n, so the first one is skiped) and sets the values of the arguments
// This function also takes the arguments and parses them into two categories: normal and predefinings arguments.
// It also evaluates the predefining ones.
func (p *parser) Parse() {
	p.parseArgs(os.Args[1:])
}
func (p *parser) parseArgs(args []string) {
	if len(args) == 0 {
		return
	}

	// ------------------------------
	// CHECK FOR -h/--help
	// ------------------------------
	if args[0] == "--help" || args[0] == "-h" {
		p.ShowHelp()
		os.Exit(1)
	}

	// ------------------------------
	// GET ALL REQUIRED ARGS
	// ------------------------------
	for _, arg := range p.args {
		if arg.required {
			p.requiredArgs = append(p.requiredArgs, arg)
		}
	}

	// ------------------------------
	// SPLIT COMBINED ARGS
	// ------------------------------
	// E.g. [-bgh] -> [-b] [-g] [-h]
	args = p.splitCombinedArgs(args)

	// ------------------------------
	// SEPARATE INTO NORMAL AND PREDEFINING
	// ------------------------------
	invalidArgExists := false
	for _, arg := range args {
		splittedArg := strings.Split(arg, "=")
		key, err := p.truncateArg(splittedArg[0])
		if err != nil {
			fmt.Println(err.Error())
			invalidArgExists = true
			break
		}

		key, err = p.toShortKey(key)
		if err != nil {
			fmt.Println(err.Error())
			invalidArgExists = true
			break
		}

		// the argument if one of the registered ones
		if strings.Contains(p.KnownShortArgs, ":"+key+":") ||
			strings.Contains(p.KnownLongArgs, ":"+key+":") { // is it a valid short or long argument?

			arg := p.args[key]
			if len(splittedArg) >= 2 { // argument with value
				arg.set(splittedArg[1])
			} else { // argument without value (=flag)
				arg.set("true")
			}

			if arg.required {
				p.requiredArgs = remove(p.requiredArgs, arg)
			}

		} else { // not valid
			if len(key) == 1 { // just to have the - in from of the argument (a bit more pretty ;) )
				key = "-" + key
			}
			invalidArgExists = true
			break
		}
	}
	if invalidArgExists {
		p.ShowHelp()
		os.Exit(1)
	}
}

// splitCombinedArgs splitts combined args :o
// A combined arg is something like -bgl which will be split into -b -g -l
func (p *parser) splitCombinedArgs(args []string) []string {
	newArgs := make([]string, 0)
	for _, arg := range args {
		// !strings.Contains(arg, "=")		- A = is not allowed because -bgl are all flags/short args
		// arg[0] == '-' && arg[1] != '-' 	- There's only one - allowed
		// len(arg) > 2						- Only string of the form -[a-Z0-9] are valid
		if !strings.Contains(arg, "=") && arg[0] == '-' && arg[1] != '-' && len(arg) > 2 { // 3 because of - and at least 2 other characters

			arg = arg[1:] // remove the -

			for _, v := range arg {
				newArgs = append(newArgs, "-"+string(v))
			}
		} else {
			newArgs = append(newArgs, arg)
		}
	}
	return newArgs
}

// truncateArg removes the dashes in the beginning of an argument and returns the plain key.
// If the argument wasn't valid (like -test or --t ) the original argument and an error will be returned.
func (p *parser) truncateArg(arg string) (string, error) {
	// notice: only --foo or -f are allowed, -foo and --f are not allowed!
	// This is just to have the normal feeling of arguments in linux, blame me but i like it ;)
	if arg[0:2] == "--" {
		if len(arg) > 3 { // greater 3 because of -- (two dashes) and at least 2 other characters
			return arg[2:], nil
		} else { // something like --f which is invalid
			return arg, errors.New("The argument " + arg + " is not valid! Something like " + arg[1:] + " would be ok.")
		}
	} else if len(arg) == 2 { // - and one other character
		arg = arg[1:]
	} else if arg == "-" {
		return arg, errors.New("What? The argument " + arg + " doesn't make any sense.")
	} else {
		return arg, errors.New("I have no idea what " + arg + " means, sorry.")
	}

	return arg, nil
}

// toShortKey takes any key and returns the short version of this. Not allowed are keys that begin with a dash!
func (p *parser) toShortKey(key string) (string, error) {
	if key[0] == '-' {
		return key, errors.New("Key beginning with a dash is not allowed here!")
	}
	if len(key) > 1 { // not a short key -> convert
		key = p.longToShortArg[key]
		if key == "" {
			return key, errors.New("There's no entry for the key " + key + ".")
		}
	}
	return key, nil
}

func (p *parser) ShowHelp() {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 4, 2, ' ', 0)
	defer writer.Flush()

	fmt.Fprintln(writer, "Usage: <argument>=<value> <flag>")
	fmt.Fprintln(writer, "")
	fmt.Fprintln(writer, "Here're all command that are available:\n")

	requiredExists := false

	for _, arg := range p.args {
		fmt.Fprint(writer, "\t-"+arg.shortKey+", --"+arg.longKey+"\t")
		if arg.required {
			fmt.Fprint(writer, "*")
			requiredExists = true
		}
		fmt.Fprintln(writer, "\t"+arg.helpText)
	}

	fmt.Fprintln(writer, "\t-h, --help\t\tShows this help message")
	fmt.Fprintln(writer, "")
	if requiredExists {
		fmt.Fprintln(writer, "Commands with a * are required and have to be set.")
		fmt.Fprintln(writer, "")
	}
	if p.description != "" {
		fmt.Fprintln(writer, p.description)
	}
	fmt.Fprintln(writer, "")
}

func remove(a []*argument, e *argument) []*argument {
	i := 0
	for j, v := range a {
		if v == e {
			i = j
			break
		}
	}

	if i != 0 && i+1 < len(a) {
		a = append(a[:i], a[i+1:]...)
	}
	return a
}
