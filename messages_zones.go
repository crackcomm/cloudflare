package cloudflare

import "time"

// Zone - Cloudflare Zone.
type Zone struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Status          string `json:"status,omitempty"`
	Paused          bool   `json:"paused,omitempty"`
	Type            string `json:"type,omitempty"`
	DevelopmentMode int    `json:"development_mode,omitempty"`

	NameServers         []string `json:"name_servers,omitempty"`
	OriginalNameServers []string `json:"original_name_servers,omitempty"`

	ModifiedOn time.Time `json:"modified_on,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	CheckedOn  time.Time `json:"checked_on,omitempty"`

	Meta  *ZoneMeta  `json:"meta,omitempty"`
	Owner *ZoneOwner `json:"owner,omitempty"`
	Plan  *ZonePlan  `json:"plan,omitempty"`

	Permissions []string `json:"permissions,omitempty"`
}

// ZoneOwner -
type ZoneOwner struct {
	Type  string `json:"type,omitempty"`
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

// ZoneMeta -
type ZoneMeta struct {
	Step                    int    `json:"step,omitempty"`
	PageRuleQuota           string `json:"page_rule_quota,omitempty"`
	CustomCertificateQuota  string `json:"custom_certificate_quota,omitempty"`
	WildcardProxiable       bool   `json:"wildcard_proxiable,omitempty"`
	PhishingDetected        bool   `json:"phishing_detected,omitempty"`
	MultipleRailgunsAllowed bool   `json:"multiple_railguns_allowed,omitempty"`
}

// ZonePlan -
type ZonePlan struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Price             int    `json:"price,omitempty"`
	Currency          string `json:"currency,omitempty"`
	Frequency         string `json:"frequency,omitempty"`
	LegacyID          string `json:"legacy_id,omitempty"`
	IsSubscribed      bool   `json:"is_subscribed,omitempty"`
	CanSubscribe      bool   `json:"can_subscribe,omitempty"`
	ExternallyManaged bool   `json:"externally_managed,omitempty"`
}

// ZonePatch -
type ZonePatch struct {
	Plan              *ZonePlan `json:"plan,omitempty"`
	Paused            bool      `json:"paused,omitempty"`
	VanityNameServers []string  `json:"vanity_name_servers,omitempty"`
}
