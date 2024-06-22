export const SITE_TITLE = 'Wails Todo App'

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
