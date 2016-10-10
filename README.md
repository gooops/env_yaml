ENV YAML
========

**this package based on `github.com/gooops/env_strings`**

Read Env from config and compile the values into yaml file.

### About ENV YAML

Sometimes we need config file as following:

`db.conf`

```yaml
{
	"host":"127.0.0.1",
	"password":"3306",
	"timeout": 1000
}
```

but, when we management more and more server and serivce, and if we need change the password or ip, it was a disaster.

So, we just want use config like this.

`db.conf`

```yaml
---
host: "{{.host}}"
password: "{{.password}}"
timeout: "{{.timeout}}"
```

We use golang's template to replace values into the config while we read the yaml file configs.

first, we set the env config at `~/.bash_profile` or `~/.zshrc`, and the default key is `ENV_YAML` and the default file extention is `.env`, the value of `ENV_YAML` could be a file or folder,it joined by`;`, it will automatic load all `*.env` files.

**Env**

```bash
export ENV_YAML='~/playgo/test.env;~/playgo/test2.env'
```

or

```bash
export ENV_YAML='~/playgo'
```


#### example program

```go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gooops/env_yaml"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Timeout  int64  `yaml:timeout`
}

func main() {
	data, _ := ioutil.ReadFile("./db.conf")

	dbConf := DBConfig{}

	if err := env_yaml.Unmarshal(data, &dbConf); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(dbConf)
}
```


`env1.yaml`

```yaml
---
host: "127.0.0.1",
password: "123456"

```


`env2.yaml`

```yaml
---
"timeout": 1000
```

**result:**

```bash
{127.0.0.1 123456 1000}
```

### More

if you want use your own `ENV` key, you could do it like this

```go
envYaml := NewEnvYaml("YOUR_ENV_KEY_NAME", ".yaml")
envyaml.Unmarshal(data, &dbConf);
```
