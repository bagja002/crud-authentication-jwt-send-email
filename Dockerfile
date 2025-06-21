FROM golang:1.21-alpine

WORKDIR /app

# Copy semua source code ke dalam container
COPY . .



# Download dependensi
RUN go mod download

# Expose port (jika kamu dengarkan di 0.0.0.0:8973)
EXPOSE 9000

# Jalankan langsung dengan go run
CMD ["go", "run", "main.go"]
