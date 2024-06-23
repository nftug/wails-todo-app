import { browser } from '$app/environment'
import { writable } from '$lib/util/util.svelte'

export const pageTitle = writable('')
export const drawerHidden = writable(true)

let isDarkModeState = $state<boolean>()
export const useDarkModeStore = () => ({
  // prettier-ignore
  isDarkMode: {
    get value() { return isDarkModeState }
  },
  initIsDarkMode: () => {
    if (!browser) return
    isDarkModeState =
      localStorage.getItem('color-theme') === 'dark' ||
      document.documentElement.classList.contains('dark')
  },
  toggleDarkMode: () => {
    isDarkModeState = document.documentElement.classList.toggle('dark')
    localStorage.setItem('color-theme', isDarkModeState ? 'dark' : 'light')
  }
})

export const headerHeight = writable(0)
export const footerHeight = writable(0)
