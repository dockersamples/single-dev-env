let server = Bun.serve({
    port: 3000,
    async fetch(req) {
        const path = "New Tuning.mp3";
        const file = Bun.file(path);
        return new Response(file);
      },
})

console.log(`Server is running on port: ${server.port}`);