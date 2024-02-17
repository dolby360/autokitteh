package sdkvalues

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"go.autokitteh.dev/autokitteh/sdk/sdkerrors"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

func Wrap(v any) (sdktypes.Value, error) { return DefaultValueWrapper.Wrap(v) }

// Wraps a native go type in a Value.
func (w ValueWrapper) Wrap(v any) (sdktypes.Value, error) {
	if v, ok := v.(sdktypes.Value); ok {
		return v, nil
	}

	if t, ok := v.(time.Time); ok {
		return sdktypes.NewTimeValue(t), nil
	}

	if d, ok := v.(time.Duration); ok {
		return sdktypes.NewDurationValue(d), nil
	}

	if msg, ok := v.(json.RawMessage); ok {
		var vv sdktypes.Value
		if err := json.Unmarshal(msg, &vv); err != nil {
			return nil, fmt.Errorf("unmarshal value error: %w", err)
		}
		return vv, nil
	}

	vv := reflect.ValueOf(v)
	switch vk := vv.Kind(); vk {
	case reflect.Func:
		// Never wrap functions here as we have no way to know its associated data.
		fallthrough

	case reflect.Invalid:
		return nil, sdkerrors.ErrInvalidArgument

	case reflect.Ptr:
		if !vv.IsNil() {
			return w.Wrap(vv.Elem().Interface())
		}

		return sdktypes.NewNothingValue(), nil

	case reflect.Struct:
		vt := reflect.TypeOf(v)

		if vt.Size() == 0 {
			return sdktypes.NewNothingValue(), nil
		}

		fs := make(map[string]sdktypes.Value)

		for _, vfs := range reflect.VisibleFields(vt) {
			if !vfs.IsExported() || vfs.Anonymous {
				continue
			}

			fv := vv.FieldByIndex(vfs.Index)

			wfv, err := w.Wrap(fv.Interface())
			if err != nil {
				return nil, fmt.Errorf("unable to convert struct field: %w", err)
			}

			n := w.fromStructCaser(vfs.Name)

			fs[n] = wfv
		}

		if w.WrapStructAsMap {
			return sdktypes.NewDictValueFromStringMap(fs), nil
		} else {
			n := vt.Name()
			if len(n) == 0 {
				n = "struct"
			}

			sym, err := sdktypes.ParseSymbol(n)
			if err != nil {
				return nil, fmt.Errorf("wrapping invalid name %q: %w", n, err)
			}

			return sdktypes.NewStructValue(sdktypes.NewSymbolValue(sym), fs), nil
		}

	case reflect.Array, reflect.Slice:
		if vv.Type().Elem().Kind() == reflect.Uint8 {
			return sdktypes.NewBytesValue(vv.Interface().([]byte)), nil
		}

		vs := make([]sdktypes.Value, vv.Len())
		for i := 0; i < vv.Len(); i++ {
			var err error
			if vs[i], err = w.Wrap(vv.Index(i).Interface()); err != nil {
				return nil, fmt.Errorf("%d: %w", i, err)
			}
		}

		return sdktypes.NewListValue(vs), nil

	case reflect.Map:
		vs := make([]*sdktypes.DictValueItem, 0, vv.Len())
		for i := vv.MapRange(); i.Next(); {
			k, v := i.Key(), i.Value()

			di := &sdktypes.DictValueItem{}

			var err error

			if di.K, err = w.Wrap(k.Interface()); err != nil {
				return nil, fmt.Errorf("key %v: %w", k, err)
			}

			if di.V, err = w.Wrap(v.Interface()); err != nil {
				return nil, fmt.Errorf("key %v: %w", k, err)
			}

			vs = append(vs, di)
		}

		return sdktypes.NewDictValue(vs), nil

	default:
		if num, ok := v.(json.Number); ok {
			if i64, err := num.Int64(); err == nil {
				return sdktypes.NewIntegerValue(i64), nil
			} else if f64, err := num.Float64(); err == nil {
				return sdktypes.NewFloatValue(f64), nil
			} else {
				return sdktypes.NewStringValue(string(num)), nil
			}
		}

		// force a float if the type is a float. see comment below
		// for conversion from float64.
		if f64, ok := v.(float64); ok {
			return sdktypes.NewFloatValue(f64), nil
		}

		// Checking using CanConvert in case of a custom type that
		// wraps a primitive type. Example: type I int. In this case,
		// I.(int) will fail.

		boolType := reflect.TypeOf(true)
		if vv.CanConvert(boolType) {
			return sdktypes.NewBooleanValue(vv.Convert(boolType).Interface().(bool)), nil
		}

		float64Type := reflect.TypeOf(float64(9.0))
		if vv.CanConvert(float64Type) {
			// float64 is convertible from int, so we need to decide if we want
			// to create an int or a float.

			f64 := vv.Convert(float64Type).Interface().(float64)
			i64 := int64(f64)
			if f64 == float64(i64) {
				return sdktypes.NewIntegerValue(i64), nil
			}

			return sdktypes.NewFloatValue(f64), nil
		}

		strType := reflect.TypeOf("meow")
		if vv.CanConvert(strType) {
			return sdktypes.NewStringValue(vv.Convert(strType).Interface().(string)), nil
		}

		return nil, fmt.Errorf("unhandled type: %q/%q", vk, reflect.TypeOf(v))
	}
}
