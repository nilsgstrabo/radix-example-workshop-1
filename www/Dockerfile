#
# -- Base node image with app
#
FROM node:10-alpine AS base
WORKDIR /usr/src/app
COPY package.json package-lock.json app.js ./
RUN mkdir -p ./bin ./routes
COPY bin bin
COPY public public
COPY views views
COPY routes routes

#
# -- Dependencies
#
FROM base as dependencies
WORKDIR /usr/src/app
RUN npm install --only=production
RUN cp -R node_modules node_modules_production
RUN npm install

#
# Potentially running test, linting, security check +++
#
FROM dependencies as test
WORKDIR /usr/src/app
#COPY test test
#RUN ["npm","test"]
COPY .eslintrc.json .eslintignore ./
RUN npm run lint
RUN npm audit

#
# Release image
#
FROM base as release
RUN addgroup -S -g 1000 radix-operator
RUN adduser -S -u 1000 -G radix-operator radix-operator
WORKDIR /usr/src/app
COPY --from=dependencies /usr/src/app/node_modules_production ./node_modules
EXPOSE 3001
ENV NODE_ENV=production
USER 1000:1000
ENTRYPOINT [ "npm", "start"]