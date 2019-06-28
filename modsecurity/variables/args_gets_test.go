package variables

import (
	"net/url"
	"testing"

	"github.com/senghoo/modsecurity-go/modsecurity"
	"github.com/senghoo/modsecurity-go/utils"
)

func TestVariableArgsGet(t *testing.T) {
	v := NewVariableArgsGet()
	v.Include(`/a/`)
	u, _ := url.Parse("http://localhost/query?a1=1&a2=2&b1=3&b2=4")
	tr, err := modsecurity.NewTransaction(modsecurity.NewEngine(), modsecurity.NewSecRuleSet())
	if err != nil {
		t.Error(err)
		return
	}
	tr.ProcessRequestURL(u, "GET", "HTTP/1.1")
	res := v.Fetch(tr)
	if !utils.SameStringSlice(res, []string{"1", "2"}) {
		t.Errorf("variable args get fail got %q", res)
	}
}