# lcvsimmu

command launch

benchmark {number of insertions} {apikey} {lc address} {lc port} {immudb address} {immudb port} 
```bash
./benchmark 1000 athlndeyjnyfqcnvoexjwlpjtltgcxtnqjyu localhost 3324 localhost 3322
```
to skip immudb benchmark

```bash
./benchmark 1000 athlndeyjnyfqcnvoexjwlpjtltgcxtnqjyu localhost 3324 
```

Enable debug:
```bash
export LCVSIMMU_DEBUG=debug
```