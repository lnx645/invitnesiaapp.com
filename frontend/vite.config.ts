import { svelte } from "@sveltejs/vite-plugin-svelte";
import { router } from "sv-router/vite-plugin";
import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  plugins: [tailwindcss(), svelte(), router()],
  build: {
    rolldownOptions: {
      output: {
        entryFileNames: "dist/app-dist.js",
        chunkFileNames: "chunks/[name]-dist.js",
        assetFileNames: (assetInfo) => {
          if (assetInfo.name && assetInfo.name.endsWith(".css")) {
            return "[name]-dist.[ext]";
          }
          return "[name].[ext]";
        },
      },
    },
  },
});
