<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";
  import { cn } from "@sglara/cn";
  interface HtmlInputProps extends HTMLInputAttributes {
    type?: "text" | "password" | "email" | "number" | "search" | "tel" | "url";
    value: string;
  }
  let {
    value = $bindable(),
    type = "text",
    ...props
  }: HtmlInputProps = $props();
</script>

<input
  {type}
  bind:value
  {...props}
  class={cn(
    "w-full border hover:border-zinc-300 text-sm transition-all focus-within:border-blue-800 active:border-blue-800 focus:border-blue-800 border-zinc-300 bg-zinc-50 outline-0 rounded-md px-3 py-2 disabled:bg-zinc-200 disabled:text-zinc-500 disabled:border-zinc-200 disabled:cursor-not-allowed",
  )}
/>
