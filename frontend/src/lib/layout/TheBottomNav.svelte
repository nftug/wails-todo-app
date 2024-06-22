<script lang="ts">
  import { page } from '$app/stores'
  import { BottomNav, BottomNavItem } from 'flowbite-svelte'
  import { HomeSolid, InfoCircleSolid } from 'flowbite-svelte-icons'

  type Props = { height: number }
  let { height = $bindable(0) }: Props = $props()

  const activeUrl = $derived($page.url.pathname)

  let footer = $state<HTMLElement | null>()
  $effect(() => {
    footer = document.getElementById('footer')
  })

  const onresize = () => (height = footer?.offsetHeight ?? 0)
</script>

<svelte:window {onresize} />

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
