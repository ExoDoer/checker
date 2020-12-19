package _checker

import (
	"testing"

	"github.com/liangyaopei/checker"
)

type enum struct {
	Int    int
	Uint   uint
	Float  float64
	String string
}

func TestEnumTest(t *testing.T) {
	eChecker := checker.NewChecker()

	intRule := checker.NewEnumRuleInt("Int", 1, 2, 3)
	eChecker.Add(intRule, "invalid Int")

	uintRule := checker.NewEnumRuleUint("Uint", 4, 5, 6)
	eChecker.Add(uintRule, "invalid Uint")

	floatRule := checker.NewEnumRuleFloat("Float", 3.14, 6.28, 9.42, 18.0)
	eChecker.Add(floatRule, "invalid Float")

	strRule := checker.NewEnumRuleString("String", "github", "str")
	eChecker.Add(strRule, "invalid String")

	e := enum{
		Int:    3,
		Uint:   5,
		Float:  3.14,
		String: "str",
	}

	isValid, prompt, errMsg := eChecker.Check(e)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid cross enum")

}
