// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type Upload struct {
	Id string `json:"id,omitempty"`
	// Max time in seconds for the signed upload URL to be valid. If a successful upload has not occurred before the timeout limit, the direct upload is marked `timed_out`
	Timeout int32 `json:"timeout,omitempty"`
	Status string `json:"status,omitempty"`
	NewAssetSettings Asset `json:"new_asset_settings,omitempty"`
	// Only set once the upload is in the `asset_created` state.
	AssetId string `json:"asset_id,omitempty"`
	Error UploadError `json:"error,omitempty"`
	// If the upload URL will be used in a browser, you must specify the origin in order for the signed URL to have the correct CORS headers.
	CorsOrigin string `json:"cors_origin,omitempty"`
}
