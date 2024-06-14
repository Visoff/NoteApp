/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {
            colors: {
                "mantle": "var(--mantle)",
                "crust": "var(--crust)",
                "base": "var(--base)",
                "surface0": "var(--surface0)",
                "surface1": "var(--surface1)",
                "text": "var(--text)",
            }
        },
    },
    plugins: [],
}
