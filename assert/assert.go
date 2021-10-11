package assert

import (
	"fmt"
	"strconv"
	"strings"
)

type testcommon interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

func AssertEqual(t testcommon, a, b int) {
	t.Helper()
	if a != b {
		t.Errorf("%d != %d", a, b)
	}
}

func AssertEqualArr(t testcommon, a, b []int) {
	t.Helper()
	fmt.Println("AssertEqualArr:")
	fmt.Println(a)
	fmt.Println(b)
	length := len(a)
	if length != len(b) {
		t.Errorf("len(a): %d != len(b): %d", length, len(b))
		return
	}
	tips := make([]string, 0, length)
	fmt.Print(" ")
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			fmt.Print("^")
			tips = append(tips, fmt.Sprintf("a[%d]: %d != b[%d]: %d", i, a[i], i, b[i]))
		} else {
			fmt.Print(" ")
		}
		fmt.Print(strings.Repeat(" ", len(strconv.Itoa(b[i]))))
	}
	fmt.Println()
	if len(tips) != 0 {
		t.Error(fmt.Sprintf("\n%s", strings.Join(tips, "\n")))
	}
}

func AssertEqualMatrix(t testcommon, a, b [][]int) {
	t.Helper()
	fmt.Println("AssertEqualMatrix:")
	fmt.Println(a)
	fmt.Println(b)
	if len(a) != len(b) {
		t.Errorf("len(a): %d != len(b): %d", len(a), len(b))
		return
	}
	var tips []string
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			for j := 0; j < len(b[i]); j++ {
				fmt.Print(strings.Repeat(" ", len(strconv.Itoa(b[i][j]))+1))
			}
			fmt.Println(" ^")
			t.Errorf("len(a[%d]): %d != len(b[%d]): %d", i, len(a[i]), i, len(b[i]))
			return
		}
		fmt.Print("  ")
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				fmt.Print("^")
				tips = append(tips, fmt.Sprintf("a[%d][%d]: %d != b[%d][%d]: %d", i, j, a[i][j], i, j, b[i][j]))
			} else {
				fmt.Print(" ")
			}
			fmt.Print(strings.Repeat(" ", len(strconv.Itoa(b[i][j]))))
		}
	}
	fmt.Println()
	if len(tips) != 0 {
		t.Error(fmt.Sprintf("\n%s", strings.Join(tips, "\n")))
	}
}
