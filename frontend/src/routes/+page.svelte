<script lang="ts">
  import { AddFile, Select, Table } from "$widgets";
  import { Csv } from "../shared/utils/сsv";
  import { readFile } from "../shared/utils/readFile";
  import { ClipboardCopy } from "@lucide/svelte";
  import { Copy } from "$lib/wailsjs/go/app";

  let files: File[] = $state([]);
  let data: string[][] = $state([]);
  let fields: string[] = $state([]);
  let filtered = $derived(filterData(data, fields));

  function filterData(data: string[][], fields: string[]) {
    if (!data || !data.length || !data[0].length) return [];

    let newData: string[][] = [];
    let cols = [];
    for (let i = 0; i < data[0].length; i++) {
      if (fields.includes(data[0][i])) {
        cols.push(i);
      }
    }

    for (let i = 0; i < data.length; i++) {
      newData[i] = [];
      for (let j = 0; j < data[i].length; j++) {
        if (cols.includes(j)) {
          newData[i].push(data[i][j]);
        }
      }
    }
    return newData;
  }

  $effect(() => {
    handleFiles();
  });

  async function handleFiles() {
    if (!files.length) return;
    await Promise.all(files.map((f) => readFile(f)))
      .then((res) => {
        return res.reduce((acc, el) => {
          if (!el) return acc;
          let newCsv = new Csv(el);
          if (!acc.data.length) return newCsv;
          acc.join(newCsv);
          return acc;
        }, new Csv());
      })
      .then((res) => (data = res.data));
  }

  async function handleCopy() {
    if (!Array.isArray(data[0])) return;
    let text = data.map((el) => el.join("\t")).join("\n");
    Copy(text);
  }
</script>

<div class="h-full flex flex-col">
  <div class="flex gap-4">
    <Select bind:value={fields} options={data[0]} />
  </div>
  <div class=" bg-black overflow-scroll"><Table data={filtered} /></div>
  <div class="mt-auto w-full flex justify-center p-2 gap-2">
    <AddFile bind:files />
    <div class="border rounded flex items-center justify-center">
      <button onclick={handleCopy} class="cursor-pointer p-2">
        <ClipboardCopy class="w-8 h-8" />
      </button>
    </div>
  </div>
</div>
