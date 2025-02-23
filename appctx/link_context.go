package appctx

import (
	"context"
)

var KeyCurrentPath struct{}

func CurrentPath(ctx context.Context) string {
	if path, ok := ctx.Value(KeyCurrentPath).(string); ok {
		return path
	}

	return ""
}

func IsCurrentPath(ctx context.Context, path string) bool {
	return CurrentPath(ctx) == path
}
