import { writable } from 'svelte/store'

export const SITE_TITLE = 'Wails Note App'

export const pageTitle = writable<string | null>()

export const drawerHidden = writable(true)
