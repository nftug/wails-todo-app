<script lang="ts">
  import { beforeNavigate } from '$app/navigation'
  import { SITE_TITLE } from '$lib'
  import TheBottomNav from '$lib/layout/TheBottomNav.svelte'
  import TheDrawer from '$lib/layout/TheDrawer.svelte'
  import TheHeader from '$lib/layout/TheHeader.svelte'
  import { pageTitle, useDarkModeStore } from '$lib/layout/stores.svelte'
  import { DarkMode } from 'flowbite-svelte'
  import { type Snippet } from 'svelte'
  import '../app.css'

  type Props = { children: Snippet }
  const { children }: Props = $props()

  const title = $derived(pageTitle.value ? `${pageTitle.value} - ${SITE_TITLE}` : SITE_TITLE)
  beforeNavigate(() => (pageTitle.value = ''))

  const { initIsDarkMode } = useDarkModeStore()
  $effect(() => initIsDarkMode())
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<TheHeader />

<TheDrawer />

<main>
  {@render children()}
</main>

<TheBottomNav />

<DarkMode class="hidden" />
