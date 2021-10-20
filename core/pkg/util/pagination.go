package util

import (
	"github.com/monitprod/core/pkg/vo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PaginateFind(findOpts *options.FindOptions, pgOpts vo.PaginateOptions) *options.FindOptions {
	skipElements := pgOpts.PageSize * pgOpts.CurrentPage

	findOpts.Limit = &pgOpts.PageSize
	findOpts.Skip = &skipElements

	return findOpts
}
