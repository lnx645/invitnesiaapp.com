<script lang="ts">
  import { p } from "sv-router/generated";
  import Input from "../components/input.svelte";
  import toast from "svelte-french-toast";

  let data = $state<{
    email: string;
    password: string;
  }>({
    email: "",
    password: "",
  });
  let isLoading = $state(false);
  function onsubmit(e: SubmitEvent) {
    isLoading = true;
    e.preventDefault();
    setTimeout(() => {
      toast.success("Login berhasil!");
      isLoading = false;
    }, 5000);
  }
</script>

<div class="container mx-auto p-3">
  <div
    class="card rounded-2xl shadow-2xl bg-white border-zinc-300 max-w-md mx-auto mt-12"
  >
    <div class="p-8 flex flex-col gap-4">
      <div
        class="text-2xl text-[#3A3541DE] border-b pb-4 mb-1 border-b-zinc-200"
      >
        <h2>Welcome Back!</h2>
        <p class="text-sm">
          Masuk ke akun kamu untuk kelola pernikahan impianmu
        </p>
      </div>
      <form {onsubmit} class="flex flex-col gap-y-5">
        <div class="flex-1 w-full flex flex-col">
          <label class="text-xs text-[#3A3541DE] inline-block mb-1" for="email"
            >Email</label
          >
          <Input
            disabled={isLoading}
            type="email"
            bind:value={data.email}
            placeholder="Masukan email anda"
          />
        </div>
        <div class="flex-1 w-full flex flex-col">
          <label
            class="text-xs text-[#3A3541DE] inline-block mb-1"
            for="password">Password</label
          >
          <Input
            disabled={isLoading}
            type="password"
            bind:value={data.password}
            placeholder="Ketikan kata sandi anda"
          />
          <div class="text-right">
            <a
              href={p("auth/forgot-password")}
              class="text-xs text-blue-800 hover:underline">Forgot password?</a
            >
          </div>
        </div>
        <div class="flex-1 w-full flex flex-col">
          <button
            type="submit"
            disabled={isLoading}
            class="w-full disabled:bg-zinc-400 disabled:text-[#cecece] bg-blue-600 transition-all hover:bg-blue-900 text-white cursor-pointer shadow-2xl border-none rounded-md text-sm border p-2"
          >
            {isLoading ? "Loading..." : "Login"}
          </button>
        </div>
      </form>
      <div>
        <p class="text-xs text-[#3A3541DE] text-center">
          Don't have an account? <a
            href={p("auth/register")}
            class="text-blue-800 hover:underline">Register</a
          >
        </p>
      </div>
    </div>
  </div>
</div>
