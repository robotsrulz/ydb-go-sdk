// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// scriptingComposeOptions is a holder of options
type scriptingComposeOptions struct {
	panicCallback func(e interface{})
}

// ScriptingOption specified Scripting compose option
type ScriptingComposeOption func(o *scriptingComposeOptions)

// WithScriptingPanicCallback specified behavior on panic
func WithScriptingPanicCallback(cb func(e interface{})) ScriptingComposeOption {
	return func(o *scriptingComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Scripting which has functional fields composed both from t and x.
func (t Scripting) Compose(x Scripting, opts ...ScriptingComposeOption) (ret Scripting) {
	options := scriptingComposeOptions{}
	for _, opt := range opts {
		opt(&options)
	}
	{
		h1 := t.OnExecute
		h2 := x.OnExecute
		ret.OnExecute = func(s ScriptingExecuteStartInfo) func(ScriptingExecuteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(ScriptingExecuteDoneInfo) 
			if h1 != nil {
				r = h1(s)
			}
			if h2 != nil {
				r1 = h2(s)
			}
			return func(s ScriptingExecuteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(s)
				}
				if r1 != nil {
					r1(s)
				}
			}
		}
	}
	{
		h1 := t.OnStreamExecute
		h2 := x.OnStreamExecute
		ret.OnStreamExecute = func(s ScriptingStreamExecuteStartInfo) func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) 
			if h1 != nil {
				r = h1(s)
			}
			if h2 != nil {
				r1 = h2(s)
			}
			return func(s ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(ScriptingStreamExecuteDoneInfo) 
				if r != nil {
					r2 = r(s)
				}
				if r1 != nil {
					r3 = r1(s)
				}
				return func(s ScriptingStreamExecuteDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(s)
					}
					if r3 != nil {
						r3(s)
					}
				}
			}
		}
	}
	{
		h1 := t.OnExplain
		h2 := x.OnExplain
		ret.OnExplain = func(s ScriptingExplainStartInfo) func(ScriptingExplainDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(ScriptingExplainDoneInfo) 
			if h1 != nil {
				r = h1(s)
			}
			if h2 != nil {
				r1 = h2(s)
			}
			return func(s ScriptingExplainDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(s)
				}
				if r1 != nil {
					r1(s)
				}
			}
		}
	}
	{
		h1 := t.OnClose
		h2 := x.OnClose
		ret.OnClose = func(s ScriptingCloseStartInfo) func(ScriptingCloseDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(ScriptingCloseDoneInfo) 
			if h1 != nil {
				r = h1(s)
			}
			if h2 != nil {
				r1 = h2(s)
			}
			return func(s ScriptingCloseDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				if r != nil {
					r(s)
				}
				if r1 != nil {
					r1(s)
				}
			}
		}
	}
	return ret
}
func (t Scripting) onExecute(s ScriptingExecuteStartInfo) func(ScriptingExecuteDoneInfo) {
	fn := t.OnExecute
	if fn == nil {
		return func(ScriptingExecuteDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(ScriptingExecuteDoneInfo) {
			return
		}
	}
	return res
}
func (t Scripting) onStreamExecute(s ScriptingStreamExecuteStartInfo) func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
	fn := t.OnStreamExecute
	if fn == nil {
		return func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
			return func(ScriptingStreamExecuteDoneInfo) {
				return
			}
		}
	}
	res := fn(s)
	if res == nil {
		return func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
			return func(ScriptingStreamExecuteDoneInfo) {
				return
			}
		}
	}
	return func(s ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
		res := res(s)
		if res == nil {
			return func(ScriptingStreamExecuteDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t Scripting) onExplain(s ScriptingExplainStartInfo) func(ScriptingExplainDoneInfo) {
	fn := t.OnExplain
	if fn == nil {
		return func(ScriptingExplainDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(ScriptingExplainDoneInfo) {
			return
		}
	}
	return res
}
func (t Scripting) onClose(s ScriptingCloseStartInfo) func(ScriptingCloseDoneInfo) {
	fn := t.OnClose
	if fn == nil {
		return func(ScriptingCloseDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(ScriptingCloseDoneInfo) {
			return
		}
	}
	return res
}
func ScriptingOnExecute(t Scripting, c *context.Context, query string, parameters scriptingQueryParameters) func(result scriptingResult, _ error) {
	var p ScriptingExecuteStartInfo
	p.Context = c
	p.Query = query
	p.Parameters = parameters
	res := t.onExecute(p)
	return func(result scriptingResult, e error) {
		var p ScriptingExecuteDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
func ScriptingOnStreamExecute(t Scripting, c *context.Context, query string, parameters scriptingQueryParameters) func(error) func(error) {
	var p ScriptingStreamExecuteStartInfo
	p.Context = c
	p.Query = query
	p.Parameters = parameters
	res := t.onStreamExecute(p)
	return func(e error) func(error) {
		var p ScriptingStreamExecuteIntermediateInfo
		p.Error = e
		res := res(p)
		return func(e error) {
			var p ScriptingStreamExecuteDoneInfo
			p.Error = e
			res(p)
		}
	}
}
func ScriptingOnExplain(t Scripting, c *context.Context, query string) func(plan string, _ error) {
	var p ScriptingExplainStartInfo
	p.Context = c
	p.Query = query
	res := t.onExplain(p)
	return func(plan string, e error) {
		var p ScriptingExplainDoneInfo
		p.Plan = plan
		p.Error = e
		res(p)
	}
}
func ScriptingOnClose(t Scripting, c *context.Context) func(error) {
	var p ScriptingCloseStartInfo
	p.Context = c
	res := t.onClose(p)
	return func(e error) {
		var p ScriptingCloseDoneInfo
		p.Error = e
		res(p)
	}
}
