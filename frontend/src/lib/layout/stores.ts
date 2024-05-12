import { writable } from 'svelte/store'

export const SITE_TITLE = 'Hello SvelteKit'

export const pageTitle = writable<string | null>()

export const drawerHidden = writable(true)
