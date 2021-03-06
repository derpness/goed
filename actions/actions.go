/*
Set of actions that can be dispatched.
Actions are dispatched, and processed one at a time by the action bus for
concurency safety.
*/
package actions

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/tcolar/goed/core"
)

// Action runner, all actions are defined on this type
// AFAIK need to be on a type for easy reflection (go/importer may help ??)
var Ar *ar = &ar{}

func Exec(action string, args []string) (res []string, err error) {
	proto, ok := actions[action]
	if !ok {
		return res, fmt.Errorf("No such action %s", action)
	}
	if len(proto.ins) != len(args) {
		return res, fmt.Errorf("Incorrect number of arguments for %s, got %d, want %d",
			action, len(args), len(proto.ins))
	}
	in := []reflect.Value{reflect.ValueOf(Ar)}
	for i := 0; i < len(proto.ins); {
		argType := proto.ins[i]
		end := i + 1
		if strings.HasPrefix(argType.String(), "[]") {
			end = len(proto.ins)
		}
		val, err := argToVal(args[i:end], argType)
		if err != nil {
			return res, err
		}
		in = append(in, val)
		i = end
	}
	out := proto.f.Call(in)
	for i, argType := range proto.outs {
		strs, err := valToStrs(out[i], argType)
		if err != nil {
			return res, err
		}
		res = append(res, strs...)
	}
	return res, nil
}

type ar struct{}

var actions map[string]actionProto = map[string]actionProto{}

type actionProto struct {
	f    reflect.Value
	ins  []reflect.Type
	outs []reflect.Type
	sig  string
}

func RegisterActions() {
	Ar = &ar{}
	t := reflect.TypeOf(Ar)
	for i := 0; i != t.NumMethod(); i++ {
		registerAction(t.Method(i))
	}
}

func registerAction(m reflect.Method) {
	proto := actionProto{
		f: m.Func,
	}
	ins := []string{}
	outs := []string{}
	for i := 1; i < m.Type.NumIn(); i++ {
		proto.ins = append(proto.ins, m.Type.In(i))
		ins = append(ins, m.Type.In(i).String())
	}
	for i := 0; i < m.Type.NumOut(); i++ {
		proto.outs = append(proto.outs, m.Type.Out(i))
		outs = append(outs, m.Type.Out(i).String())
	}
	proto.sig = fmt.Sprintf("(%s) %s", strings.Join(ins, ", "), strings.Join(outs, ", "))
	actions[toCamel(m.Name)] = proto
}

// dispath an action to the evnt bus
func d(action core.Action) {
	core.Bus.Dispatch(action)
}

func viewExists(vid int64) bool {
	v := core.Ed.ViewById(vid)
	if v == nil {
		return false
	}
	return true
}

func Usage() string {
	u := ""
	var keys []string
	for k := range actions {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		proto := actions[k]
		u += fmt.Sprintf("%s%s\n", k, proto.sig)
	}
	return u
}

func toCamel(s string) string {
	ns := []rune{}
	for i, c := range s {
		if unicode.IsUpper(c) {
			c = unicode.ToLower(c)
			if i > 0 && i < len(s)-1 {
				ns = append(ns, '_')
			}
		}
		ns = append(ns, c)
	}
	return string(ns)
}

func argToVal(args []string, toType reflect.Type) (v reflect.Value, err error) {
	t := toType.String()
	switch t {
	case "string":
		return reflect.ValueOf(args[0]), nil
	case "core.CursorMvmt":
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return v, err
		}
		return reflect.ValueOf(core.CursorMvmt(i)), nil
	case "int":
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return v, err
		}
		return reflect.ValueOf(i), nil
	case "int64":
		i, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return v, err
		}
		return reflect.ValueOf(i), nil
	case "bool":
		if strings.ToLower(args[0]) == "false" || args[0] == "0" {
			return reflect.ValueOf(false), nil
		}
		return reflect.ValueOf(true), nil
	case "[]string":
		arr := []string{}
		for _, arg := range args {
			arr = append(arr, arg)
		}
		return reflect.ValueOf(arr), nil
	default:
		return v, fmt.Errorf("Unhandled type : %s !", t)
	}
}

func valToStrs(v reflect.Value, fromType reflect.Type) (s []string, err error) {
	t := fromType.String()
	switch t {
	case "string":
		return []string{v.Interface().(string)}, nil
	case "int64":
		return []string{fmt.Sprintf("%d", v.Interface().(int64))}, nil
	case "int":
		return []string{fmt.Sprintf("%d", v.Interface().(int))}, nil
	case "bool":
		return []string{fmt.Sprintf("%t", v.Interface().(bool))}, nil
	case "error":
		return []string{fmt.Sprintf("%s", v.Interface().(error).Error())}, nil
	case "[]int64":
		array := []string{}
		for _, i := range v.Interface().([]int64) {
			array = append(array, fmt.Sprintf("%d", i))
		}
		return array, nil
	case "[]core.Selection":
		array := []string{}
		for _, s := range v.Interface().([]core.Selection) {
			array = append(array, s.String())
		}
		return array, nil
	case "[]string":
		array := []string{}
		for _, s := range v.Interface().([]string) {
			array = append(array, s)
		}
		return array, nil
	default:
		return s, fmt.Errorf("Unhandled type : %s !", t)
	}
}
