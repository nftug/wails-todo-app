export const writable = <T>(initial: T) => {
  let state = $state<T>(initial)
  // prettier-ignore
  return {
    get value() { return state },
    set value(v) { state = v }
  }
}

export const debounceDerived = <T>(getValue: () => T, waitFor: number = 500) => {
  let result = $state(getValue())
  let timeout: ReturnType<typeof setTimeout>

  $effect(() => {
    const newValue = getValue()
    if (newValue === result) return
    clearTimeout(timeout)
    timeout = setTimeout(() => (result = newValue), waitFor)
  })

  return {
    // prettier-ignore
    get value() { return result }
  }
}
