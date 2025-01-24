# Ghaymah CLI (غيمة)

Command line interface for the Ghaymah Cloud Platform.

## Features

- 🚀 Easy application deployment
  - Deploy using config file or Docker image
  - Customize resources (CPU, Memory, Storage)
  - Set environment variables
- 📊 Application status monitoring
  - Real-time deployment status
  - Resource usage metrics
  - Health checks
- 📝 Real-time log viewing
  - Follow logs in real-time
  - Filter logs by time
  - Customize output format
- ⚙️ Resource and configuration management
  - YAML configuration
  - Environment variables
  - Resource limits
- 🔐 Docker Registry Management
  - Add private registry credentials
  - List registered registries
  - Remove registry access
- 🎫 API Token Management
  - Generate secure API tokens
  - List active tokens
  - Revoke tokens when needed

## Installation

```bash
go install github.com/your-username/ghaymah-cli@latest
```

Or download the binary from the releases page.

## Configuration

### 1. Environment Variables

Required environment variables:
```bash
# The URL of the Ghaymah Cloud API
export GHAYMAH_API_URL="https://api.ghaymah.cloud"

# Your Ghaymah Cloud API token
export GHAYMAH_API_TOKEN="your-token-here"
```

### 2. Configuration File (Optional)

Create a `config.yaml` file for deployment configuration:
```yaml
# Application name
appName: "my-app"

# Path to Dockerfile (if building from source)
dockerfilePath: "./Dockerfile"

# Docker Registry configuration (optional)
registry:
  registryUrl: "docker.io"
  username: "your-username"
  password: "your-password"

# Deployment region
region: "us-east-1"

# Environment variables for your application
envVars:
  NODE_ENV: "production"
  PORT: "8080"
  DB_URL: "mongodb://localhost:27017"

# Resource requirements
resources:
  cpu: "1"       # Number of CPU cores
  memory: "512M" # Memory limit
  storage: "1G"  # Storage limit
```

## Usage

### Deploy Command

Deploy using a configuration file:
```bash
# Deploy using default config file (ghaymah.yaml)
ghaymah deploy

# Deploy using custom config file
ghaymah deploy -c myconfig.yaml
```

Deploy using Docker image:
```bash
# Deploy using image only (name will be extracted from image)
ghaymah deploy --image username/app:tag

# Deploy using image with custom name
ghaymah deploy --image username/app:tag --name my-custom-name
```

### Registry Commands

Manage Docker registry credentials:
```bash
# Add new registry credentials
ghaymah registry add --url docker.io --username user123 --password pass123

# List all registered registries
ghaymah registry list

# Remove registry credentials
ghaymah registry remove --url docker.io
```

### Token Commands

Manage API tokens:
```bash
# Generate new API token (default expiry: 30 days)
ghaymah token generate

# Generate token with custom expiry
ghaymah token generate --expiry-days 60

# List all active tokens
ghaymah token list

# Revoke a token
ghaymah token revoke --token <token-value>
```

### Status Command

Check application status:
```bash
# Basic status check
ghaymah status --name my-app

# Get detailed status including metrics
ghaymah status --name my-app --detailed
```

### Logs Command

View application logs:
```bash
# View recent logs (default: last 100 lines)
ghaymah logs --name my-app

# Follow logs in real-time
ghaymah logs --name my-app --follow

# View specific number of lines
ghaymah logs --name my-app --tail 50

# View logs since specific time
ghaymah logs --name my-app --since 2024-01-23T00:00:00Z

# Combine options
ghaymah logs --name my-app --follow --tail 50
```

## Development

### Mock API for Testing

For development and testing purposes, you can use the included Mock API server. This allows you to test the CLI functionality without connecting to a real Ghaymah Cloud API.

#### 1. Start the Mock API Server

```bash
cd mock-api
go run main.go
```

The server will start on `http://localhost:8080` and display the test token to use.

#### 2. Configure Environment Variables

In a new terminal, set the environment variables to use the Mock API:
```bash
export GHAYMAH_API_URL="http://localhost:8080"
export GHAYMAH_API_TOKEN="test-token-123"
```

#### 3. Test CLI Commands

Now you can test all CLI commands:
```bash
# Deploy an application
./ghaymah-cli deploy --image nginx:latest --name test-app

# Check application status
./ghaymah-cli status --name test-app

# View application logs
./ghaymah-cli logs --name test-app
```

The Mock API provides simulated responses for:
- Application deployment
- Status checking with resource metrics
- Log viewing with timestamp-based filtering

#### Mock API Endpoints

The Mock API implements these endpoints:
- `POST /apps`: Deploy applications
- `GET /apps/status`: Get application status
- `GET /apps/logs`: Get application logs

All endpoints require the `Authorization` header with the test token.

## Common Flags

These flags are available for most commands:

- `-h, --help`: Show help for any command
- `--name`: Specify application name
- `-c, --config`: Path to configuration file
- `-v, --verbose`: Enable verbose output

## Examples

1. Deploy a Node.js application:
```bash
ghaymah deploy --image node-app:1.0 --name my-node-app
```

2. Monitor application status:
```bash
ghaymah status --name my-node-app
```

3. View and follow logs:
```bash
ghaymah logs --name my-node-app --follow --tail 50
```

## Error Handling

Common error messages and solutions:

- `Error: environment variable GHAYMAH_API_URL is not set`
  - Solution: Set the GHAYMAH_API_URL environment variable
- `Error: environment variable GHAYMAH_API_TOKEN is not set`
  - Solution: Set the GHAYMAH_API_TOKEN environment variable
- `Error: config file not found`
  - Solution: Create a config.yaml file or specify path with -c flag

## Contributing

We welcome your contributions! Please follow these steps:

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/amazing-feature`
3. Commit your changes: `git commit -m 'Add amazing feature'`
4. Push to the branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
