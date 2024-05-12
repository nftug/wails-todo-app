<script lang="ts">
  import { Greet } from '$lib/wailsjs/go/main/App.js'
  import { Button, Heading, Input, P } from 'flowbite-svelte'
  import { onMount } from 'svelte'

  let resultText: string = 'Please enter your name below 👇'
  let name: string

  function greet(): void {
    Greet(name).then((result) => (resultText = result))
  }

  let contentHeightPx: number

  onMount(() => {
    onResizeWindow()
  })

  function onResizeWindow() {
    const header = document.getElementById('header')
    const footer = document.getElementById('footer')
    if (!header || !footer) return
    contentHeightPx = window.innerHeight - header.offsetHeight - footer.offsetHeight
  }
</script>

<svelte:window on:resize={onResizeWindow} />

<div
  class="h-full w-screen flex justify-center items-center text-center flex-col"
  style={contentHeightPx ? `height: ${contentHeightPx}px` : null}
>
  <Heading tag="h1">Welcome to Wails</Heading>

  <P class="my-8">{resultText}</P>

  <div class="flex space-x-4">
    <Input type="text" bind:value={name} placeholder="Your name" />
    <Button on:click={greet}>Greet</Button>
  </div>
</div>
