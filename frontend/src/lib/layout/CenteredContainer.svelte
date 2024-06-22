<script lang="ts">
  import { type Snippet } from 'svelte'
  import { restHeight } from './stores.svelte'

  type Props = { children: Snippet }
  let { children }: Props = $props()

  let innerHeight = $state(0)
  const contentHeight = $derived(innerHeight - restHeight.value.header - restHeight.value.footer)
</script>

<svelte:window bind:innerHeight />

<div
  class="w-screen flex justify-center items-center text-center flex-col overflow-y-auto hidden-scroll-bar"
  style={contentHeight ? `height: ${contentHeight}px` : null}
>
  {@render children()}
</div>
