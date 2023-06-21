FROM alpine:latest

# Install bash
RUN apk add --no-cache bash

# Install bash and PostgreSQL client
RUN apk add --no-cache bash postgresql-client

# Set the working directory in the container to /app
WORKDIR /app

# Copy the resource into the container
COPY . .

# Make the batch script executable
RUN chmod +x *.sh

# Command to run the batch script
CMD ["./batch_script.sh"]
