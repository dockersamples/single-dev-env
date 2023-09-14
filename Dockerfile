FROM oven/bun

ADD . .
RUN bun install
EXPOSE 3000

CMD bun run index.ts