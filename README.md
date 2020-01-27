# didactic-succotash

## Usage
After the default target of the Makefile was triggered, one instance of the main app service becomes available and 
accepts requests.
The easiest way to check the functionality of the service is to send an HTTP request with curl:
```cgo
curl -H "Source-Type: game" -H "Content-Type: application/json" \
-d '{"state": "win", "amount": 10.15, "transaction_id": "some id"}' -X POST "localhost:32838/update"
``` 
