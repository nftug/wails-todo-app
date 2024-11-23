<script lang="ts">
  import { SITE_TITLE } from '$lib'
  import { SearchOutline } from 'flowbite-svelte-icons'
  import { Button, NavBrand, NavLi, NavUl, Navbar, Search, SidebarButton } from 'svelte-5-ui-lib'
  import { headerHeight, useDrawerStore } from './stores.svelte'

  const { toggleDrawer } = useDrawerStore()
</script>

<header
  class="sticky top-0 z-40 flex-none w-full"
  id="header"
  bind:offsetHeight={headerHeight.value}
>
  <Navbar fluid class="dark:bg-gray-900">
    <SidebarButton class="m-0 ms-2 me-4 md:block" onclick={toggleDrawer} />

    <NavBrand siteName={SITE_TITLE} href="/"></NavBrand>

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

    <NavUl>
      <NavLi href="/">Home</NavLi>
      <NavLi href="/about">About</NavLi>
    </NavUl>
  </Navbar>
</header>
