# CI/CD Setup Guide

This project includes a comprehensive CI/CD pipeline using GitHub Actions.

## Workflows

### 1. Pull Request Check (`pr-check.yml`)
- Runs on every pull request to `main` and `develop` branches
- Performs testing, linting, and build verification
- Ensures code quality before merging

### 2. CI/CD Pipeline (`ci-cd.yml`)
- Runs on pushes to `main` and `develop` branches
- Includes testing, building, security scanning, and deployment stages

## Pipeline Stages

### Test Stage
- Runs unit tests with MySQL service
- Performs code linting with `golint`
- Checks code formatting with `go fmt`
- Runs `go vet` for static analysis

### Build Stage
- Builds Docker image using multi-stage build
- Pushes to Docker Hub with tags
- Only runs on pushes to `main` branch

### Security Scan Stage
- Runs Trivy vulnerability scanner
- Uploads results to GitHub Security tab
- Only runs on pushes to `main` branch

### Deployment Stages
- **Staging**: Deploys on pushes to `develop` branch
- **Production**: Deploys on pushes to `main` branch

## Required Secrets

Add these secrets to your GitHub repository (Settings > Secrets and variables > Actions):

### Docker Hub Credentials
- `DOCKER_USERNAME`: Your Docker Hub username
- `DOCKER_PASSWORD`: Your Docker Hub password or access token

### Optional: Environment-specific secrets
- `STAGING_DB_HOST`: Staging database host
- `STAGING_DB_PASSWORD`: Staging database password
- `PRODUCTION_DB_HOST`: Production database host
- `PRODUCTION_DB_PASSWORD`: Production database password

## Environment Setup

### 1. Create Environments (Optional)
In your GitHub repository, go to Settings > Environments and create:
- `staging`
- `production`

### 2. Configure Branch Protection (Recommended)
Go to Settings > Branches and add branch protection rules for `main`:
- Require status checks to pass before merging
- Require branches to be up to date before merging
- Require pull request reviews before merging

## Local Development

### Running Tests Locally
```bash
# Start MySQL for testing
docker run -d --name mysql-test \
  -e MYSQL_ROOT_PASSWORD=admin123 \
  -e MYSQL_DATABASE=multi_finance_test \
  -p 3306:3306 mysql:8.0

# Run tests
go test -v ./...

# Run linting
go install golang.org/x/lint/golint@latest
golint ./...

# Check formatting
go fmt ./...
```

### Building Locally
```bash
# Build application
go build -o main .

# Build Docker image
docker build -t multifinance-go .
```

## Customization

### Adding More Tests
1. Create test files with `_test.go` suffix
2. Tests will automatically run in the CI pipeline

### Adding Code Quality Checks
1. Add new steps in the workflow files
2. Common additions:
   - `go vet` (already included)
   - `gosec` for security scanning
   - `staticcheck` for static analysis

### Customizing Deployment
1. Update the deployment steps in `ci-cd.yml`
2. Add your specific deployment commands
3. Examples:
   - Kubernetes: `kubectl apply -f k8s/`
   - Docker Compose: `docker-compose up -d`
   - AWS ECS: `aws ecs update-service`

## Troubleshooting

### Common Issues

1. **Tests failing in CI but passing locally**
   - Check environment variables
   - Ensure MySQL service is running
   - Verify database connection settings

2. **Docker build failing**
   - Check if `go.sum` exists (should be auto-generated)
   - Verify Dockerfile syntax
   - Check for missing files in `.dockerignore`

3. **Security scan failing**
   - Review Trivy results in GitHub Security tab
   - Update dependencies if vulnerabilities found
   - Consider using `--severity HIGH,CRITICAL` for stricter scanning

### Debugging Workflows
1. Check workflow runs in Actions tab
2. Review logs for specific step failures
3. Use `echo` statements for debugging in workflow files
4. Enable debug logging by setting `ACTIONS_STEP_DEBUG` secret to `true`

## Best Practices

1. **Branch Strategy**
   - Use feature branches for development
   - Merge to `develop` for testing
   - Merge to `main` for production releases

2. **Commit Messages**
   - Use conventional commits: `feat:`, `fix:`, `docs:`, etc.
   - Keep commits atomic and focused

3. **Security**
   - Never commit secrets to the repository
   - Use GitHub Secrets for sensitive data
   - Regularly update dependencies

4. **Monitoring**
   - Set up notifications for workflow failures
   - Monitor deployment health
   - Track performance metrics 