package node

import (
	"reflect"

	"gopkg.in/urfave/cli.v1"
)

type LoggedCliFlags = map[string]interface{}

// LogCliFlags parses all cli defined flags and stores their default values
// or the ones passed by the user.
func LogCliFlags(ctx *cli.Context) LoggedCliFlags {
	activeFlags := LoggedCliFlags{}

	args := ctx.App.Flags
	for _, f := range args {
		fname := f.GetName()
		isLocalSet := ctx.IsSet(fname)

		switch reflect.TypeOf(f) {
		case reflect.TypeOf(cli.BoolFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Bool(fname)
			} else {
				activeFlags[fname] = ctx.GlobalBool(fname)
			}
		case reflect.TypeOf(cli.BoolTFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.BoolT(fname)
			} else {
				activeFlags[fname] = ctx.GlobalBoolT(fname)
			}
		case reflect.TypeOf(cli.DurationFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Duration(fname)
			} else {
				activeFlags[fname] = ctx.GlobalDuration(fname)
			}
		case reflect.TypeOf(cli.Float64Flag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Float64(fname)
			} else {
				activeFlags[fname] = ctx.GlobalFloat64(fname)
			}
		case reflect.TypeOf(cli.GenericFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Generic(fname)
			} else {
				activeFlags[fname] = ctx.GlobalGeneric(fname)
			}
		case reflect.TypeOf(cli.Int64Flag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Int64(fname)
			} else {
				activeFlags[fname] = ctx.GlobalInt64(fname)
			}
		case reflect.TypeOf(cli.Int64SliceFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Int64Slice(fname)
			} else {
				activeFlags[fname] = ctx.GlobalInt64Slice(fname)
			}
		case reflect.TypeOf(cli.IntFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Int(fname)
			} else {
				activeFlags[fname] = ctx.GlobalInt(fname)
			}
		case reflect.TypeOf(cli.IntSliceFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.IntSlice(fname)
			} else {
				activeFlags[fname] = ctx.GlobalIntSlice(fname)
			}
		case reflect.TypeOf(cli.StringFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.String(fname)
			} else {
				activeFlags[fname] = ctx.GlobalString(fname)
			}
		case reflect.TypeOf(cli.StringSliceFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.StringSlice(fname)
			} else {
				activeFlags[fname] = ctx.GlobalStringSlice(fname)
			}
		case reflect.TypeOf(cli.UintFlag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Uint(fname)
			} else {
				activeFlags[fname] = ctx.GlobalUint(fname)
			}
		case reflect.TypeOf(cli.Uint64Flag{}):
			if isLocalSet {
				activeFlags[fname] = ctx.Uint64(fname)
			} else {
				activeFlags[fname] = ctx.GlobalUint64(fname)
			}
		}
	}

	return activeFlags
}
