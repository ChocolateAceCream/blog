FROM nginx:alpine
ARG build_env=staging

RUN cd /var && mkdir -p www
# COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY ./dist /var/www/dist
