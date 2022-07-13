# base image
FROM node:16.14.2-alpine
RUN apk add --update --no-cache gcc g++
WORKDIR /app

COPY ./frontend /app
RUN yarn install --network-timeout 1000000000 --frozen-lockfile --ignore-optional --non-interactive --silent

EXPOSE 4200 4200
ENTRYPOINT ["yarn", "start"]
