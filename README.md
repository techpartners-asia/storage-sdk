# MinIO Service Implementation

A Go-based service implementation for interacting with MinIO object storage. This service provides a simple interface for managing files in MinIO buckets, including uploading, deleting files, and creating buckets.

## Features

- File upload to MinIO buckets
- File deletion from MinIO buckets
- Bucket creation with public access policy
- Automatic bucket creation if it doesn't exist during file upload
- SSL support for secure connections

## Prerequisites

- Go 1.16 or higher
- MinIO server instance
- Access credentials for MinIO (access key and secret key)

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd storage-sdk
```

2. Install dependencies:

```bash
go mod download
```

## Usage

### Initialization

```go
service, err := NewMinioService(
    "minio.example.com",  // MinIO server host
    "your-access-key",    // MinIO access key
    "your-secret-key",    // MinIO secret key
    true,                 // Use SSL
)
if err != nil {
    log.Fatal(err)
}
```

### Uploading a File

```go
// Assuming you have a multipart.FileHeader from an HTTP request
filename, err := service.UploadFile("my-bucket", fileHeader)
if err != nil {
    log.Fatal(err)
}
```

### Deleting a File

```go
err := service.DeleteFile("my-bucket", "filename.txt")
if err != nil {
    log.Fatal(err)
}
```

### Creating a Bucket

```go
err := service.CreateBucket("new-bucket")
if err != nil {
    log.Fatal(err)
}
```

## API Reference

### MinioService Interface

```go
type MinioService interface {
    UploadFile(bucketName string, file *multipart.FileHeader) (string, error)
    DeleteFile(bucketName string, fileName string) error
    CreateBucket(bucketName string) error
}
```

## License

This project is licensed under the terms specified in the LICENSE file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
