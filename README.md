# Ticket

Go script that checks the availability of tickets at a specific theatre and gives a push notification when tickets are available.

## Usage

```bash
go run main.go --url="bookmyshow_url" --theatre="theatre_name"
```

## Running as background service (Linux)

1. Install the binary

```bash
go get -u github.com/saurabhmittal16/ticket
```

2. Run binary and detach from current session

```bash
nohup ticket --url="bookmyshow_url" --theatre="theatre_name" &
```

The service runs in background and logs are appended to `~/nohup.out`. The process quits when tickets are available but it can be manually stopped by following the mentioned procedure

## Stopping the process

1. Get the pid of the process

```bash
ps -e | grep ticket
```

2. Kill the process with obtained pid

```bash
kill 1234
```
