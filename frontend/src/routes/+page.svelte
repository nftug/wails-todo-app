<script lang="ts">
  import ConfirmDialog, { type DialogProps } from '$lib/common/ConfirmDialog.svelte'
  import CenteredContainer from '$lib/layout/CenteredContainer.svelte'
  import { Button, Heading } from 'flowbite-svelte'

  let dialog = $state<ConfirmDialog>()
  let dialogProps = $state<DialogProps>({})

  async function onClickButton() {
    if (!dialog) return
    dialogProps = { message: 'Are you sure you want to continue?' }
    const ans = await dialog.openDialog()
    alert(`Your answer: ${ans}`)
  }
</script>

<CenteredContainer>
  <Heading tag="h1">Welcome!</Heading>

  <div class="mt-16">
    <Button onclick={onClickButton}>Dialog</Button>
  </div>
</CenteredContainer>

<ConfirmDialog bind:this={dialog} {...dialogProps}>
  {#snippet icon()}{/snippet}
</ConfirmDialog>
