package main

import (
	"strings"
	"syscall/js"

	"github.com/Visoff/NoteApp/wasm_lib/editor"
	"github.com/Visoff/NoteApp/wasm_lib/signals"
)

func main() {
    wasm_api := js.Global().Get("Object").New()
    js.Global().Set("wasm_api", wasm_api)

    wasm_signals := js.Global().Get("Object").New()
    wasm_api.Set("signals", wasm_signals)

    tf := editor.NewTextField()
    render := func () func() {
        sig := signals.NewSignal()
        sig.AttachToJS(wasm_signals, "text_field_content")
        return func() {
            sig.Data <- strings.Join(tf.Render(), "<br>")
        }
    }()

    go func() {
        sig := signals.NewSignal()
        sig.AttachToJS(wasm_signals, "keypress")
        for {
            event := (<-sig.Data).(js.Value)
            tf.HandleHandleKeyPress(
                    event.Get("key").String(),
                    event.Get("ctrlKey").Bool(),
                    event.Get("shiftKey").Bool(),
                    event.Get("altKey").Bool(),
            )
            render()
        }
    }()

    <-make(chan struct{})
}
