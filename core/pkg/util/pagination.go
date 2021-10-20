package util

import "go.mongodb.org/mongo-driver/mongo/options"

type PaginateOptions struct {
	CurrentPage int64
	PageSize    int64
}

func PaginateFind(findOpts *options.FindOptions, pgOpts PaginateOptions) *options.FindOptions {
	skipElements := pgOpts.PageSize * pgOpts.CurrentPage

	findOpts.Limit = &pgOpts.PageSize
	findOpts.Skip = &skipElements

	return findOpts
}
