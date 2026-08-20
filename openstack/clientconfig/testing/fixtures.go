package testing

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/utils/v2/openstack/clientconfig"
)

var iTrue = true
var iFalse = false

var VirginiaEnvAuth = map[string]string{
	"OS_AUTH_URL":                      "https://va.example.com:5000/v3",
	"OS_APPLICATION_CREDENTIAL_ID":     "app-cred-id",
	"OS_APPLICATION_CREDENTIAL_SECRET": "secret",
}

var VirginiaAuthOpts = &gophercloud.AuthOptions{
	Scope:                       &gophercloud.AuthScope{},
	IdentityEndpoint:            "https://va.example.com:5000/v3",
	ApplicationCredentialID:     "app-cred-id",
	ApplicationCredentialSecret: "secret",
}

var VirginiaCloudYAML = clientconfig.Cloud{
	RegionName: "VA",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:                     "https://va.example.com:5000/v3",
		ApplicationCredentialID:     "app-cred-id",
		ApplicationCredentialSecret: "secret",
	},
	Verify:   &iTrue,
	AuthType: "v3applicationcredential",
}

var PhiladelphiaCloudYAML = clientconfig.Cloud{
	RegionName: "PHL",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://phl.example.com:5000/v3",
		Username:    "admin",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Verify:   &iTrue,
	AuthType: "password",
}

var PhiladelphiaComplexPhl1CloudYAML = clientconfig.Cloud{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://phl1.example.com:5000/v3",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Regions: []clientconfig.Region{
		{
			Name:   "PHL1",
			Values: clientconfig.Cloud{AuthInfo: &clientconfig.AuthInfo{AuthURL: "https://phl1.example.com:5000/v3"}},
		},
		{
			Name:   "PHL2",
			Values: clientconfig.Cloud{},
		},
	},
	Verify: &iTrue,
}

var PhiladelphiaComplexPhl2CloudYAML = clientconfig.Cloud{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://phl.example.com:5000/v3",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Regions: []clientconfig.Region{
		{
			Name:   "PHL1",
			Values: clientconfig.Cloud{AuthInfo: &clientconfig.AuthInfo{AuthURL: "https://phl1.example.com:5000/v3"}},
		},
		{
			Name:   "PHL2",
			Values: clientconfig.Cloud{},
		},
	},
	Verify: &iTrue,
}

var disconnected_regions = []clientconfig.Region{
	{
		Name:   "SOMEWHERE",
		Values: clientconfig.Cloud{AuthInfo: &clientconfig.AuthInfo{AuthURL: "https://somewhere.example.com:5000/v3"}},
	},
	{
		Name:   "ANYWHERE",
		Values: clientconfig.Cloud{AuthInfo: &clientconfig.AuthInfo{AuthURL: "https://anywhere.example.com:5000/v3"}},
	},
	{
		Name:   "NOWHERE",
		Values: clientconfig.Cloud{AuthInfo: &clientconfig.AuthInfo{AuthURL: "https://nowhere.example.com:5000/v3"}},
	},
}

var DisconnectedSomewhereCloudYAML = clientconfig.Cloud{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://somewhere.example.com:5000/v3",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Regions: disconnected_regions,
	Verify:  &iTrue,
}

var DisconnectedAnywhereCloudYAML = clientconfig.Cloud{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://anywhere.example.com:5000/v3",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Regions: disconnected_regions,
	Verify:  &iTrue,
}

var DisconnectedNowhereCloudYAML = clientconfig.Cloud{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://nowhere.example.com:5000/v3",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Regions: disconnected_regions,
	Verify:  &iTrue,
}

var ChicagoCloudYAML = clientconfig.Cloud{
	Profile:    "rackspace",
	RegionName: "ORD",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://identity.api.rackspacecloud.com/v2.0/",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Verify: &iTrue,
}

var ChicagoCloudLegacyYAML = clientconfig.Cloud{
	Cloud:      "rackspace",
	RegionName: "ORD",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://identity.api.rackspacecloud.com/v2.0/",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Verify: &iTrue,
}

var ChicagoCloudUseProfileYAML = clientconfig.Cloud{
	Cloud:      "rackspace",
	Profile:    "rackspace",
	RegionName: "ORD",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://identity.api.rackspacecloud.com/v2.0/",
		Username:    "jdoe",
		Password:    "securepassword",
		ProjectName: "Some Project",
	},
	Verify: &iTrue,
}

var HawaiiCloudYAML = clientconfig.Cloud{
	RegionName: "HNL",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://hi.example.com:5000/v3",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
		DomainName:  "default",
	},
	Verify: &iTrue,
}

var HawaiiClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://hi.example.com:5000/v3",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
		DomainName:  "default",
	},
}

var HawaiiEnvAuth = map[string]string{
	"OS_AUTH_URL":     "https://hi.example.com:5000/v3",
	"OS_USERNAME":     "jdoe",
	"OS_PASSWORD":     "password",
	"OS_PROJECT_NAME": "Some Project",
	"OS_DOMAIN_NAME":  "default",
}

var HawaiiAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectName: "Some Project",
		DomainName:  "default",
	},
	IdentityEndpoint: "https://hi.example.com:5000/v3",
	Username:         "jdoe",
	Password:         "password",
	TenantName:       "Some Project",
	DomainName:       "default",
}

var FloridaCloudYAML = clientconfig.Cloud{
	RegionName: "MIA",
	Interface:  "admin",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:      "https://fl.example.com:5000/v3",
		Username:     "jdoe",
		Password:     "password",
		ProjectID:    "12345",
		UserDomainID: "abcde",
	},
	Verify: &iTrue,
}

var InsecureFloridaCloudYAML = clientconfig.Cloud{
	RegionName: "MIA",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:      "https://fl.example.com:5000/v3",
		Username:     "jdoe",
		Password:     "password",
		ProjectID:    "12345",
		UserDomainID: "abcde",
	},
	Verify:         &iFalse,
	ClientKeyFile:  "",
	ClientCertFile: "",
	CACertFile:     "",
}

var SecureFloridaCloudYAML = clientconfig.Cloud{
	RegionName: "MIA",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:      "https://fl.example.com:5000/v3",
		Username:     "jdoe",
		Password:     "password",
		ProjectID:    "12345",
		UserDomainID: "abcde",
	},
	Verify:         &iTrue,
	ClientKeyFile:  "/home/myhome/client-cert.key",
	ClientCertFile: "/home/myhome/client-cert.crt",
	CACertFile:     "/home/myhome/ca.crt",
}

var FloridaClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:      "https://fl.example.com:5000/v3",
		Username:     "jdoe",
		Password:     "password",
		ProjectID:    "12345",
		UserDomainID: "abcde",
	},
}

var FloridaEnvAuth = map[string]string{
	"OS_AUTH_URL":       "https://fl.example.com:5000/v3",
	"OS_USERNAME":       "jdoe",
	"OS_PASSWORD":       "password",
	"OS_PROJECT_ID":     "12345",
	"OS_USER_DOMAIN_ID": "abcde",
}

var FloridaAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectID: "12345",
	},
	IdentityEndpoint: "https://fl.example.com:5000/v3",
	Username:         "jdoe",
	Password:         "password",
	TenantID:         "12345",
	DomainID:         "abcde",
}

var CaliforniaCloudYAML = clientconfig.Cloud{
	EndpointType: "internal",
	Regions: []clientconfig.Region{{
		Name: "SAN",
	}, {
		Name: "LAX",
	}},
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:           "https://ca.example.com:5000/v3",
		Username:          "jdoe",
		Password:          "password",
		ProjectName:       "Some Project",
		ProjectDomainName: "default",
		UserDomainName:    "default",
	},
	Verify: &iTrue,
}

var CaliforniaClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:           "https://ca.example.com:5000/v3",
		Username:          "jdoe",
		Password:          "password",
		ProjectName:       "Some Project",
		ProjectDomainName: "default",
		UserDomainName:    "default",
	},
}

var CaliforniaEnvAuth = map[string]string{
	"OS_AUTH_URL":            "https://ca.example.com:5000/v3",
	"OS_USERNAME":            "jdoe",
	"OS_PASSWORD":            "password",
	"OS_PROJECT_NAME":        "Some Project",
	"OS_PROJECT_DOMAIN_NAME": "default",
	"OS_USER_DOMAIN_NAME":    "default",
}

var CaliforniaAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectName: "Some Project",
		DomainName:  "default",
	},
	IdentityEndpoint: "https://ca.example.com:5000/v3",
	Username:         "jdoe",
	Password:         "password",
	TenantName:       "Some Project",
	DomainName:       "default",
}

var ArizonaCloudYAML = clientconfig.Cloud{
	RegionName:   "PHX",
	EndpointType: "public",
	AuthType:     clientconfig.AuthToken,
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://az.example.com:5000/v3",
		Token:       "12345",
		ProjectName: "Some Project",
		DomainName:  "default",
	},
	Verify: &iTrue,
}

var ArizonaClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://az.example.com:5000/v3",
		Token:       "12345",
		ProjectName: "Some Project",
		DomainName:  "default",
	},
}

