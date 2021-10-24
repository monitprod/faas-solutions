package util

import (
	"github.com/monitprod/core/pkg/vo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PaginateFind starts at index 1
func PaginateFind(findOpts *options.FindOptions, pgOpts vo.PaginateOptions) *options.FindOptions {
	skipElements := pgOpts.PageSize * (pgOpts.CurrentPage - 1)

	findOpts.Limit = &pgOpts.PageSize
	findOpts.Skip = &skipElements

	return findOpts
}
