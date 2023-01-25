FROM node:14.17 as build 

WORKDIR /react-app

ARG REACT_APP_BASE_URL

ENV REACT_APP_BASE_URL ${REACT_APP_BASE_URL}

COPY package.json .
COPY yarn.lock .

RUN yarn install

COPY . .

RUN yarn build

FROM nginx:1.19

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=build /react-app/build /usr/share/nginx/html