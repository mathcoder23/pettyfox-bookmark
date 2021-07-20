FROM node:16-alpine3.11

WORKDIR /app
ADD ./src /app/src
ADD ./views /app/views
ADD ./test /app/test
ADD ./package.json /app/
ADD ./tsconfig.build.json /app/
ADD ./tsconfig.json /app/
ADD ./nest-cli.json /app/
RUN yarn install --registry https://registry.npm.taobao.org/
RUN yarn run build
RUN ls
EXPOSE 3000
CMD ["yarn", "start:prod"]
