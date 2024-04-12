# Use the official PostgreSQL image
FROM postgres:latest

# Set environment variables
ENV POSTGRES_DB=disappoint_db
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=root

# Copy the initialization script to the container
COPY init_db.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432
