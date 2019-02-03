# corectl
Management tool for g8os

# Usage

```bash
corectl -h
```
```raw
NAME:
   corectl - manage g8os

USAGE:
   Query or send commands to g8os manager
   
VERSION:
   1.0
   
COMMANDS:
     ping     checks connectivity with g8os
     execute  execute arbitary commands
     stop     stops a process with `id`
     info     query various infomation
     reboot   reboot the machine
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --socket value, -s value   Path to core socket (default: "/var/run/core.sock")
   --timeout value, -t value  Commands that takes longer than this will get killed (default: 0)
   --async                    Run command asyncthronuslly (only commands that supports this)
   --id value                 Speicify porcess id, if not given a random guid will be generated
   --container value          Container numeric ID or comma seperated list with tags (only with execute)
   --help, -h                 show help
   --version, -v              print the version

```

```bash
corectl info -h
```
```raw
NAME:
   corectl info - query various infomation

USAGE:
   corectl info command [command options] [arguments...]

COMMANDS:
   cpu		display CPU info
   memory, mem	display memory info
   disk		display disk info
   nic		display NIC info
   os		display OS info
   process, ps	display processes info
   help, h	Shows a list of commands or help for one command
   
OPTIONS:
   --help, -h	show help
```
