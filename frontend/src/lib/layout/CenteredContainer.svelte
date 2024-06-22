<script lang="ts">
  import { RestHeightContext, type RestHeight } from '$routes/+layout.svelte'
  import { getContext, type Snippet } from 'svelte'

  type Props = { children: Snippet }
  let { children }: Props = $props()

  let innerHeight = $state(0)
  const restHeight = getContext<RestHeight>(RestHeightContext)

  const contentHeight = $derived(innerHeight - restHeight?.header - restHeight?.footer)
</script>

<svelte:window bind:innerHeight />

<div
  class="w-screen flex justify-center items-center text-center flex-col overflow-y-auto hidden-scroll-bar"
  style={contentHeight ? `height: ${contentHeight}px` : null}
>
  {@render children()}
</div>
