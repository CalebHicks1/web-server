# Caleb's Webserver
This is the start of my personal webserver. 
Uses:
- docker compose
    - nginx
    - webapp

# Setup
## Install
1. Install docker: 
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-18-04
2. Install docker-compose
## Run
1. Run nginx proxy: 
```
docker-compose -f nginx-proxy-compose.yaml up -d
```
2. Run webapp:
```
docker-compose up -d
```