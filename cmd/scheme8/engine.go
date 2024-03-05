package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/adzeitor/goscheme/scheme"
	"github.com/adzeitor/goscheme/sexpr"
)

type EngineContext struct {
	Editor bool
	Time   int

	// FIXME: move it to dependency, or depend sdl on EngineContext.
	// sdl specific
	Surface *sdl.Surface
	Keys    map[sdl.Scancode]bool
	Pressed []sdl.Keysym
	Font    *ttf.Font
}

func environmentWithSDL(ctx *EngineContext) scheme.Environment {
	env := scheme.DefaultEnvironment()
	scheme.AddFuncToEnv(env, "rectfill", rectFillBuiltin(ctx))
	scheme.AddFuncToEnv(env, "cls", clsBuiltin(ctx))
	scheme.AddFuncToEnv(env, "pset", psetBuiltin(ctx))
	scheme.AddFuncToEnv(env, "btn", btnBuiltin(ctx))
	scheme.AddFuncToEnv(env, "print", printBuiltin(ctx))
	return env
}

func rectFillBuiltin(ctx *EngineContext) scheme.Builtin {
	return func(args []sexpr.Expr, env scheme.Environment) sexpr.Expr {
		x := args[0].(int)
		y := args[1].(int)
		w := args[2].(int)
		h := args[3].(int)
		c := args[4].(int)
		rectFill(ctx, x, y, w, h, c)
		return nil
	}
}

func clsBuiltin(ctx *EngineContext) scheme.Builtin {
	return func(args []sexpr.Expr, env scheme.Environment) sexpr.Expr {
		cls(ctx, args[0].(int))
		return nil
	}
}

func psetBuiltin(ctx *EngineContext) scheme.Builtin {
	return func(args []sexpr.Expr, env scheme.Environment) sexpr.Expr {
		x := args[0].(int)
		y := args[1].(int)
		color := args[2].(int)
		pset(ctx, x, y, color)
		return nil
	}
}

func btnBuiltin(ctx *EngineContext) scheme.Builtin {
	return func(args []sexpr.Expr, env scheme.Environment) sexpr.Expr {
		return ctx.Keys[buttons[args[0].(int)]]
	}
}

func printBuiltin(ctx *EngineContext) scheme.Builtin {
	return func(args []sexpr.Expr, env scheme.Environment) sexpr.Expr {
		x := args[1].(int)
		y := args[2].(int)
		s := args[0].(string)
		printText(ctx, s, x, y, 7)
		return nil
	}
}
