## pscan completion

Generate bash completion for your command

### Synopsis

To load your completions run
	source <(pScan completion)
To load completions automatically on login, add this line to you .bashrc file: $ ~/.bashrc
source <(pScan completion)


```
pscan completion [flags]
```

### Options

```
  -h, --help   help for completion
```

### Options inherited from parent commands

```
      --config string       config file (default is $HOME/.pscan.yaml)
  -f, --hosts-file string   file to store hosts (default "pscan.hosts")
```

### SEE ALSO

* [pscan](../README.md)	 - Fast TCP port scanner

