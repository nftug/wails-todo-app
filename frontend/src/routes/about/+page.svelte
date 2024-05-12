<script lang="ts">
  import Container from '$lib/layout/Container.svelte'
  import { SITE_TITLE, pageTitle } from '$lib/layout/stores'
  import { A, Heading, P } from 'flowbite-svelte'

  $pageTitle = 'About'
</script>

<Container>
  <Heading tag="h1">Welcome to {SITE_TITLE}</Heading>
  <P class="mt-5 ml-1">
    Visit <A href="https://kit.svelte.dev">kit.svelte.dev</A> to read the documentation
  </P>
</Container>
