package clientconfig

import (
	"context"
	"crypto/tls"
	"testing"
)

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

func TestBuildOIDCScopeDoesNotMutateAuthInfo(t *testing.T) {
	authInfo := &AuthInfo{
		ProjectName: "project-name",
		DomainID:    "domain-id",
	}

	scope := buildOIDCScope(authInfo)
	if scope.ProjectName != "project-name" || scope.DomainID != "domain-id" {
		t.Fatalf("unexpected project scope: %#v", scope)
	}
	if authInfo.DomainID != "domain-id" || authInfo.ProjectDomainID != "" || authInfo.UserDomainID != "" {
		t.Fatalf("buildOIDCScope mutated AuthInfo: %#v", authInfo)
	}
}

func TestSpecializedAuthDoesNotMutateClientOpts(t *testing.T) {
	t.Setenv("OS_CLIENT_ID", "client-id")
	t.Setenv("OS_IDENTITY_PROVIDER", "idp")
	t.Setenv("OS_PROTOCOL", "openid")

	authInfo := &AuthInfo{
		AuthURL:             "http://127.0.0.1:1/v3/",
		AccessTokenEndpoint: "http://127.0.0.1:1/token",
	}
	opts := &ClientOpts{
		AuthType: AuthV3OIDCClientCredentials,
		AuthInfo: authInfo,
	}

	_, _ = newOIDCProviderClient(context.Background(), new(Cloud), opts, "OS_", new(tls.Config))
	if authInfo.ClientID != "" || authInfo.IdentityProvider != "" || authInfo.Protocol != "" {
		t.Fatalf("specialized authentication mutated ClientOpts.AuthInfo: %#v", authInfo)
	}
}
