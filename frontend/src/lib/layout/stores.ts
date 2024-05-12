import { writable } from 'svelte/store'

export const SITE_TITLE = 'My Wails App'

export const pageTitle = writable<string | null>()

export const drawerHidden = writable(true)

export const headerHeight = writable(0)

export const footerHeight = writable(0)
