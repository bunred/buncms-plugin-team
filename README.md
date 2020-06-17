## buncms plugin team

This plugin provides team permission(CRUD) for different users, and needs to be used in [buncms](https://github.com/bunred/buncms)

### How to use

Copy the codes from `.env-example.env` append to `[buncms path]/plugins/.env`

```
import (
...
    "github.com/bunred/buncms-plugin-team"
)
...
    buncms_plugin_user.Routes(types.Routes{Server: server})
...
```