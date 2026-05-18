FROM oven/bun:alpine AS build
WORKDIR /app
COPY app/package.json app/bun.lockb* ./
RUN bun install --frozen-lockfile
COPY app .
ARG VITE_API_URL
ENV VITE_API_URL=$VITE_API_URL
RUN bun run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
