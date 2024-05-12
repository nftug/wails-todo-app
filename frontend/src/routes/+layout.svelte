<script lang="ts">
  import { pageTitle, SITE_TITLE } from '$lib/layout/stores'
  import TheBottomNav from '$lib/layout/TheBottomNav.svelte'
  import TheDrawer from '$lib/layout/TheDrawer.svelte'
  import TheHeader from '$lib/layout/TheHeader.svelte'
  import '../app.css'

  $: title = $pageTitle ? `${$pageTitle} - ${SITE_TITLE}` : SITE_TITLE
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<TheHeader />

<TheDrawer />

<main>
  <slot />
</main>

<TheBottomNav />
