# gen-ssh-key
Generates new ssh key, copies it to clipboard (only for MacOS) and appends config file.

Usage:
```bash
gen-ssh-key NAME HOST
```

Arguments:\
`NAME` - Name for key file to idetify it. E.g. "github" will generate files "id_ed25519.github" and "id_ed25519.github.pub"
`HOST` - SSH host. E.g. "github.com"

Example:
```bash
gen-ssh-key corp-gitlab gitlab.corp.com
```

This will produce files *id_ed25519.corp-gitlab* and *id_ed25519.corp-gitlab.pub* and append following to the end of *config* file:

```conf
Host gitlab.corp.com
    Hostname gitlab.corp.com
    IdentityFile /{home}/.ssh/id_ed25519.github
    IdentitiesOnly yes
    AddKeysToAgent yes
    PreferredAuthentications publickey
```