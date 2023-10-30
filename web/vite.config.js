import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  define: {
    "process.env": {
      NODE_ENV: JSON.stringify(process.env.NODE_ENV),
    },
  },
  plugins: [react()],
  server: {
    port: 3000,
  },
  build: {
    outDir: "build",
    emptyOutDir: true,
  },
  test: {
    globals: true,
    environment: "jsdom",
  },
});
