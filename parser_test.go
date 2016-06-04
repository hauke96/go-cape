package cape

import (
	"testing"
)

func TestAddArgument(t *testing.T) {
	p := NewParser()
	len1 := len(p.args)
	p.RegisterArgument("foo", "f", "")
	len2 := len(p.args)

	if len1 == len2 {
		t.Error("The argument has not been added.")
	}
}

func TestParser(t *testing.T) {
	p := NewParser()

	integerArg := p.RegisterArgument("foo", "f", "A sample argument").Required().Int()
	stringArg := p.RegisterArgument("bar", "b", "A second sample argument").String()
	boolArg := p.RegisterArgument("short", "s", "And another sample argument").Bool()
	flagArg := p.RegisterArgument("flag", "l", "A sample flag").Bool()

	args := make([]string, 4)
	args[0] = "-f=123"
	args[1] = "--bar=hello world!"
	args[2] = "-s=true"
	args[3] = "-l"

	p.parseArgs(args)

	if integerArg == nil || *integerArg != 123 {
		t.Error("The integerArg has not been parsed correctly.")
	}
	if stringArg == nil || *stringArg != "hello world!" {
		t.Error("The stringArg has not been parsed correctly.")
	}
	if boolArg == nil || *boolArg != true {
		t.Error("The boolArg has not been parsed correctly.")
	}
	if flagArg == nil || *flagArg != true {
		t.Error("The flagArg has not been parsed correctly.")
	}
}

func TestArgumentSetValue(t *testing.T) {
	a := NewParser().RegisterArgument("foo", "f", "")
	var v1, v2 int

	v1 = *a.intValue
	a.set("42")
	v2 = *a.intValue

	if v1 == v2 {
		t.Error("The settings failed: both values v1(", v1, ",", &v1, ") and v2(", v2, ",", &v2, ") are equal :(")
	}
}

func TestLongToShortArgument(t *testing.T) {
	p := NewParser()

	longArgs := make([]string, 4)
	shortArgs := make([]string, len(longArgs))

	longArgs[0] = "foo"
	longArgs[1] = "bar"
	longArgs[2] = "short"
	longArgs[3] = "--foo"

	shortArgs[0] = "f"
	shortArgs[1] = "b"
	shortArgs[2] = "s"
	shortArgs[3] = "-f"

	for i, longArg := range longArgs {
		p.RegisterArgument(longArg, shortArgs[i], "").String()
	}
	for i, longArg := range longArgs {
		shortA, errA := p.toShortKey(longArg)
		shortB, errB := p.toShortKey(shortArgs[i])
		if shortA != shortArgs[i] {
			if errA == nil {
				t.Error("Wrong convertion of", longArg, "to", shortA+". Error: ", errA)
			}
		}
		if shortB != shortArgs[i] {
			if errA == nil {
				t.Error("Wrong convertion of", shortArgs[i], "to", shortB+". Error: ", errB)
			}
		}
	}
}

func TestTruncateArgument(t *testing.T) {
	p := NewParser()

	args := make([]string, 5)

	args[0] = "--foo"
	args[1] = "--b"
	args[2] = "-s"
	args[3] = "-short"
	args[4] = "whatever"

	i := 0
	if trunc, err := p.truncateArg(args[i]); trunc != "foo" || err != nil {
		t.Error("Truncating of", args[i], "failed. Error:", err)
	}
	i++
	if trunc, err := p.truncateArg(args[i]); trunc != args[i] || err == nil {
		t.Error("Truncating of", args[i], "failed. Error:", err)
	}
	i++
	if trunc, err := p.truncateArg(args[i]); trunc != "s" || err != nil {
		t.Error("Truncating of", args[i], "failed. Error:", err)
	}
	i++
	if trunc, err := p.truncateArg(args[i]); trunc != args[i] || err == nil {
		t.Error("Truncating of", args[i], "failed. Error:", err)
	}
	i++
	if trunc, err := p.truncateArg(args[i]); trunc != args[i] || err == nil {
		t.Error("Truncating of", args[i], "failed. Error:", err)
	}
}

func TestSplitCombinedArgument(t *testing.T) {
	p := NewParser()

	args := make([]string, 4)

	args[0] = "-foo"
	args[1] = "--bar"
	args[2] = "-s"
	args[3] = "whatever"

	expected := make([]string, 6)
	expected[0] = "-f"
	expected[1] = "-o"
	expected[2] = "-o"
	expected[3] = "--bar"
	expected[4] = "-s"
	expected[5] = "whatever"

	result := p.splitCombinedArgs(args)

	for i, resultValue := range result {
		if resultValue != expected[i] {
			t.Error("Different values at index", i, ". Expected", expected[i], "but got", resultValue)
		}
	}
}
