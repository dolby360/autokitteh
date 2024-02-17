package sdkruntimes

import (
	"context"
	"fmt"
	"strings"

	"go.autokitteh.dev/autokitteh/internal/kittehs"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

type NewRuntimeFunc = func() (sdkservices.Runtime, error)

type Runtime struct {
	Desc sdktypes.Runtime
	New  func() (sdkservices.Runtime, error)
}

type runtimes []*Runtime

var _ sdkservices.Runtimes = (runtimes)(nil)

func New(rts []*Runtime) (sdkservices.Runtimes, error) {
	exts := make(map[string]bool)

	for _, rt := range rts {
		for _, ext := range sdktypes.GetRuntimeFileExtensions(rt.Desc) {
			if exts[ext] {
				return nil, fmt.Errorf("duplicate extension found: %q", ext)
			}

			exts[ext] = true
		}
	}

	return runtimes(rts), nil
}

func (s runtimes) List(context.Context) ([]sdktypes.Runtime, error) {
	return kittehs.Transform(s, func(rt *Runtime) sdktypes.Runtime { return rt.Desc }), nil
}

func (s runtimes) get(f func(sdktypes.Runtime) bool) *Runtime {
	_, rt := kittehs.FindFirst(s, func(rt *Runtime) bool { return f(rt.Desc) })
	return rt
}

func (s runtimes) New(ctx context.Context, n sdktypes.Name) (sdkservices.Runtime, error) {
	if rt := s.get(func(rt sdktypes.Runtime) bool { return sdktypes.GetRuntimeName(rt).String() == n.String() }); rt != nil {
		return rt.New()
	}

	return nil, nil
}

func MatchRuntimeByPath(rts []sdktypes.Runtime, path string) sdktypes.Runtime {
	// find longest match (the most specific) between all registered extensions.

	var (
		lastExt string
		lastRT  sdktypes.Runtime
	)

	for _, rt := range rts {
		exts := sdktypes.GetRuntimeFileExtensions(rt)
		for _, ext := range exts {
			if strings.HasSuffix(path, "."+ext) {
				if len(lastExt) < len(ext) {
					lastExt, lastRT = ext, rt
				}
			}
		}
	}

	return lastRT
}
