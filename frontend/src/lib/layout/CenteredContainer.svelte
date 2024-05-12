<script lang="ts">
  import { footerHeight, headerHeight } from './stores'

  let innerHeight: number
  let contentHeight: number

  $: contentHeight = innerHeight - $headerHeight - $footerHeight
</script>

<svelte:window bind:innerHeight />

<div
  class="w-screen flex justify-center items-center text-center flex-col overflow-y-auto hidden-scroll-bar"
  style={contentHeight ? `height: ${contentHeight}px` : null}
>
  <slot />
</div>
