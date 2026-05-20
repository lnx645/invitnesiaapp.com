<script lang="ts">
  import { Router } from "sv-router";
  import "sv-router/generated";
  import { Toaster } from "svelte-french-toast";
  import BaseLayout from "./layouts/base-layout.svelte";
</script>

<Toaster />
<BaseLayout>
  <Router base="/app" />
</BaseLayout>
