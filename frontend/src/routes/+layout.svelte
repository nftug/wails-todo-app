<script context="module" lang="ts">
  export const HeaderHeightContext = Symbol('header-height-getter')
  export const FooterHeightContext = Symbol('footer-height-getter')
  export type HeightReadable = { value: number }
</script>

<script lang="ts">
  import { pageTitle, SITE_TITLE } from '$lib/layout/stores.svelte'
  import TheBottomNav from '$lib/layout/TheBottomNav.svelte'
  import TheDrawer from '$lib/layout/TheDrawer.svelte'
  import TheHeader from '$lib/layout/TheHeader.svelte'
  import { setContext, type Snippet } from 'svelte'
  import '../app.css'

  type Props = { children: Snippet }
  const { children }: Props = $props()

  const title = $derived(pageTitle.value ? `${pageTitle.value} - ${SITE_TITLE}` : SITE_TITLE)

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
