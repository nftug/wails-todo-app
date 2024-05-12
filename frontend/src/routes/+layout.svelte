<script context="module" lang="ts">
  export const HeaderHeightContext = Symbol('header-height-getter')
  export const FooterHeightContext = Symbol('footer-height-getter')
  export type HeightStore = Readable<number>
</script>

<script lang="ts">
  import { pageTitle, SITE_TITLE } from '$lib/layout/stores'
  import TheBottomNav from '$lib/layout/TheBottomNav.svelte'
  import TheDrawer from '$lib/layout/TheDrawer.svelte'
  import TheHeader from '$lib/layout/TheHeader.svelte'
  import { setContext } from 'svelte'
  import { derived, writable, type Readable } from 'svelte/store'
  import '../app.css'

  $: title = $pageTitle ? `${$pageTitle} - ${SITE_TITLE}` : SITE_TITLE

  const headerHeight = writable(0)
  const footerHeight = writable(0)
  const headerHeightDerived = derived(headerHeight, (x) => x)
  const footerHeightDerived = derived(footerHeight, (x) => x)

  setContext<HeightStore>(HeaderHeightContext, headerHeightDerived)
  setContext<HeightStore>(FooterHeightContext, footerHeightDerived)
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<TheHeader bind:height={$headerHeight} />

<TheDrawer />

<main>
  <slot />
</main>

<TheBottomNav bind:height={$footerHeight} />
