# Use the official Go image with version 1.19
FROM golang:1.19.6-bullseye

# Set the working directory inside the container
# WORKDIR /app

# Copy the Go module files and download dependencies

RUN apt update -y
RUN apt upgrade -y
RUN apt install -y curl apt-transport-https -y
RUN apt install rdiff-backup -y
RUN curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor -o /usr/share/keyrings/postgresql-keyring.gpg
RUN echo "deb [signed-by=/usr/share/keyrings/postgresql-keyring.gpg] http://apt.postgresql.org/pub/repos/apt/ bullseye-pgdg main" | tee /etc/apt/sources.list.d/postgresql.list
RUN apt update -y
RUN apt install postgresql-client-14 -y
# RUN mkdir -p /root/.ssh
# COPY ["/home/renato/.ssh/id_ed25519", "/root/id_ed25519"]
# RUN chmod 700 -R /root/.ssh
# RUN chmod +x /root/.ssh
# RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
# COPY go.mod go.sum ./
# RUN go mod download


# Copy the rest of the application code
# COPY . .

# Build the application
# RUN go build -o ./dist/server-backup

# Set the container command to run the compiled binary
# ENTRYPOINT ["./dist/server-backup"]
