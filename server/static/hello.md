# 2023-5-5 multiple nginx container deployment
<a> #nginx </a><a> #docker </a>

## Background
In case you only have bought one domain <a>www.easydelivery.ltd</a>, but you want to host multiple web apps, then only choice left here is using sub-domains, such as <a>photo.easydelivery.ltd</a> or <a>blog.easydelivery.ltd</a>

Normally you can deploy multiple web apps using one nginx docker containers, however, one issue with that approach is that deployment of each app will affect others, since you will need to re-create nginx containers for every deployment. So, one solution for this problem would be utilizing multiple nginx containers to deploy each web apps separately.

<img src="/assets/images/2023.5.5/1.png" alt="图片alt" title="图片title">

As shown in the network topology, we use a main nginx container to forward requests to each sub-domains, which has fixed configs for its single task --- just forwarding requests.

Then in each sub-nginx container, we use different nginx config to suit each apps' different demands, and these sub containers can be altered without affecting others.

## Configuration
1. main nginx config

```nginx
# load balance
upstream blog_server {
  server blog.easydelivery.ltd:3040;
}

upstream blog_swagger {
  server blog.easydelivery.ltd:3030;
}

# log config
log_format  blog_log  'remote_addr: $remote_addr '
                      'remote_user: $remote_user '
                      'X-Forwarded-Proto $scheme '
                      'Host $http_host '
                      'http_x_forwarded_for: $http_x_forwarded_for';

server {
  listen 80;
  listen [::]:80;
  root /var/www/dist-1.12/;
  index index.html;
  server_name         www.easydelivery.ltd;

  #将 http 重定向 https
  return 301 https://$server_name$request_uri;
}

# ssl verification related config
ssl_session_timeout  5m;    #缓存有效期
ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;    #encrypt algorithm
ssl_protocols TLSv1 TLSv1.1 TLSv1.2;    #安全链接可选的加密协议
ssl_prefer_server_ciphers on;   #使用服务器端的首选算法

#https
server {
  listen 443 ssl;
  server_name         www.easydelivery.ltd;
  root /var/www/dist-1.12/;
  index index.html;

  # for https connection, need to provide these certificate infos
  ssl_certificate     /etc/nginx/cert/www.easydelivery.ltd.pem;
  ssl_certificate_key /etc/nginx/cert/www.easydelivery.ltd.key;
}

server {
  listen 443 ssl;
  server_name         photo.easydelivery.ltd;
  root /var/www/photo-art/;
  index index.html;

  ssl_certificate     /etc/nginx/cert/photo.easydelivery.ltd.pem;
  ssl_certificate_key /etc/nginx/cert/photo.easydelivery.ltd.key;
}

server {
  listen 443 ssl;
  server_name         blog.easydelivery.ltd;
  ssl_certificate     /etc/nginx/cert/blog.easydelivery.ltd.pem;
  ssl_certificate_key /etc/nginx/cert/blog.easydelivery.ltd.key;

  # $http_host is blog.easydelivery.ltd in this case
  proxy_set_header Host $http_host;

  # customized header should always start with x

  # $remote_addr is the client ip where request was sent from
  proxy_set_header  X-Real-IP $remote_addr;

  # if the request go through multiple proxy, we use $proxy_add_x_forwarded_for to track each proxy ip.
  # so we can obtain the whole chain of proxy path, start from origin client ip from left, followed by each proxy ip request passed through
  proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

  # $scheme is https in this case.
  # this header only useful if we want to identify the origin request is http or https
  proxy_set_header X-Forwarded-Proto $scheme;

  # output access_log to logs/blog-web.log, with log format blog_log configured above. Remember to map the log file location if you are using docker containers.
  access_log  logs/blog-web.log blog_log;

  location / {
    proxy_pass https://blog_server; # 设置代理服务器的协议和地址
  }

  location /swagger/ {
    proxy_pass https://blog_swagger;
  }
}
```

We deployed three web apps in this main nginx, one listen to 80 (http) and two other listen to 443 (https). For https web app, we need to specify the ssl_certificates, so don't forget to map the files location when creating nginx images.

```yml
# docker-compose.yml for main nginx
version: '3'
services:
  nginx:
    image: nginx:alpine
    container_name: nginx
    ports:
      - "80:80"
      - "443:443"
    restart: always
    privileged: true
    volumes:
      - /root/docker/nginx/nginx.conf:/etc/nginx/conf.d/default.conf

      # web app dist
      - /root/docker/nginx/www/:/var/www/

      #ssl certificate
      - /root/docker/nginx/cert/:/etc/nginx/cert/

      #logs
      - /root/docker/nginx/logs/:/etc/nginx/logs/

```

now you can start your main nginx containers using
```bash
docker-compose up --force-recreate -d
```

As you should noticed in main nginx config, we forwarded the request for the third app blog, to the sub nginx server running on <a>blog.easydelivery.ltd:3040</a>. So next we need to config the sub nginx server which actually hosting the blog web app.

2. sub nginx config
```nginx
log_format  blog_log  'remote_addr: $remote_addr '
                      'remote_user: $remote_user '
                      'X-Forwarded-Proto $scheme '
                      'Host $http_host '
                      'X-Real-IP: $http_x_real_ip '
                      'Cookie $http_Cookie '
                      'proxy_add_x_forwarded_for: $proxy_add_x_forwarded_for';
#https
server {
  listen 3040 ssl;
  server_name         blog.easydelivery.ltd;
  root /var/www/dist/;
  index index.html;

  gzip on;
  gzip_min_length 1024;
  gzip_comp_level 4;
  gzip_http_version 1.1;
  gzip_types text/plain text/css application/x-javascript application/javascript application/xml application/json;

  # still need to specify certificate since we are listening to ssl port
  ssl_certificate     /etc/nginx/cert/blog.easydelivery.ltd.pem;
  ssl_certificate_key /etc/nginx/cert/blog.easydelivery.ltd.key;



  location /blog/api {
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    # rewrite ^/api/(.*)$ /$1 break;  #重写

    # where actual backend server listening to
    proxy_pass http://blog.easydelivery.ltd:3030/api; # 设置代理服务器的协议和地址

    # proxy_cookie_path  / /blog/api;
    # proxy_cookie_domain blog.easydelivery.ltd xxx.xx.xx.xx;
    access_log  logs/blog-web.log blog_log;

  }
  location / {
    try_files $uri $uri/ /index.html;
  }

  location /swagger/ {
    proxy_pass http://blog.easydelivery.ltd:3030;
  }
}
```

This sub nginx listen to <a>blog.easydelivery.ltd:3040</a>, and forward all api request (/blog/api) to <a>blog.easydelivery.ltd:3030</a>, where actual backend server running on.

```yml
# docker-compose.yml for sub nginx
version: '3'
services:
  nginx:
    image: nuodi/blog-web:latest
    container_name: blog_web
    ports:
      - "3040:3040"
    restart: always
    privileged: true
    volumes:
      - /root/docker/blog/web/nginx.conf:/etc/nginx/conf.d/default.conf
      - /root/docker/blog/web/dist/:/var/www/blog/
      #ssl certificate
      - /root/docker/blog/web/cert/:/etc/nginx/cert/
      - /root/docker/blog/web/logs/:/etc/nginx/logs/
```

## Summary
Now every time frontend web app re-packed, we only need to recreate the sub-nginx containers, which will not affect any other web apps since main nginx container remain the same.