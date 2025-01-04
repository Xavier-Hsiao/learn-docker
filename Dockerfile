FROM debian:stable-slim

# COPY source destination
COPY learn-docker /bin/learn-docker

# Run the server process when the container starts running
CMD [ "/bin/learn-docker" ]