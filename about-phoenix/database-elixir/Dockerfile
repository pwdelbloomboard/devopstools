# Use an official Elixir runtime as the base image
FROM elixir:1.12

# Set the working directory inside the container
WORKDIR /app

# Install hex package manager and rebar
# RUN mix local.hex --force && \
#    mix local.rebar --force

# Copy the application source code (if available)
COPY ./volumebindmount/ /home/app/
# Enable start script
RUN chmod +x /home/app/start

# Set the command to run when the container starts
CMD ["/home/app/start"]
