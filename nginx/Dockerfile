# syntax=docker/dockerfile:1

FROM nginx:alpine
COPY nginx.conf /etc/nginx/nginx.conf
COPY /etc/letsencrypt/live/trisciuzzi.dev/cert.pem /etc/nginx/ssl/cert.pem
COPY /etc/letsencrypt/live/trisciuzzi.dev/privkey.pem /etc/nginx/ssl/key.pem
