package schemas

import (
	"github.com/HewlettPackard/hpegl-vmaas-terraform-resources/internal/resources/validations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HTTPProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "HTTP Profile configuration",
		MaxItems:    1,
		ExactlyOneOf: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		ConflictsWith: []string{
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBHttpProfile",
						"LBFastTcpProfile",
						"LBFastUdpProfile",
						"LBClientSslProfile",
						"LBServerSslProfile",
						"LBCookiePersistenceProfile",
						"LBGenericPersistenceProfile",
						"LBSourceIpPersistenceProfile",
					}, false),
					Required:     true,
					InputDefault: "LBHttpProfile",
					Description:  "Provide the Supported values for serviceTypes",
				},
				"http_idle_timeout": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          15,
					ValidateDiagFunc: validations.IntBetween(1, 5400),
					Description:      "http_idle_timeout for Network Load balancer Profile",
				},
				"request_header_size": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          1024,
					ValidateDiagFunc: validations.IntBetween(1, 65536),
					Description:      "request_header_size for Network Load balancer Profile",
				},
				"response_header_size": {
					Type:             schema.TypeInt,
					Optional:         true,
					ValidateDiagFunc: validations.IntBetween(1, 65536),
					Default:          4096,
					Description:      "response_header_size for Network Load balancer Profile",
				},
				"redirection": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "redirection for Network Load balancer Profile",
				},
				"x_forwarded_for": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"INSERT",
						"REPLACE",
					}, false),
					Required:     true,
					InputDefault: "INSERT",
					Description:  "x_forwarded_for for Network Load balancer Profile",
				},
				"request_body_size": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "request_body_size for Network Load balancer Profile",
				},
				"response_timeout": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          60,
					ValidateDiagFunc: validations.IntBetween(1, 2147483647),
					Description:      "response_timeout for Network Load balancer Profile",
				},
				"ntlm_authentication": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "ntlm_authentication for Network Load balancer Profile",
				},
			},
		},
	}
}

func TCPProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "TCP Profile configuration",
		MaxItems:    1,
		ExactlyOneOf: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		ConflictsWith: []string{
			"http_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBHttpProfile",
						"LBFastTcpProfile",
						"LBFastUdpProfile",
						"LBClientSslProfile",
						"LBServerSslProfile",
						"LBCookiePersistenceProfile",
						"LBGenericPersistenceProfile",
						"LBSourceIpPersistenceProfile",
					}, false),
					Required:     true,
					InputDefault: "LBFastTcpProfile",
					Description:  "Provide the  Supported values for serviceType",
				},
				"fast_tcp_idle_timeout": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          1800,
					ValidateDiagFunc: validations.IntBetween(1, 2147483647),
					Description:      "http_idle_timeout for Network Load balancer Profile",
				},
				"connection_close_timeout": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          8,
					ValidateDiagFunc: validations.IntBetween(1, 60),
					Description:      "connection_close_timeout for Network Load balancer Profile",
				},
				"ha_flow_mirroring": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "ha_flow_mirroring for Network Load balancer Profile",
				},
			},
		},
	}
}

func UDPProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "UDP Profile configuration",
		MaxItems:    1,
		ExactlyOneOf: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		ConflictsWith: []string{
			"http_profile",
			"tcp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBHttpProfile",
						"LBFastTcpProfile",
						"LBFastUdpProfile",
						"LBClientSslProfile",
						"LBServerSslProfile",
						"LBCookiePersistenceProfile",
						"LBGenericPersistenceProfile",
						"LBSourceIpPersistenceProfile",
					}, false),
					Required:     true,
					InputDefault: "LBFastUdpProfile",
					Description:  "Provide the  Supported values for serviceType",
				},
				"fast_udp_idle_timeout": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          300,
					ValidateDiagFunc: validations.IntBetween(1, 2147483647),
					Description:      "fast_udp_idle_timeout for Network Load balancer Profile",
				},
				"ha_flow_mirroring": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "ha_flow_mirroring for Network Load balancer Profile",
				},
			},
		},
	}
}

func CookieProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Cookie Profile configuration",
		MaxItems:    1,
		ExactlyOneOf: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		ConflictsWith: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBHttpProfile",
						"LBFastTcpProfile",
						"LBFastUdpProfile",
						"LBClientSslProfile",
						"LBServerSslProfile",
						"LBCookiePersistenceProfile",
						"LBGenericPersistenceProfile",
						"LBSourceIpPersistenceProfile",
					}, false),
					Required:     true,
					InputDefault: "LBCookiePersistenceProfile",
					Description:  "Provide the  Supported values for serviceType",
				},
				"cookie_name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "cookie_name for Network Load balancer Profile",
				},
				"cookie_fallback": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "cookie_fallback for Network Load balancer Profile",
				},
				"cookie_garbling": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "cookie_garbling for Network Load balancer Profile",
				},
				"cookie_mode": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"INSERT",
						"PREFIX",
						"REWRITE",
					}, false),
					Required:     true,
					InputDefault: "INSERT",
					Description:  "Provide the  Supported values for cookie_mode",
				},
				"cookie_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBPersistenceCookieTime",
						"LBSessionCookieTime",
					}, false),
					Required:     true,
					InputDefault: "LBSessionCookieTime",
					Description:  "Provide the  Supported values for cookie_type",
				},
				"cookie_domain": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "cookie_domain for Network Load balancer Profile",
				},
				"cookie_path": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "cookie_path for Network Load balancer Profile",
				},
				"max_idle_time": {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: "max_idle_time for Network Load balancer Profile",
				},
				"max_cookie_age": {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: "max_cookie_age for Network Load balancer Profile",
				},
				"share_persistence": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "ntlm_authentication for Network Load balancer Profile",
				},
			},
		},
	}
}

func SourceIPProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "SourceIP Profile configuration",
		MaxItems:    1,
		ExactlyOneOf: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		ConflictsWith: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBHttpProfile",
						"LBFastTcpProfile",
						"LBFastUdpProfile",
						"LBClientSslProfile",
						"LBServerSslProfile",
						"LBCookiePersistenceProfile",
						"LBGenericPersistenceProfile",
						"LBSourceIpPersistenceProfile",
					}, false),
					Required:     true,
					InputDefault: "LBSourceIpPersistenceProfile",
					Description:  "Provide the  Supported values for serviceType",
				},
				"share_persistence": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "ntlm_authentication for Network Load balancer Profile",
				},
				"ha_persistence_mirroring": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "ha_persistence_mirroring for Network Load balancer Profile",
				},
				"persistence_entry_timeout": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          300,
					ValidateDiagFunc: validations.IntBetween(1, 2147483647),
					Description:      "persistence_entry_timeout for Network Load balancer Profile",
				},
				"purge_entries_when_full": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "purge_entries_when_full for Network Load balancer Profile",
				},
			},
		},
	}
}

func GenericProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Generic Profile configuration",
		MaxItems:    1,
		ExactlyOneOf: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		ConflictsWith: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"client_profile",
			"server_profile",
		},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBHttpProfile",
						"LBFastTcpProfile",
						"LBFastUdpProfile",
						"LBClientSslProfile",
						"LBServerSslProfile",
						"LBCookiePersistenceProfile",
						"LBGenericPersistenceProfile",
						"LBSourceIpPersistenceProfile",
					}, false),
					Required:     true,
					InputDefault: "LBGenericPersistenceProfile",
					Description:  "Provide the  Supported values for serviceType",
				},
				"share_persistence": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "ntlm_authentication for Network Load balancer Profile",
				},
				"ha_persistence_mirroring": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					Description: "ha_persistence_mirroring for Network Load balancer Profile",
				},
				"persistence_entry_timeout": {
					Type:             schema.TypeInt,
					Optional:         true,
					Default:          300,
					ValidateDiagFunc: validations.IntBetween(1, 2147483647),
					Description:      "persistence_entry_timeout for Network Load balancer Profile",
				},
			},
		},
	}
}

func ClientProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Client Profile configuration",
		MaxItems:    1,
		ExactlyOneOf: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		ConflictsWith: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"server_profile",
		},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBHttpProfile",
						"LBFastTcpProfile",
						"LBFastUdpProfile",
						"LBClientSslProfile",
						"LBServerSslProfile",
						"LBCookiePersistenceProfile",
						"LBGenericPersistenceProfile",
						"LBSourceIpPersistenceProfile",
					}, false),
					Required:     true,
					InputDefault: "LBClientSslProfile",
					Description:  "Provide the  Supported values for serviceType",
				},
				"ssl_suite": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"BALANCED",
						"HIGH_SECURITY",
						"HIGH_COMPATIBILITY",
						"CUSTOM",
					}, false),
					Required:     true,
					InputDefault: "BALANCED",
					Description:  "Provide the  Supported values for ssl_suite",
				},
				"session_cache": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "session_cache for Network Load balancer Profile",
				},
				"session_cache_entry_timeout": {
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     300,
					Description: "session_cache_entry_timeout for Network Load balancer Profile",
				},
				"prefer_server_cipher": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "prefer_server_cipher for Network Load balancer Profile",
				},
			},
		},
	}
}

func ServerProfileSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Server Profile configuration",
		MaxItems:    1,
		ExactlyOneOf: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
			"server_profile",
		},
		ConflictsWith: []string{
			"http_profile",
			"tcp_profile",
			"udp_profile",
			"cookie_profile",
			"sourceip_profile",
			"generic_profile",
			"client_profile",
		},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"service_type": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"LBHttpProfile",
						"LBFastTcpProfile",
						"LBFastUdpProfile",
						"LBClientSslProfile",
						"LBServerSslProfile",
						"LBCookiePersistenceProfile",
						"LBGenericPersistenceProfile",
						"LBSourceIpPersistenceProfile",
					}, false),
					Required:     true,
					InputDefault: "LBServerSslProfile",
					Description:  "Provide the  Supported values for serviceType",
				},
				"ssl_suite": {
					Type: schema.TypeString,
					ValidateDiagFunc: validations.StringInSlice([]string{
						"BALANCED",
						"HIGH_SECURITY",
						"HIGH_COMPATIBILITY",
						"CUSTOM",
					}, false),
					Required:     true,
					InputDefault: "BALANCED",
					Description:  "Provide the  Supported values for ssl_suite",
				},
				"session_cache": {
					Type:        schema.TypeBool,
					Required:    true,
					Description: "session_cache for Network Load balancer Profile",
				},
			},
		},
	}
}