var ArizonaEnvAuth = map[string]string{
	"OS_AUTH_URL":     "https://az.example.com:5000/v3",
	"OS_TOKEN":        "12345",
	"OS_PROJECT_NAME": "Some Project",
	"OS_DOMAIN_NAME":  "default",
}

var ArizonaAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectName: "Some Project",
		DomainName:  "default",
	},
	IdentityEndpoint: "https://az.example.com:5000/v3",
	TokenID:          "12345",
	TenantName:       "Some Project",
}

var NewMexicoCloudYAML = clientconfig.Cloud{
	RegionName:   "SAF",
	EndpointType: "admin",
	AuthType:     clientconfig.AuthPassword,
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:           "https://nm.example.com:5000/v3",
		Username:          "jdoe",
		Password:          "password",
		ProjectName:       "Some Project",
		ProjectDomainName: "Some Domain",
		UserDomainName:    "Some OtherDomain",
		DomainName:        "default",
	},
	Verify: &iTrue,
}

var NewMexicoClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:           "https://nm.example.com:5000/v3",
		Username:          "jdoe",
		Password:          "password",
		ProjectName:       "Some Project",
		ProjectDomainName: "Some Domain",
		UserDomainName:    "Other Domain",
		DomainName:        "default",
	},
}

var NewMexicoEnvAuth = map[string]string{
	"OS_AUTH_URL":            "https://nm.example.com:5000/v3",
	"OS_USERNAME":            "jdoe",
	"OS_PASSWORD":            "password",
	"OS_PROJECT_NAME":        "Some Project",
	"OS_PROJECT_DOMAIN_NAME": "Some Domain",
	"OS_USER_DOMAIN_NAME":    "Other Domain",
	"OS_DOMAIN_NAME":         "default",
}

var NewMexicoAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectName: "Some Project",
		DomainName:  "Some Domain",
	},
	IdentityEndpoint: "https://nm.example.com:5000/v3",
	Username:         "jdoe",
	Password:         "password",
	TenantName:       "Some Project",
	DomainName:       "Other Domain",
}

var NevadaCloudYAML = clientconfig.Cloud{
	RegionName:   "LAS",
	AuthType:     "password",
	EndpointType: "internal",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:           "https://nv.example.com:5000/v3",
		UserID:            "12345",
		Password:          "password",
		ProjectName:       "Some Project",
		ProjectDomainName: "Some Domain",
	},
	Verify: &iTrue,
}

var NevadaClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:           "https://nv.example.com:5000/v3",
		UserID:            "12345",
		Password:          "password",
		ProjectName:       "Some Project",
		ProjectDomainName: "Some Domain",
	},
}

var NevadaEnvAuth = map[string]string{
	"OS_AUTH_URL":            "https://nv.example.com:5000/v3",
	"OS_USER_ID":             "12345",
	"OS_PASSWORD":            "password",
	"OS_PROJECT_NAME":        "Some Project",
	"OS_PROJECT_DOMAIN_NAME": "Some Domain",
}

var NevadaAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectName: "Some Project",
		DomainName:  "Some Domain",
	},
	IdentityEndpoint: "https://nv.example.com:5000/v3",
	UserID:           "12345",
	Password:         "password",
	TenantName:       "Some Project",
}

var TexasCloudYAML = clientconfig.Cloud{
	RegionName: "AUS",
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:        "https://tx.example.com:5000/v3",
		Username:       "jdoe",
		Password:       "password",
		ProjectName:    "Some Project",
		UserDomainName: "Some Domain",
		DefaultDomain:  "default",
	},
	Verify: &iTrue,
}

var TexasClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:        "https://tx.example.com:5000/v3",
		Username:       "jdoe",
		Password:       "password",
		ProjectName:    "Some Project",
		UserDomainName: "Some Domain",
		DefaultDomain:  "default",
	},
}

var TexasEnvAuth = map[string]string{
	"OS_AUTH_URL":         "https://tx.example.com:5000/v3",
	"OS_USERNAME":         "jdoe",
	"OS_PASSWORD":         "password",
	"OS_PROJECT_NAME":     "Some Project",
	"OS_USER_DOMAIN_NAME": "Some Domain",
	"OS_DEFAULT_DOMAIN":   "default",
}

var TexasAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectName: "Some Project",
		DomainID:    "default",
	},
	IdentityEndpoint: "https://tx.example.com:5000/v3",
	Username:         "jdoe",
	Password:         "password",
	TenantName:       "Some Project",
	DomainName:       "Some Domain",
}

