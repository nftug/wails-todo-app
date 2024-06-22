<script lang="ts" context="module">
  type ButtonText = { yesText?: string; noText?: string }
  type DialogResult = 'yes' | 'no'
  export type DialogProps = { message?: string; icon?: Snippet } & ButtonText
</script>

<script lang="ts">
  import { Button, Modal } from 'flowbite-svelte'
  import { ExclamationCircleOutline } from 'flowbite-svelte-icons'
  import { type Snippet } from 'svelte'

  const { message, icon, yesText = 'Yes', noText = 'No' }: DialogProps = $props()

  let open = $state(false)
  let returnResult: (result: DialogResult) => void

  export async function openDialog() {
    open = true
    const result = await new Promise<DialogResult>((resolve) => (returnResult = resolve))
    open = false
    return result
  }
</script>

<Modal bind:open size="xs" dismissable={false}>
  <div class="text-center">
    {#if icon}
      {@render icon()}
    {:else}
      <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
    {/if}

    <h3 class="mb-5 text-md font-normal whitespace-pre-wrap">
      {message}
    </h3>

    <Button color="red" class="me-2" on:click={() => returnResult('yes')}>
      {yesText}
    </Button>
    <Button color="alternative" on:click={() => returnResult('no')}>
      {noText}
    </Button>
  </div>
</Modal>
