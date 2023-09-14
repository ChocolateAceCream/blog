# called by github action to deploy images on host machine

sudo docker pull nuodi/blog-server:latest
sudo docker pull nuodi/blog-web:latest
sudo docker-compose -f /root/docker/blog/server/docker-compose.yml up --force-recreate --remove-orphans -d
sudo docker-compose -f /root/docker/blog/web/docker-compose.yml up --force-recreate --remove-orphans -d

# delete all unused docker images
sudo docker image prune -f --filter "label!=nuodi-blog" --filter "dangling=true"