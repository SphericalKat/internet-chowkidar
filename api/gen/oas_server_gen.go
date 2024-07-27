// Code generated by ogen, DO NOT EDIT.

package genapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateAbuseReport implements createAbuseReport operation.
	//
	// Create a new abuse report.
	//
	// POST /abuse-reports
	CreateAbuseReport(ctx context.Context, req *AbuseReportInput) (*AbuseReport, error)
	// CreateBlock implements createBlock operation.
	//
	// Create a new block.
	//
	// POST /blocks
	CreateBlock(ctx context.Context, req *BlockInput) (*Block, error)
	// CreateCategory implements createCategory operation.
	//
	// Create a new category.
	//
	// POST /categories
	CreateCategory(ctx context.Context, req *CreateCategoryReq) (*Category, error)
	// CreateISP implements createISP operation.
	//
	// Create a new ISP.
	//
	// POST /isps
	CreateISP(ctx context.Context, req *ISPInput) (*ISP, error)
	// CreateSite implements createSite operation.
	//
	// Create a new site.
	//
	// POST /sites
	CreateSite(ctx context.Context, req *SiteInput) (*SiteCreate, error)
	// CreateSiteSuggestion implements createSiteSuggestion operation.
	//
	// Create a new site suggestion.
	//
	// POST /sites/suggestions
	CreateSiteSuggestion(ctx context.Context, req *SiteSuggestionInput) (*SiteSuggestion, error)
	// HealthCheck implements healthCheck operation.
	//
	// Health check.
	//
	// GET /health
	HealthCheck(ctx context.Context) (string, error)
	// ListAbuseReports implements listAbuseReports operation.
	//
	// List all abuse reports.
	//
	// GET /abuse-reports
	ListAbuseReports(ctx context.Context) ([]AbuseReport, error)
	// ListBlocks implements listBlocks operation.
	//
	// List all blocks.
	//
	// GET /blocks
	ListBlocks(ctx context.Context) ([]Block, error)
	// ListCategories implements listCategories operation.
	//
	// List all categories.
	//
	// GET /categories
	ListCategories(ctx context.Context) ([]Category, error)
	// ListISPs implements listISPs operation.
	//
	// List all ISPs.
	//
	// GET /isps
	ListISPs(ctx context.Context, params ListISPsParams) ([]ISP, error)
	// ListSiteSuggestions implements listSiteSuggestions operation.
	//
	// List all site suggestions.
	//
	// GET /sites/suggestions
	ListSiteSuggestions(ctx context.Context) ([]SiteSuggestion, error)
	// ListSites implements listSites operation.
	//
	// List all sites.
	//
	// GET /sites
	ListSites(ctx context.Context, params ListSitesParams) ([]Site, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
