<script lang="ts">
  import CenteredContainer from '$lib/layout/CenteredContainer.svelte'
  import { Greet, ShowErrorMessage } from '$lib/wailsjs/go/main/App.js'
  import { Button, Heading, Input, P } from 'flowbite-svelte'

  let resultText: string = 'Please enter your name below 👇'
  let name: string

  async function greet() {
    try {
      resultText = await Greet(name)
      name = ''
    } catch (e) {
      ShowErrorMessage(e as string)
    }
  }
</script>

<CenteredContainer>
  <Heading tag="h1">Welcome!</Heading>

  <P class="my-8">{resultText}</P>

  <form on:submit|preventDefault={greet} class="flex space-x-4">
    <Input type="text" bind:value={name} placeholder="Your name" />
    <Button on:click={greet}>Greet</Button>
  </form>
</CenteredContainer>
