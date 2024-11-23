<script lang="ts">
  import ConfirmDialog, { type DialogProps } from '$lib/common/ConfirmDialog.svelte'
  import CenteredContainer from '$lib/layout/CenteredContainer.svelte'
  import { Button, Heading } from 'svelte-5-ui-lib'

  let dialog = $state<ConfirmDialog>()
  let dialogProps = $state<DialogProps>({})

  async function onClickButton() {
    if (!dialog) return
    dialogProps = { message: 'Are you sure you want to continue?' }
    await dialog.openDialog()
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