var CloudYAML = clientconfig.Clouds{
	Clouds: map[string]clientconfig.Cloud{
		"hawaii":     HawaiiCloudYAML,
		"florida":    FloridaCloudYAML,
		"california": CaliforniaCloudYAML,
		"arizona":    ArizonaCloudYAML,
		"newmexico":  NewMexicoCloudYAML,
		"nevada":     NevadaCloudYAML,
		"texas":      TexasCloudYAML,
	},
}

var AlbertaCloudYAML = clientconfig.Cloud{
	RegionName: "YYC",
	AuthType:   clientconfig.AuthPassword,
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://ab.example.com:5000/v2.0",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
	Verify: &iTrue,
}

var AlbertaClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://ab.example.com:5000/v2.0",
		Username:    "jdoe",
		Password:    "password",
		ProjectName: "Some Project",
	},
}

var AlbertaEnvAuth = map[string]string{
	"OS_AUTH_URL":             "https://ab.example.com:5000/v2.0",
	"OS_USERNAME":             "jdoe",
	"OS_PASSWORD":             "password",
	"OS_PROJECT_NAME":         "Some Project",
	"OS_IDENTITY_API_VERSION": "2.0",
}

var AlbertaAuthOpts = &gophercloud.AuthOptions{
	IdentityEndpoint: "https://ab.example.com:5000/v2.0",
	Username:         "jdoe",
	Password:         "password",
	TenantName:       "Some Project",
}

var YukonCloudYAML = clientconfig.Cloud{
	RegionName: "YXY",
	AuthType:   clientconfig.AuthV2Token,
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://yt.example.com:5000/v2.0",
		Token:       "12345",
		ProjectName: "Some Project",
	},
	Verify: &iTrue,
}

var YukonClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:     "https://yt.example.com:5000/v2.0",
		Token:       "12345",
		ProjectName: "Some Project",
	},
}

var YukonEnvAuth = map[string]string{
	"OS_AUTH_URL":             "https://yt.example.com:5000/v2.0",
	"OS_TOKEN":                "12345",
	"OS_PROJECT_NAME":         "Some Project",
	"OS_IDENTITY_API_VERSION": "2",
}

var YukonAuthOpts = &gophercloud.AuthOptions{
	IdentityEndpoint: "https://yt.example.com:5000/v2.0",
	TokenID:          "12345",
	TenantName:       "Some Project",
}

var LegacyCloudYAML = clientconfig.Clouds{
	Clouds: map[string]clientconfig.Cloud{
		"alberta": AlbertaCloudYAML,
		"yukon":   YukonCloudYAML,
	},
}

var OregonCloudYAML = clientconfig.Cloud{
	RegionName: "PDX",
	AuthType:   clientconfig.AuthV3OIDCClientCredentials,
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:             "https://or.example.com:5000/v3",
		ClientID:            "my-client-id",
		ClientSecret:        "my-client-secret",
		AccessTokenEndpoint: "https://idp.example.com/oauth2/token",
		IdentityProvider:    "myidp",
		Protocol:            "openid",
		ProjectName:         "Some Project",
		ProjectDomainName:   "default",
	},
	Verify: &iTrue,
}

var OregonClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:             "https://or.example.com:5000/v3",
		ClientID:            "my-client-id",
		ClientSecret:        "my-client-secret",
		AccessTokenEndpoint: "https://idp.example.com/oauth2/token",
		IdentityProvider:    "myidp",
		Protocol:            "openid",
		ProjectName:         "Some Project",
		ProjectDomainName:   "default",
	},
}

var OregonEnvAuth = map[string]string{
	"OS_AUTH_URL":              "https://or.example.com:5000/v3",
	"OS_CLIENT_ID":             "my-client-id",
	"OS_CLIENT_SECRET":         "my-client-secret",
	"OS_ACCESS_TOKEN_ENDPOINT": "https://idp.example.com/oauth2/token",
	"OS_IDENTITY_PROVIDER":     "myidp",
	"OS_PROTOCOL":              "openid",
	"OS_PROJECT_NAME":          "Some Project",
	"OS_PROJECT_DOMAIN_NAME":   "default",
}

// OregonAuthOpts is the expected standard auth configuration for Oregon.
var OregonAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectName: "Some Project",
		DomainName:  "default",
	},
	IdentityEndpoint: "https://or.example.com:5000/v3",
	TenantName:       "Some Project",
}

var WashingtonCloudYAML = clientconfig.Cloud{
	RegionName: "SEA",
	AuthType:   clientconfig.AuthV3OIDCClientCredentials,
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:             "https://wa.example.com:5000/v3",
		ClientID:            "wa-client-id",
		ClientSecret:        "wa-client-secret",
		AccessTokenEndpoint: "https://idp.example.com/oauth2/token",
		IdentityProvider:    "myidp",
		Protocol:            "openid",
		AccessTokenType:     "id_token",
		OpenIDScope:         "openid profile",
		ProjectID:           "67890",
	},
	Verify: &iTrue,
}

var WashingtonClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:             "https://wa.example.com:5000/v3",
		ClientID:            "wa-client-id",
		ClientSecret:        "wa-client-secret",
		AccessTokenEndpoint: "https://idp.example.com/oauth2/token",
		IdentityProvider:    "myidp",
		Protocol:            "openid",
		AccessTokenType:     "id_token",
		OpenIDScope:         "openid profile",
		ProjectID:           "67890",
	},
}

var WashingtonEnvAuth = map[string]string{
	"OS_AUTH_URL":              "https://wa.example.com:5000/v3",
	"OS_CLIENT_ID":             "wa-client-id",
	"OS_CLIENT_SECRET":         "wa-client-secret",
	"OS_ACCESS_TOKEN_ENDPOINT": "https://idp.example.com/oauth2/token",
	"OS_IDENTITY_PROVIDER":     "myidp",
	"OS_PROTOCOL":              "openid",
	"OS_ACCESS_TOKEN_TYPE":     "id_token",
	"OS_OPENID_SCOPE":          "openid profile",
	"OS_PROJECT_ID":            "67890",
}

var WashingtonAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectID: "67890",
	},
	IdentityEndpoint: "https://wa.example.com:5000/v3",
	TenantID:         "67890",
}

var MontanaCloudYAML = clientconfig.Cloud{
	RegionName: "BIL",
	AuthType:   clientconfig.AuthV3OIDCClientCredentials,
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:           "https://mt.example.com:5000/v3",
		ClientID:          "mt-client-id",
		ClientSecret:      "mt-client-secret",
		DiscoveryEndpoint: "https://idp.example.com/.well-known/openid-configuration",
		IdentityProvider:  "myidp",
		Protocol:          "openid",
		ProjectName:       "Some Project",
		ProjectDomainName: "default",
	},
	Verify: &iTrue,
}

var MontanaClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:           "https://mt.example.com:5000/v3",
		ClientID:          "mt-client-id",
		ClientSecret:      "mt-client-secret",
		DiscoveryEndpoint: "https://idp.example.com/.well-known/openid-configuration",
		IdentityProvider:  "myidp",
		Protocol:          "openid",
		ProjectName:       "Some Project",
		ProjectDomainName: "default",
	},
}

var MontanaEnvAuth = map[string]string{
	"OS_AUTH_URL":            "https://mt.example.com:5000/v3",
	"OS_CLIENT_ID":           "mt-client-id",
	"OS_CLIENT_SECRET":       "mt-client-secret",
	"OS_DISCOVERY_ENDPOINT":  "https://idp.example.com/.well-known/openid-configuration",
	"OS_IDENTITY_PROVIDER":   "myidp",
	"OS_PROTOCOL":            "openid",
	"OS_PROJECT_NAME":        "Some Project",
	"OS_PROJECT_DOMAIN_NAME": "default",
}

var MontanaAuthOpts = &gophercloud.AuthOptions{
	Scope: &gophercloud.AuthScope{
		ProjectName: "Some Project",
		DomainName:  "default",
	},
	IdentityEndpoint: "https://mt.example.com:5000/v3",
	TenantName:       "Some Project",
}

var ColoradoCloudYAML = clientconfig.Cloud{
	RegionName: "COS",
	AuthType:   clientconfig.AuthV3OAuth2MTLSClientCredential,
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:  "https://co.example.com:5000/v3",
		ClientID: "co-client-id",
	},
	Verify:         &iTrue,
	ClientCertFile: "/home/myhome/client-cert.crt",
	ClientKeyFile:  "/home/myhome/client-cert.key",
	CACertFile:     "/home/myhome/ca.crt",
}

var ColoradoClientOpts = &clientconfig.ClientOpts{
	AuthInfo: &clientconfig.AuthInfo{
		AuthURL:  "https://co.example.com:5000/v3",
		ClientID: "co-client-id",
	},
}

var ColoradoEnvAuth = map[string]string{
	"OS_AUTH_URL":  "https://co.example.com:5000/v3",
	"OS_CLIENT_ID": "co-client-id",
}

// ColoradoAuthOpts is the expected standard auth configuration for Colorado.
var ColoradoAuthOpts = &gophercloud.AuthOptions{
	Scope:            &gophercloud.AuthScope{},
	IdentityEndpoint: "https://co.example.com:5000/v3",
}
