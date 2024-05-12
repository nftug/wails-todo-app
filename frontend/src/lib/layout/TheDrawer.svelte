<script lang="ts">
  import { onNavigate } from '$app/navigation'
  import { page } from '$app/stores'
  import {
    CloseButton,
    Drawer,
    DropdownDivider,
    Sidebar,
    SidebarGroup,
    SidebarItem,
    SidebarWrapper
  } from 'flowbite-svelte'
  import { CogSolid, HomeSolid, InfoCircleOutline } from 'flowbite-svelte-icons'
  import { sineIn } from 'svelte/easing'
  import { drawerHidden } from './stores'

  let transitionParams = {
    x: -320,
    duration: 200,
    easing: sineIn
  }

  onNavigate(() => {
    $drawerHidden = true
  })

  $: activeUrl = $page.url.pathname
</script>

<Drawer transitionType="fly" bind:hidden={$drawerHidden} {transitionParams}>
  <div class="flex items-center">
    <h5
      id="drawer-navigation-label-3"
      class="text-base font-semibold text-gray-500 uppercase dark:text-gray-400"
    >
      Menu
    </h5>
    <CloseButton on:click={() => ($drawerHidden = true)} class="mb-4 dark:text-white" />
  </div>

  <Sidebar {activeUrl}>
    <SidebarWrapper divClass="overflow-y-auto py-4 px-3 rounded dark:bg-gray-800">
      <SidebarGroup>
        {@const navIconClass = `w-5 h-5 text-gray-500 transition duration-75
          dark:text-gray-400group-hover:text-gray-900 dark:group-hover:text-white`}

        <SidebarItem label="Home" href="/">
          <svelte:fragment slot="icon">
            <HomeSolid class={navIconClass} />
          </svelte:fragment>
        </SidebarItem>
        <SidebarItem label="About" href="/about">
          <svelte:fragment slot="icon">
            <InfoCircleOutline class={navIconClass} />
          </svelte:fragment>
        </SidebarItem>

        <DropdownDivider />

        <SidebarItem label="Setting">
          <svelte:fragment slot="icon">
            <CogSolid class={navIconClass} />
          </svelte:fragment>
        </SidebarItem>
      </SidebarGroup>
    </SidebarWrapper>
  </Sidebar>
</Drawer>
