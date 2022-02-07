# The base go-image
FROM golang:1.17.6
 
# Create a directory for the app
RUN mkdir /app
 
# Copy all files from the current directory to the app directory
COPY . /app
 
# Set working directory
WORKDIR /app
 
# Run the server executable
CMD [ "go", "run", "./cmd/main.go" ]