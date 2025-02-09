FROM node:18-alpine AS dependencies
WORKDIR /frontend
COPY /frontend/react-intro/package-lock.json .
COPY /frontend/react-intro/package.json .
RUN npm install


FROM dependencies AS build
WORKDIR /frontend
COPY frontend/react-intro/. .
RUN npm run build


FROM nginx:stable-alpine
COPY --from=build frontend/build /usr/share/nginx/html
COPY frontend/react-intro/nginx.conf /etc/nginx/conf.d/default.conf
CMD ["nginx", "-g", "daemon off;"]
