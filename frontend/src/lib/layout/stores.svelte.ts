import { browser } from '$app/environment'

let pageTitleState = $state('')
// prettier-ignore
export const pageTitle = {
  get value() { return pageTitleState },
  set value(v) { pageTitleState = v }
}

let drawerHiddenState = $state(true)
// prettier-ignore
export const drawerHidden = {
  get value() { return drawerHiddenState },
  set value(v) { drawerHiddenState = v }
}

let isDarkModeState = $state<boolean>()
export const useDarkModeStore = () => ({
  // prettier-ignore
  isDarkMode: {
    get value() { return isDarkModeState },
    set value(v) { isDarkModeState = v }
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

let restHeightState = $state({ header: 0, footer: 0 })
// prettier-ignore
export const restHeight = {
  get value() { return restHeightState },
  set value(v) { restHeightState = v }
}
