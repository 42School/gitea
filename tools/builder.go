package tools

import "xorm.io/builder"

//Union is a helper to build union with builder
func Union(args *builder.Builder, unionTp string, unionCond *builder.Builder) *builder.Builder {
	if args == nil {
		return unionCond
	}
	return args.Union(unionTp, unionCond)
}
