import { writable } from "svelte/store";

let wasm_ready = false;;

export function init() {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("/wasm/main.wasm"), go.importObject).then((result) => {
        go.run(result.instance)
        console.log("ready")
        wasm_ready = true;            
    });
}

/** @typedef {[(callback: (data: any) => void) => void, (data: any) => void]} signal */
/** @param {string} signal */
export function signal_to_writable(signal) {
    const wr = writable(undefined);
    if (typeof window === "undefined") {
        return wr
    }
    const indx = setInterval(() => {
        if (!wasm_ready) {
            return
        }
        if (!(signal in wasm_api.signals)) {
            return
        }
        /** @type {signal} */
        const [getter, _] = wasm_api.signals[signal];
        getter((data) => {
            wr.set(data);
        })
        clearInterval(indx)
    }, 100)
    return wr
}

/** @param {string} signal */
export function signal_setter(signal) {
    if (typeof window === "undefined") {
        return
    }
    if (!wasm_ready) {
        return
    }
    if (!(signal in wasm_api.signals)) {
        return
    }
    /** @type {signal} */
    const [_, setter] = wasm_api.signals[signal];
    return setter
}
