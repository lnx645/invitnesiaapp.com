<script lang="ts">
  import { p } from "sv-router/generated";
  import toast from "svelte-french-toast";
  function onsubmit(e: SubmitEvent) {
    e.preventDefault();
    toast.success("Login failed, please try again later");
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
          <input
            type="text"
            placeholder="Masukan email anda"
            class="w-full border hover:border-zinc-300 text-sm transition-all focus-within:border-blue-800 active:border-blue-800 focus:border-blue-800 border-zinc-300 bg-zinc-50 outline-0 rounded-md px-3 py-2"
          />
        </div>
        <div class="flex-1 w-full flex flex-col">
          <label
            class="text-xs text-[#3A3541DE] inline-block mb-1"
            for="password">Password</label
          >
          <input
            type="password"
            placeholder="Ketikan kata sandi anda"
            class="w-full border hover:border-zinc-300 text-sm transition-all focus-within:border-blue-800 active:border-blue-800 focus:border-blue-800 border-zinc-300 bg-zinc-50 outline-0 rounded-md px-3 py-2"
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
            class="w-full disabled:bg-zinc-400 disabled:text-[#cecece] bg-blue-800 transition-all hover:bg-blue-900 text-white cursor-pointer shadow-2xl border-none rounded-md text-sm border p-2"
            >Login</button
          >
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
