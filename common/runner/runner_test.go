package runner

import (
	"github.com/Tencent/AI-Infra-Guard/internal/gologger"
	"github.com/Tencent/AI-Infra-Guard/internal/options"
	"testing"
)

func TestRunner_RunEnumeration(t *testing.T) {
	targets := []string{
		"http://127.0.0.1:5000",
	}
	parseOptions := &options.Options{
		Target:       targets,
		Output:       "",
		ProxyURL:     "",
		TimeOut:      10,
		JSON:         false,
		RateLimit:    10,
		FPTemplates:  "data/fingerprints",
		AdvTemplates: "data/advisories",
	}
	r, err := New(parseOptions)
	if err != nil {
		gologger.Fatalf("Could not create runner: %s\n", err)
	}
	defer r.Close()
	r.RunEnumeration()
}
