# Getting Started

This chapter will guide you on how to quickly deploy and use A.I.G.

## One-Click Installation

### Deploying with Docker
```bash
docker-compose -f docker-compose.images.yml up -d
```

After the installation is complete, you can access the A.I.G Web UI by visiting `http://localhost:8088` in your browser.

## Configure Model KEY

A.I.G's `MCP Scan` and `Large Model Security Health Check` features require the use of a large model API. If you need to use these two functions, you can first configure the large model API KEY.

![image-20250814173229996](./assets/image-20250814173229996.png)

Configure the required Model Name, API Key, and Base URL, then click Save.

![image-20250813113550192](./assets/image-20250813113550192.png)


## Frequently Asked Questions

1. **Port Conflict**
   ```bash
   # Modify the webserver port mapping
   ports:
     - "8080:8088"  # Use port 8080
   ```

2. **Permission Issues**
   ```bash
   # Check data directory permissions
   sudo chown -R $USER:$USER ./data
   ```

3. **Service Startup Failure**
   ```bash
   # View detailed logs
   docker-compose logs webserver
   docker-compose logs agent
   ```

4. **Stopping the Service**
    ```bash
    # Stop the service
    docker-compose down
    
    # Stop the service and remove data volumes (use with caution)
    docker-compose down -v
    ```


## Updates and Upgrades

```bash
# Rebuild and start
docker-compose -f docker-compose.images.yml up -d --build
# Clean up old images
docker image prune -f
```