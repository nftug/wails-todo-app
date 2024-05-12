<script lang="ts">
  import { page } from '$app/stores'
  import { BottomNav, BottomNavItem } from 'flowbite-svelte'
  import { HomeSolid, InfoCircleSolid } from 'flowbite-svelte-icons'

  $: activeUrl = $page.url.pathname
</script>

<footer id="footer">
  <BottomNav {activeUrl} classInner="grid-cols-2" classOuter="md:hidden z-40 h-14">
    <BottomNavItem btnName="Home" href="/">
      <HomeSolid />
    </BottomNavItem>
    <BottomNavItem btnName="About" href="/about">
      <InfoCircleSolid />
    </BottomNavItem>
  </BottomNav>
</footer>
