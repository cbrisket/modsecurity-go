package variables

import (
	"net/url"
	"testing"

	"github.com/senghoo/modsecurity-go/modsecurity"
)

func TestVariableRequestFilename(t *testing.T) {
	v := NewVariableRequestFilename()
	if v.Name() != "REQUEST_FILENAME" {
		t.Errorf("got unexcepted variable name %s", v.Name())
		return
	}
	inputs := map[string]string{
		"http://localhost/query?a1=1&a2=2&b1=3&b2=4":         "/query",
		"http://localhost/query.php?a1=1&a2=2&b1=3&b2=4":     "/query.php",
		"http://localhost/query.php?a1=1&a2=2&b1=3&b2=4#123": "/query.php",
	}
	for input, out := range inputs {
		u, _ := url.Parse(input)
		tr, err := modsecurity.NewTransaction(modsecurity.NewEngine(), modsecurity.NewSecRuleSet())
		if err != nil {
			t.Error(err)
			return
		}
		tr.ProcessRequestURL(u, "GET", "HTTP/1.1")
		vars := v.Fetch(tr)
		if len(vars) != 1 {
			t.Errorf("unexcepted count %d", len(vars))
			return
		}
		if vars[0] != out {
			t.Errorf("variable args get fail got %q", vars)
		}
	}
}
