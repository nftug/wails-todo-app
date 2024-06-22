<script lang="ts">
  import { Button, Modal } from 'flowbite-svelte'
  import { ExclamationCircleOutline } from 'flowbite-svelte-icons'
  import { type Snippet } from 'svelte'

  type ButtonText = { yesText?: string; noText?: string }
  type DialogResult = 'yes' | 'no'
  type Props = { message: string; iconSnippet?: Snippet } & ButtonText
  let props = $state<Props>()

  const open = $derived(!!props)
  let returnResult: (result: PromiseLike<DialogResult> | DialogResult) => void

  export async function openDialog({
    message,
    iconSnippet = undefined,
    yesText = 'Yes',
    noText = 'No'
  }: Props) {
    props = { message, iconSnippet, yesText, noText }
    const result = await new Promise<DialogResult>((resolve) => (returnResult = resolve))
    props = undefined
    return result
  }
</script>

<Modal {open} size="xs" dismissable={false}>
  <div class="text-center">
    {#if props?.iconSnippet}
      {@render props.iconSnippet()}
    {:else}
      <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
    {/if}

    <h3 class="mb-5 text-md font-normal whitespace-pre-wrap">
      {props?.message}
    </h3>

    <Button color="red" class="me-2" on:click={() => returnResult('yes')}>
      {props?.yesText}
    </Button>
    <Button color="alternative" on:click={() => returnResult('no')}>
      {props?.noText}
    </Button>
  </div>
</Modal>
