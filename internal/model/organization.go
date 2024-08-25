package model

import (
	"fmt"
	"github.com/dembygenesis/local.tools/internal/sysconsts"
	"github.com/dembygenesis/local.tools/internal/utilities/validationutils"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"strings"
	"time"
)

type UpdateOrganization struct {
	Id     int         `json:"id" validate:"required,greater_than_zero"`
	UserId null.Int    `json:"userId"`
	Name   null.String `json:"name"`
}

func (c *UpdateOrganization) Validate() error {
	if err := validationutils.Validate(c); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	hasAtLeastOneUpdateParameters := false

	if c.Name.Valid {
		if c.Name.Valid && strings.TrimSpace(c.Name.String) != "" {
			hasAtLeastOneUpdateParameters = true
		}
	}

	if !hasAtLeastOneUpdateParameters {
		return errors.New(sysconsts.ErrHasNotASingleValidateUpdateParameter)
	}

	return nil
}

type CreateOrganization struct {
	Name   string `json:"name" validate:"required"`
	UserId int    `json:"user_id"`
}

type Organization struct {
	Id            int       `json:"id" boil:"id"`
	Name          string    `json:"name" boil:"name"`
	CreatedBy     string    `json:"created_by" boil:"created_by"`
	LastUpdatedBy string    `json:"last_updated_by" boil:"last_updated_by"`
	CreatedAt     time.Time `json:"created_at" boil:"created_at"`
	LastUpdatedAt null.Time `json:"last_updated_at" boil:"last_updated_at"`
	IsActive      bool      `json:"is_active" boil:"is_active"`
}

type PaginatedOrganizations struct {
	Organizations []Organization `json:"organizations"`
	Pagination    *Pagination    `json:"pagination"`
}

type OrganizationFilters struct {
	OrganizationNameIn     []string  `query:"organization_name_in" json:"organization_name_in"`
	OrganizationIsActive   null.Bool `query:"is_active" json:"is_active"`
	IdsIn                  []int     `query:"ids_in" json:"ids_in"`
	CreatedBy              null.Int  `query:"created_by" json:"created_by"`
	LastUpdatedBy          null.Int  `query:"last_updated_by" json:"last_updated_by"`
	PaginationQueryFilters `swaggerignore:"true"`
}

type DeleteOrganization struct {
	ID int `json:"id" validate:"required,greater_than_zero"`
}

type RestoreOrganization struct {
	ID int `json:"id" validate:"required,greater_than_zero"`
}

func (c *CreateOrganization) Validate() error {
	if err := validationutils.Validate(c); err != nil {
		return fmt.Errorf("validate: %v", err)
	}
	return nil
}

func (c *CreateOrganization) ToOrganization() *Organization {
	organization := &Organization{
		Name:      c.Name,
		CreatedBy: fmt.Sprint(c.UserId),
	}
	return organization
}

func (c *OrganizationFilters) Validate() error {
	if err := c.ValidatePagination(); err != nil {
		return fmt.Errorf("pagination: %v", err)
	}
	if err := validationutils.Validate(c); err != nil {
		return fmt.Errorf("organization filters: %v", err)
	}

	return nil
}
