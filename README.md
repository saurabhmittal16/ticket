# Ticket

Go script that checks the availability of tickets at a specific theatre and gives a push notification when tickets are available.

## Usage

```bash
go run main.go --url="bookmyshow_url" --theatre="theatre_name" --duration=5
```

## Running as background service (Linux)

- Install the binary

```bash
go get -u github.com/saurabhmittal16/ticket
```

- Run binary and detach from current session

```bash
nohup ticket --url="bookmyshow_url" --theatre="theatre_name" --duration=5 &
```

The service runs in background and logs are appended to `~/nohup.out`. The process quits when tickets are available but it can be manually stopped by following the mentioned procedure

## Stopping the process

- Get the pid of the process

```bash
ps -e | grep ticket
```

- Kill the process with obtained pid

```bash
kill 1234
```
