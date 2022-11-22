grpcurl --plaintext -d '{"name": "reid", "policy": "balanced", "devices": "cpu,gpu", "endpoints": "10.0.0.1:55555,10.0.0.2:55555"}' localhost:55554 localhost:55554/Push
