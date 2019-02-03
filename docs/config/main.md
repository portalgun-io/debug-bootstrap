# Main Configuration

The main configuration is auto-loaded from the `zero-os.toml` file.

In the [zero-os/0-initramfs](https://github.com/zero-os/0-initramfs) repository, used for creating a Zero-OS kernel, `zero-os.toml` can be found in the `/config/etc/zero-os/zero-os.toml` directory.

`zero-os.toml` has the following sections:

- [\[main\]](#main)
- [\[containers\]](#containers)
- [\[logging\]](#logging)
- [\[stats\]](#stats)
- [\[globals\]](#globals)
- [\[extension\]](#extension)


<a id="main"></a>
## [main]

```toml
[main]
max_jobs = 200
include = "/config/root"
network = "/config/g8os/network.toml"
```

- **max_jobs**: Max parallel jobs the core can execute concurrently (as its own direct children), once this limit is reached 0-core will not pull for any new jobs from its dedicated Redis queue until it has at least one free job slot to fill
- **include**: Path to the directory with TOML files to include, this directory can have configurations for startup services and extensions, when Zero-OS boots it will try to load all `.toml` files from the given locations, each of these TOML file can define one or more extensions to the 0-core commands, and/or start up services
- **network**: Path to the network configuration file, discussed in [Network Configuration](network.md)


<a id="containers"></a>
## [containers]
Contains containers creation limits

```
[containers]
max_count = 300 (max number of running containers, defaults to 1000 if not set)
```


<a id="logging"></a>
## [logging]

In this section you define how Core0 processes logs from running processes.

There are 2 built in loggers that are used by zero-os to log jobs outputs that can be refined by the following two seconds

- **logging.file**: writes logs to `/var/log/core.log`
- **ledis**: forwards logs to Ledis

For each logger you define log levels, specifying which log levels are logged to this logger.

Example:

```
[logging.file]
levels = [2, 4, 7, 8, 9]

[logging.ledis]
levels = [1, 2, 4, 7, 8, 9]
size = 1000
```

In the above example:

- The second logger, of type `ledis`, specifies with `size` how many log messages are kept in the queue before older log messages will get dropped

See the section [Logging](../monitoring/logging.md) for more details about logging.

<a id="stats"></a>
## [stats]

This is where the statistics is configured.

Here's an example:

```
[stats]
enabled = true
```

See [Monitoring](../monitoring/README.md) for more details about statistics.


<a id="globals"></a>
## [globals]

Here all global module parameters are set.

Example:

```
[globals]
storage = "ardb://hub.gig.tech:16379"
```

With `storage` you set the default key-value store that will be mounted by the [Zero-OS File System](https://github.com/zero-os/0-fs) when creating containers using the [container.create()](../interacting/commands/container.md#create) command. The default, as shown above, is the ARDB storage cluster implemented in [0-Hub](https://github.com/zero-os/-hub?). When creating a new container you can override this default by specifying any other ARDB storage cluster, as documented in [Creating Containers](../containers/creating.md).


<a id="extension"></a>
## [extension]

An extension is simply a new command or functionality to extend what Zero-OS can do. This allows you to add new functionality and commands to Zero-OS without actually changing its code. An extension works as a wrapper around the `core.system` command by wrapping the actual command call.

The below example is a user management extension for adding and removing users, and changing their passwords:

```toml
[extension."user.add"]
binary = "useradd"
args = ["-m", "{username}"]

[extension."user.delete"]
binary = "userdel"
args = ["-f", "-r", "{username}"]

[extension."user.chpasswd"]
binary = "sh"
args = ["-c", "echo '{username}:{password}' | chpasswd"]
```

Adding the above into a TOML file and saving it in one of the paths specified in the `include` section of the main configuration file will add the following commands to Zero-OS:

 - **user.add**
   - Args: `{"username": "user name to add"}`
 - **user.delete**
   - Args: `{"username": "user name to remove"}`
 - **user.chpasswd**
   - Args: `{"username": "user", "password": "password to set"}`

This allows you to call the extension from the Python client as follows:

```python
client.raw("user.add", {"username": "testuser"})
client.raw("user.chpasswd", {"username": "testuser", "password": "new-password"})
```

> Core0 takes care of substituting the `{key}` notation in the extension arguments with the ones passed from the client.

Extension also supports the following attributes:

```toml
[extension.test]
binary = "binary"
args = ["args", "list", "to", "binary"]
# cwd of the binary
cwd = "/path"

#env variables will be available during command execution
[extension.test.env]
env1 = "value-1"
env2 = "value-2"
```
