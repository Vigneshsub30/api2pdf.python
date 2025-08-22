package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ChromeAdvancedOptions represents the ChromeAdvancedOptions schema from the OpenAPI specification
type ChromeAdvancedOptions struct {
	Landscape string `json:"landscape,omitempty"`
	Printbackground bool `json:"printBackground,omitempty"`
}

// WkHtmlToPdfAdvancedOptions represents the WkHtmlToPdfAdvancedOptions schema from the OpenAPI specification
type WkHtmlToPdfAdvancedOptions struct {
	Orientation string `json:"orientation,omitempty"`
	Pagesize string `json:"pageSize,omitempty"`
}

// WkHtmlToPdfUrlToPdfRequest represents the WkHtmlToPdfUrlToPdfRequest schema from the OpenAPI specification
type WkHtmlToPdfUrlToPdfRequest struct {
	Inlinepdf bool `json:"inlinePdf,omitempty"`
	Options WkHtmlToPdfAdvancedOptions `json:"options,omitempty"`
	Url string `json:"url"`
	Filename string `json:"fileName,omitempty"`
}

// ApiResponseFailure represents the ApiResponseFailure schema from the OpenAPI specification
type ApiResponseFailure struct {
	Reason string `json:"reason,omitempty"` // The reason for the PDF generation failure
	Success bool `json:"success,omitempty"` // Will be false if the operation failed
}

// ChromeHtmlToPdfRequest represents the ChromeHtmlToPdfRequest schema from the OpenAPI specification
type ChromeHtmlToPdfRequest struct {
	Filename string `json:"fileName,omitempty"`
	Html string `json:"html"`
	Inlinepdf bool `json:"inlinePdf,omitempty"`
	Options ChromeAdvancedOptions `json:"options,omitempty"`
}

// LibreOfficeConvertRequest represents the LibreOfficeConvertRequest schema from the OpenAPI specification
type LibreOfficeConvertRequest struct {
	Filename string `json:"fileName,omitempty"`
	Inlinepdf bool `json:"inlinePdf,omitempty"`
	Url string `json:"url"`
}

// WkHtmlToPdfHtmlToPdfRequest represents the WkHtmlToPdfHtmlToPdfRequest schema from the OpenAPI specification
type WkHtmlToPdfHtmlToPdfRequest struct {
	Options WkHtmlToPdfAdvancedOptions `json:"options,omitempty"`
	Filename string `json:"fileName,omitempty"`
	Html string `json:"html"`
	Inlinepdf bool `json:"inlinePdf,omitempty"`
}

// ChromeUrlToPdfRequest represents the ChromeUrlToPdfRequest schema from the OpenAPI specification
type ChromeUrlToPdfRequest struct {
	Inlinepdf bool `json:"inlinePdf,omitempty"`
	Options ChromeAdvancedOptions `json:"options,omitempty"`
	Url string `json:"url"`
	Filename string `json:"fileName,omitempty"`
}

// MergeRequest represents the MergeRequest schema from the OpenAPI specification
type MergeRequest struct {
	Inlinepdf bool `json:"inlinePdf,omitempty"`
	Urls []string `json:"urls"`
	Filename string `json:"fileName,omitempty"`
}

// ApiResponseSuccess represents the ApiResponseSuccess schema from the OpenAPI specification
type ApiResponseSuccess struct {
	Mbout float64 `json:"mbOut,omitempty"` // The amount of megabytes of bandwidth generated from the resulting pdf
	Pdf string `json:"pdf,omitempty"` // A url to the PDF that will exist only for 24 hours
	Success bool `json:"success,omitempty"` // Will be true if the operation suceeded
	Cost float64 `json:"cost,omitempty"` // Cost of the operation (mbIn + mbOut) * $.001
	Mbin float64 `json:"mbIn,omitempty"` // The amount of megabytes of bandwidth used to process the pdf
}
