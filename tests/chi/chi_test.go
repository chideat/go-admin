package chi

import (
	"github.com/chenhg5/go-admin/tests/common"
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func TestChi(t *testing.T) {
	common.Test(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(NewChiHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
