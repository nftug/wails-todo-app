import { browser } from '$app/environment'
import { writable } from '$lib/util/util.svelte'
import { uiHelpers } from 'svelte-5-ui-lib'

export const pageTitle = writable('')

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

let drawerStatus = $state<boolean>(false)
const drawer = uiHelpers()

export const useDrawerStore = () => {
  $effect(() => {
    drawerStatus = drawer.isOpen
  })

  return {
    // prettier-ignore
    drawerStatus: {
      get value() { return drawerStatus }
    },
    closeDrawer: drawer.close,
    toggleDrawer: drawer.toggle
  }
}

export const headerHeight = writable(0)
export const footerHeight = writable(0)
