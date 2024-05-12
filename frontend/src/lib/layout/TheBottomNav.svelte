<script lang="ts">
  import { page } from '$app/stores'
  import { BottomNav, BottomNavItem } from 'flowbite-svelte'
  import { HomeSolid, InfoCircleSolid } from 'flowbite-svelte-icons'
  import { onMount } from 'svelte'
  import { footerHeight } from './stores'

  $: activeUrl = $page.url.pathname

  let footer: HTMLElement | null

  onMount(() => {
    footer = document.getElementById('footer')
  })

  function setFooterHeight() {
    $footerHeight = footer?.offsetHeight ?? 0
  }
</script>

<svelte:window on:resize={setFooterHeight} />

<footer>
  <BottomNav id="footer" {activeUrl} classInner="grid-cols-2" classOuter="md:hidden z-40 h-14">
    <BottomNavItem btnName="Home" href="/">
      <HomeSolid />
    </BottomNavItem>
    <BottomNavItem btnName="About" href="/about">
      <InfoCircleSolid />
    </BottomNavItem>
  </BottomNav>
</footer>
