package sdkmodule

import (
	"context"
	"fmt"

	"go.autokitteh.dev/autokitteh/internal/kittehs"
	"go.autokitteh.dev/autokitteh/sdk/sdkerrors"
	"go.autokitteh.dev/autokitteh/sdk/sdkexecutor"
	"go.autokitteh.dev/autokitteh/sdk/sdklogger"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

type Module interface {
	Describe() sdktypes.Module

	Configure(ctx context.Context, xid sdktypes.ExecutorID, config string) (map[string]sdktypes.Value, error)

	sdkexecutor.Caller
}

type module struct {
	desc sdktypes.Module
	opts moduleOpts
}

func New(optfns ...Optfn) Module {
	opts := moduleOpts{
		funcs: make(map[string]*funcOpts),
		vars:  make(map[string]*varOpts),
	}

	for i, optfn := range optfns {
		if err := optfn(&opts); err != nil {
			sdklogger.Panic("option error", "i", i, "err", err)
		}
	}

	if opts.dataFromConfig == nil {
		opts.dataFromConfig = func(string) ([]byte, error) { return nil, nil }
	}

	return &module{
		desc: sdktypes.NewModule(
			kittehs.TransformMapValues(opts.funcs, func(f *funcOpts) sdktypes.ModuleFunction {
				return kittehs.Must1(sdktypes.ModuleFunctionFromProto(&f.desc))
			}),
			kittehs.TransformMapValues(opts.vars, func(v *varOpts) sdktypes.ModuleVariable {
				return kittehs.Must1(sdktypes.ModuleVariableFromProto(&v.desc))
			}),
		),
		opts: opts,
	}
}

func (m *module) Configure(ctx context.Context, xid sdktypes.ExecutorID, config string) (map[string]sdktypes.Value, error) {
	values := make(map[string]sdktypes.Value)

	data, err := m.opts.dataFromConfig(config)
	if err != nil {
		return nil, fmt.Errorf("config: %w", err)
	}

	for k, v := range m.opts.vars {
		if values[k] != nil {
			return nil, fmt.Errorf("value name %q is already set: %w", k, sdkerrors.ErrConflict)
		}

		if values[k], err = v.fn(xid, data); err != nil {
			return nil, fmt.Errorf("%q: %w", k, err)
		}
	}

	for k, f := range m.opts.funcs {
		values[k] = sdktypes.NewFunctionValue(xid, k, data, f.flags, kittehs.Must1(sdktypes.ModuleFunctionFromProto(&f.desc)))
	}

	return values, nil
}

func (m *module) Call(ctx context.Context, fnv sdktypes.Value, args []sdktypes.Value, kwargs map[string]sdktypes.Value) (sdktypes.Value, error) {
	name := sdktypes.GetFunctionValueName(fnv)
	if name == nil {
		return nil, sdkerrors.ErrInvalidArgument
	}

	fn, ok := m.opts.funcs[name.String()]
	if !ok {
		return nil, sdkerrors.ErrNotFound
	}

	return fn.fn(wrapCallContext(ctx, m, fnv), args, kwargs)
}

func (m *module) Describe() sdktypes.Module { return m.desc }
