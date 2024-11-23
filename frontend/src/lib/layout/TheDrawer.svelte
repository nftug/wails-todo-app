<script lang="ts">
  import { onNavigate } from '$app/navigation'
  import { HomeSolid, InfoCircleSolid } from 'flowbite-svelte-icons'
  import { Drawer, Drawerhead, SidebarGroup, SidebarItem } from 'svelte-5-ui-lib'
  import { useDrawerStore } from './stores.svelte'

  const { drawerStatus, closeDrawer } = useDrawerStore()

  onNavigate(() => {
    closeDrawer()
  })

  const navIconClass =
    'h-5 w-5 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white'
</script>

<Drawer drawerStatus={drawerStatus.value} {closeDrawer}>
  <Drawerhead onclick={closeDrawer} class="mb-4">
    <h5
      id="drawer-label"
      class="inline-flex items-center text-xl font-semibold text-gray-500 dark:text-gray-400"
    >
      <InfoCircleSolid class="me-2.5 h-5 w-5" />Menu
    </h5>
  </Drawerhead>

  <SidebarGroup>
    <SidebarItem label="Home" href="/">
      {#snippet iconSlot()}
        <HomeSolid class={navIconClass} />
      {/snippet}
    </SidebarItem>
    <SidebarItem label="About" href="/about">
      {#snippet iconSlot()}
        <InfoCircleSolid class={navIconClass} />
      {/snippet}
    </SidebarItem>
  </SidebarGroup>
</Drawer>
