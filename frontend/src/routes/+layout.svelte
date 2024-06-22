<script context="module" lang="ts">
  export const HeaderHeightContext = Symbol('header-height-getter')
  export const FooterHeightContext = Symbol('footer-height-getter')
  export type HeightReadable = { value: number }
</script>

<script lang="ts">
  import { beforeNavigate } from '$app/navigation'
  import { SITE_TITLE } from '$lib'
  import TheBottomNav from '$lib/layout/TheBottomNav.svelte'
  import TheDrawer from '$lib/layout/TheDrawer.svelte'
  import TheHeader from '$lib/layout/TheHeader.svelte'
  import { pageTitle, useDarkModeStore } from '$lib/layout/stores.svelte'
  import { DarkMode } from 'flowbite-svelte'
  import { setContext, type Snippet } from 'svelte'
  import '../app.css'

  type Props = { children: Snippet }
  const { children }: Props = $props()

  const title = $derived(pageTitle.value ? `${pageTitle.value} - ${SITE_TITLE}` : SITE_TITLE)
  beforeNavigate(() => (pageTitle.value = ''))

  let headerHeight = $state(0)
  let footerHeight = $state(0)

  setContext<HeightReadable>(HeaderHeightContext, {
    get value() {
      return headerHeight
    }
  })
  setContext<HeightReadable>(FooterHeightContext, {
    get value() {
      return footerHeight
    }
  })

  const { setIsDarkMode } = useDarkModeStore()
  $effect(() => setIsDarkMode())
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<TheHeader bind:height={headerHeight} />

<TheDrawer />

<main>
  {@render children()}
</main>

<TheBottomNav bind:height={footerHeight} />

<DarkMode class="hidden" />
