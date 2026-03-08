<script lang="ts">
  import { AddFile, Select, Table } from "$widgets";
  import { Csv } from "../utils/сsv";
  import { readFile } from "../utils/readFile";

  let files: File[] = $state([]);
  let data: string[][] = $state([]);

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
</script>

<div class="h-full flex flex-col">
  <div class="">
    <Select value="Test" options={["test1", "test2"]} />
  </div>
  <div class=" bg-black overflow-scroll"><Table {data} /></div>
  <div class="mt-auto w-full flex justify-center p-2">
    <AddFile bind:files />
  </div>
</div>
