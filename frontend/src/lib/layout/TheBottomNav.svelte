<script lang="ts">
  import { page } from '$app/stores'
  import { HomeSolid, InfoCircleSolid } from 'flowbite-svelte-icons'
  import { BottomNav, BottomNavItem } from 'svelte-5-ui-lib'
  import { footerHeight } from './stores.svelte'

  const activeUrl = $derived($page.url.pathname)

  let footer = $state<HTMLElement | null>()
  $effect(() => {
    footer = document.getElementById('footer')
  })

  const onresize = () => (footerHeight.value = footer?.offsetHeight ?? 0)
</script>

<svelte:window {onresize} />

<footer>
  <BottomNav id="footer" {activeUrl} innerClass="grid-cols-2" outerClass="md:hidden z-40 h-14">
    <BottomNavItem btnName="Home" href="/">
      <HomeSolid />
    </BottomNavItem>
    <BottomNavItem btnName="About" href="/about">
      <InfoCircleSolid />
    </BottomNavItem>
  </BottomNav>
</footer>
