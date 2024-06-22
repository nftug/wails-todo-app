<script lang="ts">
  import { SITE_TITLE } from '$lib'
  import CenteredContainer from '$lib/layout/CenteredContainer.svelte'
  import { pageTitle } from '$lib/layout/stores.svelte'
  import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime'
  import { A, Heading, P } from 'flowbite-svelte'

  pageTitle.value = 'About'
</script>

<CenteredContainer>
  <Heading tag="h1">Welcome to {SITE_TITLE}</Heading>
  <P class="mt-8 ml-1">
    Visit
    <A onclick={() => BrowserOpenURL('https://kit.svelte.dev')}>kit.svelte.dev</A>
    to read the documentation of SvelteKit
  </P>
</CenteredContainer>
