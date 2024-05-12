<script lang="ts">
  export let className = ''

  $: containerClass = 'm-10 my-14' + (className ? ` ${className}` : '')
</script>

<div class={containerClass}>
  <slot />
</div>
