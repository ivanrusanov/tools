# good-listener
Utility that setups HTTP server on specified port and logs all incoming requests, answering to them with 200 OK. It can be configured to 
accept reuests on specific endpoint(s).

# Usage
The easiest way to run good-listener is to use Docker image:
```bash
docker run -p 80:80 ghcr.io/ivanrusanov/tools:main -port=80 -endpoint="/hook"
```