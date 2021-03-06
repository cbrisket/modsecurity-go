package seclang

import (
	"strings"
	"testing"

	"github.com/senghoo/modsecurity-go/modsecurity/operators"
	"github.com/senghoo/modsecurity-go/seclang/parser"
)

func TestMakeOperator(t *testing.T) {
	t.Run("test @rx without @rx", func(t *testing.T) {
		dr := newDireRule()
		input := `"some regex"`
		scan := parser.NewSecLangScanner(strings.NewReader(input))
		parsed, err := scan.ReadOperator()
		if err != nil {
			t.Error(err)
			return
		}
		err = dr.applyOperator(parsed)
		if err != nil {
			t.Error(err)
			return
		}
		if dr.rule.Operator == nil {
			t.Error("not get operator")
			return
		}
		rx, ok := dr.rule.Operator.(*operators.OperatorRx)
		if !ok {
			t.Errorf("except VariableRequestURI got %#v", rx)
			return
		}
		if rx.Args() != "some regex" {
			t.Errorf("expect argument %s got %s", "some regex", rx.Args())
			return
		}

	})

	t.Run("test @rx", func(t *testing.T) {
		dr := newDireRule()
		input := `"@rx nikto"`
		scan := parser.NewSecLangScanner(strings.NewReader(input))
		parsed, err := scan.ReadOperator()
		if err != nil {
			t.Error(err)
			return
		}
		err = dr.applyOperator(parsed)
		if err != nil {
			t.Error(err)
			return
		}
		if dr.rule.Operator == nil {
			t.Error("not get operator")
			return
		}
		rx, ok := dr.rule.Operator.(*operators.OperatorRx)
		if !ok {
			t.Errorf("except VariableRequestURI got %#v", rx)
			return
		}
		if rx.Args() != "nikto" {
			t.Errorf("expect argument %s got %s", "nikto", rx.Args())
			return
		}

	})

}
