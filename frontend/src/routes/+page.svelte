<script lang="ts">
  import CenteredContainer from '$lib/layout/CenteredContainer.svelte'
  import { Greet, ShowErrorMessage } from '$lib/wailsjs/go/main/App.js'
  import { Button, Heading, Input, P } from 'flowbite-svelte'

  let resultText: string = 'Please enter your name below 👇'
  let name: string
  let nameInput: HTMLInputElement

  async function greet() {
    try {
      resultText = await Greet(name)
    } catch (e) {
      await ShowErrorMessage(e as string)
      nameInput.focus()
    }

    name = ''
  }
</script>

<CenteredContainer>
  <Heading tag="h1">Welcome!</Heading>

  <P class="my-8">{resultText}</P>

  <form on:submit|preventDefault={greet} class="flex space-x-4">
    <Input let:props placeholder="Your name">
      <input type="text" {...props} bind:value={name} bind:this={nameInput} />
    </Input>
    <Button on:click={greet}>Greet</Button>
  </form>
</CenteredContainer>
