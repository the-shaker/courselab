FROM oven/bun:alpine AS build
WORKDIR /app
COPY app/package.json app/bun.lockb* ./
RUN bun install --frozen-lockfile
COPY app .
RUN bun run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
