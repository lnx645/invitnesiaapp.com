<script lang="ts">
  import toast from "svelte-french-toast";
 
  function onsubmit(e : SubmitEvent) {
      e.preventDefault();
      toast.success("Login failed, please try again later");
  }
</script>

<div class="card rounded-2xl shadow-2xl border-zinc-300 max-w-sm mx-auto mt-24">
  <div class="p-8 flex flex-col gap-4">
    <div class="text-2xl border-b pb-3 border-b-zinc-200">
      <h2>Login</h2>
    </div>
    <form onsubmit={onsubmit} class="flex flex-col gap-y-3">
      <div class="flex-1 w-full flex flex-col">
        <label class="text-xs inline-block mb-1" for="email">Email</label>
        <input
          type="text"
          placeholder="Masukan email anda"
          class="w-full border hover:border-zinc-300 text-sm transition-all focus-within:border-blue-800 active:border-blue-800 focus:border-blue-800 border-zinc-300 bg-zinc-50 outline-0 rounded-md px-2 py-2"
        />
      </div>
      <div class="flex-1 w-full flex flex-col">
        <label class="text-xs inline-block mb-1" for="password">Password</label>
        <input
          type="password"
          placeholder="Ketikan kata sandi anda"
          class="w-full border hover:border-zinc-300 text-sm transition-all focus-within:border-blue-800 active:border-blue-800 focus:border-blue-800 border-zinc-300 bg-zinc-50 outline-0 rounded-md px-2 py-2"
        />
      </div>
      <div class="flex-1 w-full flex flex-col">
        <button
          class="w-full bg-blue-800 transition-all hover:bg-blue-900 text-white cursor-pointer shadow-2xl border-none rounded-md text-sm border p-2"
          >Login</button
        >
      </div>
    </form>
  </div>
</div>
