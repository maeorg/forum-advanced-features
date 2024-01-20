#!/bin/bash

# Stop the running container
echo -e "Stopping containers...\n"
docker stop docker-forum

# Remove the stopped container
echo -e "Removing stopped container...\n"
docker rm docker-forum

# Remove the images
echo -e "Removing images...\n"
docker rmi docker-forum
