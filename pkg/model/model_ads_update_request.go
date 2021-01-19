/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdsUpdateRequest struct {
	AdId                    int64    `json:"ad_id,omitempty"`
	AdName                  string   `json:"ad_name,omitempty"`
	ConfiguredStatus        AdStatus `json:"configured_status,omitempty"`
	ImpressionTrackingUrl   *string  `json:"impression_tracking_url,omitempty"`
	ClickTrackingUrl        *string  `json:"click_tracking_url,omitempty"`
	FeedsInteractionEnabled *bool    `json:"feeds_interaction_enabled,omitempty"`
	AccountId               int64    `json:"account_id,omitempty"`
}
