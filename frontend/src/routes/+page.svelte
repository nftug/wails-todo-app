<script lang="ts">
  import ConfirmDialog from '$lib/common/ConfirmDialog.svelte'
  import CenteredContainer from '$lib/layout/CenteredContainer.svelte'
  import { Button, Heading } from 'flowbite-svelte'

  let dialog = $state<ConfirmDialog>()

  async function onClickButton() {
    if (!dialog) return
    const ans = await dialog.openDialog({
      message: 'Are you sure you want to continue?',
      iconSnippet
    })
    alert(`Your answer: ${ans}`)
  }
</script>

<CenteredContainer>
  <Heading tag="h1">Welcome!</Heading>

  <div class="mt-16">
    <Button onclick={onClickButton}>Dialog</Button>
  </div>
</CenteredContainer>

<ConfirmDialog bind:this={dialog} />

{#snippet iconSnippet()}{/snippet}
