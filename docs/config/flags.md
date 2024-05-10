# run args

```
./scm -h

Usage of scm:
  -http-host string
        http host to bind to
  -http-port int
        http listen port
  -local-cfg-file string
        local config file (default "/etc/scm/scm.config.yaml")
  -log-level string
        one of debug, info, warn, error, dpanic, panic, fatal (default "info")
  -log-paths string
        a list of comma-separated log file path. stdout means the console stdout (default "stdout,scm.log")
  -metric-push-gateway-addrs string
        a list of comma-separated prometheus push gatway address
  -push-interval int
        push interval in seconds (default 10)
  -v    show build version and quit

```