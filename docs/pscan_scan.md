## pscan scan

scan hosts ports

```
pscan scan [flags]
```

### Options

```
  -f, --filter string    filter open or closed ports
  -h, --help             help for scan
  -p, --ports string     ports to scan within hosts (default "22,80,443")
  -r, --range string     port range to scan within hosts
  -t, --timeout string   timeout in milliseconds
```

### Options inherited from parent commands

```
      --config string       config file (default is $HOME/.pscan.yaml)
  -F, --hosts-file string   file to store hosts (default "pscan.hosts")
```
