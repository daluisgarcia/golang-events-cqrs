# Golang CQRS and Event Sourcing example
A golang feeds management project with distribuited apps using Docker. This project is composed by 3 apps:
- **Command API**: This app is responsible for creating feeds.
- **Query API**: This app is responsible for listing and searching feeds. For searching this app use ElasticSearch.
- **Pusher**: This app is responsible for sending notifications when a new feed is created. As event bus this project uses NATS Streaming.

Every app and service is built over a Docker image defined in the Docker-compose file. This project uses a Makefile to build and run the project.

## Project architecture
The project architecture is based on the following diagram:

![Project architecture](arquitecture.jpg "Project architecture")

The project uses NGINX as a reverse proxy to redirect the requests to the right app. The NGINX configuration is defined in the ```./nginx/nginx.conf``` file.

## Build and run the project
The right way to build and run this project is using Docker. To build and run the project you need to run the following commands in the root of the project:
    
    docker compose up

After executing the previous command you will have the services running in ```http://localhost:8080```. The endpoints available are:
- **List feeds**: ```GET http://localhost:8080/feeds```
- **Search feeds**: ```GET http://localhost:8080/feeds/search?query=```
- **Create feed**: ```POST http://localhost:8080/feeds```
