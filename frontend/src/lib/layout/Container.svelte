<script lang="ts">
  import type { Snippet } from 'svelte'

  type Props = {
    children: Snippet
    className?: string
  }
  const { children, className }: Props = $props()

  const containerClass = $derived('p-10 md:py-8' + (className ? ` ${className}` : ''))
</script>

<div class={containerClass}>
  {@render children()}
</div>
