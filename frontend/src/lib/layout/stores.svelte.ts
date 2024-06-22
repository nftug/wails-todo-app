import { browser } from '$app/environment'

let pageTitleState = $state('')
export const pageTitle = {
  get value() {
    return pageTitleState
  },
  set value(v) {
    pageTitleState = v
  }
}

let drawerHiddenState = $state(true)
export const drawerHidden = {
  get value() {
    return drawerHiddenState
  },
  set value(v) {
    drawerHiddenState = v
  }
}

let isDarkModeState = $state<boolean>()
export const useDarkModeStore = () => {
  const setIsDarkMode = () => {
    if (!browser) return

    isDarkModeState =
      localStorage.getItem('color-theme') === 'dark' ||
      document.documentElement.classList.contains('dark')
  }

  const toggleDarkMode = () => {
    isDarkModeState = document.documentElement.classList.toggle('dark')
    localStorage.setItem('color-theme', isDarkModeState ? 'dark' : 'light')
  }

  return {
    isDarkMode: {
      get value() {
        return isDarkModeState
      },
      set value(v) {
        isDarkModeState = v
      }
    },
    setIsDarkMode,
    toggleDarkMode
  }
}
