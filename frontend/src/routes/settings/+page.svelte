<script lang="ts">
  import Container from '$lib/layout/Container.svelte'
  import { pageTitle, useDarkModeStore } from '$lib/layout/stores.svelte'
  import { Heading, Toggle } from 'flowbite-svelte'

  const { isDarkMode, toggleDarkMode } = useDarkModeStore()

  pageTitle.value = 'Settings'
</script>

<Container className="flex justify-center items-center !pt-16">
  <div class="md:w-2/3 w-full">
    <Heading tag="h1">Settings</Heading>

    <div class="mt-12 flex-column">
      <div class="mb-10">
        <Heading tag="h3" class="mb-8">Appearance</Heading>
        <Toggle size="large" checked={isDarkMode.value} on:change={toggleDarkMode}>DarkMode</Toggle>
      </div>
    </div>
  </div>
</Container>
