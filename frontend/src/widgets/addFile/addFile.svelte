<script lang="ts">
  import { FilePlusCorner } from "@lucide/svelte";

  let { files = $bindable([]) }: { files?: File[]; "bind:files"?: File[] } = $props();

  function handleDragOver(e: DragEvent) {
    e.preventDefault();
  }

  function handleDrop(e: DragEvent) {
    e.preventDefault();
    if (!e.dataTransfer) return;
    files.push(...Array.from(e.dataTransfer.files));
  }

  function handleInput(e: Event) {
    if (!e.target) return;
    files.push(...Array.from((e.target as HTMLInputElement).files!));
  }
</script>

<label ondrop={handleDrop} ondragover={handleDragOver}>
  <input type="file" multiple hidden onchange={handleInput} />
  <div class="flex gap-2 border rounded w-fit px-2 py-2 items-center justify-center cursor-pointer">
    <FilePlusCorner class="w-10 h-10" />
    <span>Add files</span>
  </div>
</label>
