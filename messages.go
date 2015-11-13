package cloudflare

import "encoding/json"

// Response - Cloudflare API Response.
type Response struct {
	Result     json.RawMessage `json:"result"`
	ResultInfo *ResultInfo     `json:"result_info"`

	Success  bool     `json:"success"`
	Errors   []string `json:"errors"`
	Messages []string `json:"messages"`
}

// ResultInfo - Cloudflare API Response Result Info.
type ResultInfo struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
	Count      int `json:"count,omitempty"`
	TotalCount int `json:"total_count,omitempty"`
}
