// Package schema provides the Chrome Debugging Protocol
// commands, types, and events for the Schema domain.
//
// Provides information about the protocol schema.
//
// Generated by the chromedp-gen command.
package schema

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
)

// GetDomainsParams returns supported domains.
type GetDomainsParams struct{}

// GetDomains returns supported domains.
func GetDomains() *GetDomainsParams {
	return &GetDomainsParams{}
}

// GetDomainsReturns return values.
type GetDomainsReturns struct {
	Domains []*Domain `json:"domains,omitempty"` // List of supported domains.
}

// Do executes Schema.getDomains against the provided context and
// target handler.
//
// returns:
//   domains - List of supported domains.
func (p *GetDomainsParams) Do(ctxt context.Context, h cdp.Handler) (domains []*Domain, err error) {
	// execute
	var res GetDomainsReturns
	err = h.Execute(ctxt, cdp.CommandSchemaGetDomains, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Domains, nil
}
