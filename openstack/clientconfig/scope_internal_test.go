package clientconfig

import "testing"

func TestBuildOIDCScopePreservesTrust(t *testing.T) {
	scope := buildOIDCScope(&AuthInfo{
		TrustID:     "trust-id",
		ProjectID:   "project-id",
		SystemScope: "all",
		DomainID:    "domain-id",
		ProjectName: "project-name",
		DomainName:  "domain-name",
	})
	if scope.TrustID != "trust-id" {
		t.Fatalf("unexpected trust scope: %#v", scope)
	}
	if scope.ProjectID != "" || scope.ProjectName != "" || scope.DomainID != "" || scope.DomainName != "" || scope.System {
		t.Fatalf("trust scope must not include another scope target: %#v", scope)
	}
}
