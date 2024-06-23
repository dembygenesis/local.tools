package api

import (
	"context"
	"github.com/dembygenesis/local.tools/internal/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . categoryService
type categoryService interface {
	ListCategories(ctx context.Context, filters *model.CategoryFilters) (*model.PaginatedCategories, error)
	CreateCategory(ctx context.Context, category *model.CreateCategory) (*model.Category, error)
	UpdateCategory(ctx context.Context, category *model.UpdateCategory) (*model.Category, error)
	DeleteCategory(ctx context.Context, params *model.DeleteCategory) error
	RestoreCategory(ctx context.Context, params *model.RestoreCategory) error
	GetCategoryByID(ctx context.Context, id int) (*model.Category, error)
}

//counterfeiter:generate . organizationService
type organizationService interface {
	ListOrganizations(ctx context.Context, filters *model.OrganizationFilters) (*model.PaginatedOrganizations, error)
	CreateOrganization(ctx context.Context, organization *model.CreateOrganization) (*model.Organization, error)
	UpdateOrganization(ctx context.Context, category *model.UpdateOrganization) (*model.Organization, error)
	DeleteOrganization(ctx context.Context, params *model.DeleteOrganization) error
	RestoreOrganization(ctx context.Context, params *model.RestoreOrganization) error
	GetOrganizationByID(ctx context.Context, id int) (*model.Organization, error)
}

//counterfeiter:generate . capturePagesService
type capturePagesService interface {
	ListCapturePages(ctx context.Context, filters *model.CapturePagesFilters) (*model.PaginatedCapturePages, error)
	CreateCapturePages(ctx context.Context, capturepage *model.CreateCapturePage) (*model.CapturePages, error)
	UpdateCapturePages(ctx context.Context, capturepage *model.UpdateCapturePages) (*model.CapturePages, error)
	DeleteCapturePages(ctx context.Context, params *model.DeleteCapturePages) error
	RestoreCapturePages(ctx context.Context, params *model.RestoreCapturePages) error
	GetCapturePageByID(ctx context.Context, id int) (*model.CapturePages, error)
}

//counterfeiter:generate . clickTrackerService
type clickTrackerService interface {
	ListClickTrackers(ctx context.Context, filters *model.ClickTrackerFilters) (*model.PaginatedClickTrackers, error)
	CreateClickTracker(ctx context.Context, params *model.CreateClickTracker) (*model.ClickTracker, error)
	UpdateClickTracker(ctx context.Context, clicktracker *model.UpdateClickTracker) (*model.ClickTracker, error)
	DeleteClickTracker(ctx context.Context, params *model.DeleteClickTracker) error
	RestoreClickTracker(ctx context.Context, params *model.RestoreClickTracker) error
	GetClickTrackerByID(ctx context.Context, id int) (*model.ClickTracker, error)
}

//counterfeiter:generate . userService
type userService interface {
	ListUsers(ctx context.Context, filters *model.UserFilters) (*model.PaginatedUsers, error)
}
