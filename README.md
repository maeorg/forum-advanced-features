# 🔥 Forum image-upload 🔥

##### 📖 _This repository contains the Forum image-upload project_

## 📜 Auditing directions 

- [Forum image-upload auditing directions](https://github.com/01-edu/public/tree/master/subjects/forum/image-upload/audit)

For testing you can use these accounts and/or create your own:

- Username: WinnieThePooh  
Password: WinnieThePooh

- Username: Eeyore  
Password: EeyoreEeyore

- Username: Tigger  
Password: TiggerTigger

- Username: Piglet  
Password: PigletPiglet

# 🥏 Running the Dockerized Go Application #
This guide will walk you through running a Go application in a Docker container. The provided Dockerfile uses the official Go image as a base and demonstrates a typical workflow for containerizing a Go application.

## Prerequisites ##
Before you begin, make sure you have the following installed on your system:  
- [Docker](https://www.docker.com/get-started/)  

## 1. Running the app in Docker with a script ##
This is the easiest way to run the application. Just clone the repo and run this script. It will automatically build the Docker image and run it.
- **./build-and-run.sh**

## 2. Docker files cleanup with a script ##
This is the easiest way to clean up. It will automatically stop the container and delete the Docker files.
- **./cleanup.sh**
 
###
## 🛠 Running the program in Docker without scripts 📰 ##

### Building the Docker Image ###
1. Clone or download the repository containing the Go application and the Dockerfile.
2. Open a terminal and navigate to the directory containing the Dockerfile and the Go application code.
3. Run the following command to build a Docker image tagged as *docker-forum*. This command instructs Docker to build an image using the Dockerfile and tag it with the name *docker-forum*. 
- **docker build --tag docker-forum .**

### Running the Docker Container ###
Once the Docker image is built, you can run a Docker container from it. This command starts a Docker container from the *docker-forum* image and maps port 8080 from the container to the same port on your host machine. Your Go application will be accessible at http://localhost:8080 in your web browser or via a web API request.
- **docker run -p 8080:8080 docker-forum**

### Stopping the Container ###
To stop the running container, open a new terminal window and use the following command. Replace <container_id> with the actual container ID or the name of the container you want to stop. You can find the container ID by running docker ps or docker ps -a to list all running and stopped containers.
- **docker stop <container_id>**

### Cleaning Up ###
After stopping the container, you can remove it by running:
- **docker rm <container_id>**

Additionally, you can remove the Docker image to save disk space:
- **docker rmi docker-forum**

## 🛠 Running the program without Docker 🏃

1. Install any necessary dependencies. Make sure you have Go installed.

2. Clone this repository to your local machine:
```bash
 git clone https://01.kood.tech/git/kmaeorg/forum.git
```

3. Navigate to the project directory:
```bash
 cd forum
```

4. Run the Go application: 
```bash
 go run .
```

5. Navigate to the server address in your browser:
```bash
  http://localhost:8080/
```

## 🌟 Authors

- **[Katrina Mäeorg](https://01.kood.tech/git/kmaeorg)**
