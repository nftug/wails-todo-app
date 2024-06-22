<script lang="ts">
  import { page } from '$app/stores'
  import { SITE_TITLE } from '$lib'
  import { Button, NavBrand, NavHamburger, NavLi, NavUl, Navbar, Search } from 'flowbite-svelte'
  import { SearchOutline } from 'flowbite-svelte-icons'
  import { drawerHidden } from './stores.svelte'

  type Props = { height: number }
  let { height = $bindable(0) }: Props = $props()

  const activeUrl = $derived($page.url.pathname)
</script>

<header class="sticky top-0 z-40 flex-none w-full" id="header" bind:offsetHeight={height}>
  <Navbar fluid class="dark:bg-gray-900">
    <NavHamburger
      class="m-0 ms-2 me-4 md:block"
      onClick={() => (drawerHidden.value = !drawerHidden.value)}
    />

    <NavBrand href="/">
      <span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
        {SITE_TITLE}
      </span>
    </NavBrand>

    <div class="flex md:order-3">
      <Button
        onclick={() => alert('search')}
        color="none"
        data-collapse-toggle="mobile-menu-3"
        aria-controls="mobile-menu-3"
        aria-expanded="false"
        class="
          text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700
          focus:outline-none focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700
          md:hidden rounded-lg text-sm p-2.5 me-1
        "
      >
        <SearchOutline class="w-5 h-5" />
      </Button>

      <div class="hidden relative md:block">
        <Search class="w-72" size="md" />
      </div>
    </div>

    <NavUl {activeUrl} class="mx-auto">
      <NavLi href="/">Home</NavLi>
      <NavLi href="/about">About</NavLi>
    </NavUl>
  </Navbar>
</header>
