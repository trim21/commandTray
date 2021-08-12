A example config, put it as `$PWD/config.yaml`

```yaml
program:
    - name: clash
      program: "clash.exe"
      args: ["-d", "cfg/clash"]

    - name: caddy
      program: caddy
      cwd: "cfg/caddy"
      args: ["run", "-watch", "-config", "Caddyfile", "-adapter", "caddyfile"]

cron:
    - spec: "12 */2 * * *"
      name: run a script as cron
      program: C:\Users\Trim21\.venv\test\Scripts\python.exe
      args: ["C:/Users/Trim21/proj/test/a.py"]
#      cwd: "~"
```
