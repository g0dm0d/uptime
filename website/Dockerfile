### STAGE 1: Build ###
FROM node:20.8-alpine3.18@sha256:80cc2c520781c1a4681e59df6791d81a432a6cb3da0c385d8d62e9f16acf8e5f as builder
WORKDIR /app
COPY ./website/package.json ./
COPY ./website/package-lock.json ./
COPY .env .
RUN npm install
COPY ./website .
RUN npm run build


### STAGE 2: Run ###
FROM nginx:1.25.2-alpine3.18@sha256:34b58b4f5c6d133d97298cbaae140283dc325ff1aeffb28176f63078baeffd14
COPY --from=builder /app/build /var/www/default