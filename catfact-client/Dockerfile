FROM node:11.6.0-alpine AS builder
COPY . ./catfact-client
WORKDIR /catfact-client
RUN npm i
RUN $(npm bin)/ng build --prod

FROM nginx:1.15.8-alpine
COPY --from=builder /catfact-client/dist /usr/share/nginx/html