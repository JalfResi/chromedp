// Package storage provides the Chrome Debugging Protocol
// commands, types, and events for the Storage domain.
//
// Generated by the chromedp-gen command.
package storage

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
)

// ClearDataForOriginParams clears storage for origin.
type ClearDataForOriginParams struct {
	Origin       string `json:"origin"`       // Security origin.
	StorageTypes string `json:"storageTypes"` // Comma separated origin names.
}

// ClearDataForOrigin clears storage for origin.
//
// parameters:
//   origin - Security origin.
//   storageTypes - Comma separated origin names.
func ClearDataForOrigin(origin string, storageTypes string) *ClearDataForOriginParams {
	return &ClearDataForOriginParams{
		Origin:       origin,
		StorageTypes: storageTypes,
	}
}

// Do executes Storage.clearDataForOrigin against the provided context and
// target handler.
func (p *ClearDataForOriginParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandStorageClearDataForOrigin, p, nil)
}

// GetUsageAndQuotaParams returns usage and quota in bytes.
type GetUsageAndQuotaParams struct {
	Origin string `json:"origin"` // Security origin.
}

// GetUsageAndQuota returns usage and quota in bytes.
//
// parameters:
//   origin - Security origin.
func GetUsageAndQuota(origin string) *GetUsageAndQuotaParams {
	return &GetUsageAndQuotaParams{
		Origin: origin,
	}
}

// GetUsageAndQuotaReturns return values.
type GetUsageAndQuotaReturns struct {
	Usage float64 `json:"usage,omitempty"` // Storage usage (bytes).
	Quota float64 `json:"quota,omitempty"` // Storage quota (bytes).
}

// Do executes Storage.getUsageAndQuota against the provided context and
// target handler.
//
// returns:
//   usage - Storage usage (bytes).
//   quota - Storage quota (bytes).
func (p *GetUsageAndQuotaParams) Do(ctxt context.Context, h cdp.Handler) (usage float64, quota float64, err error) {
	// execute
	var res GetUsageAndQuotaReturns
	err = h.Execute(ctxt, cdp.CommandStorageGetUsageAndQuota, p, &res)
	if err != nil {
		return 0, 0, err
	}

	return res.Usage, res.Quota, nil
}
