import { writable } from 'svelte/store';


/** @typedef {{type: "file", mimetype: string, content: string, name: string}|{type: "dir", name: string, children: FSEntry[]}} FSEntry */
/** @type {import('svelte/store').Writable<{tree:FSEntry, path:string}|undefined>} */
export const fs = writable(undefined);
