<script lang="ts">
  import CenteredContainer from '$lib/layout/CenteredContainer.svelte'
  import { pageTitle } from '$lib/layout/stores.svelte'
  import { Heading } from 'flowbite-svelte'

  pageTitle.value = ''
</script>

<CenteredContainer>
  <Heading tag="h1">Welcome!</Heading>
</CenteredContainer>
