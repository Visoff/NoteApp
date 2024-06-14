package signals

import "syscall/js"

type Signal struct {
    Data chan interface{}
}

func NewSignal() Signal {
    return Signal{Data: make(chan interface{})}
}

func (s Signal) AttachToJS(parent js.Value, attribute_name string) {
    signal := js.Global().Get("Array").New()
    signal.Call("push", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        go func() {
            for {
                if len(args) < 1 || args[0].Type() != js.TypeFunction {
                    break
                }
                args[0].Invoke(js.ValueOf(<-s.Data))
            } 
        }()
        return nil
    }))
    signal.Call("push", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        s.Data <- args[0]
        return nil
    }))
    parent.Set(attribute_name, signal)
}
