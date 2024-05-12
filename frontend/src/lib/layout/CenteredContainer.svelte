<script lang="ts">
  import {
    FooterHeightContext,
    HeaderHeightContext,
    type HeightStore
  } from '$routes/+layout.svelte'
  import { getContext } from 'svelte'

  let innerHeight: number
  const headerHeight = getContext<HeightStore>(HeaderHeightContext)
  const footerHeight = getContext<HeightStore>(FooterHeightContext)

  $: contentHeight = innerHeight - $headerHeight - $footerHeight
</script>

<svelte:window bind:innerHeight />

<div
  class="w-screen flex justify-center items-center text-center flex-col overflow-y-auto hidden-scroll-bar"
  style={contentHeight ? `height: ${contentHeight}px` : null}
>
  <slot />
</div>
