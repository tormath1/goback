## GoBack

Keep an eye on your Docker volumes !

## Todo

  * [ ] create a file / sqlite in order to keep current cron entries
  * [ ] remove a cron job, waiting for next release of cron lib

## Getting started

Download the latest [release](https://github.com/tormath1/goback/releases): `server` is the API server and `goback` is the command line tool. 

As `server` will manage your volumes, you'll need to start it with `sudo` command. Once your server up, you'll need to move `goback` in a `$PATH` location like `/usr/local/bin`.

Now you are ready to backup your volumes ! 

## How to setting up Goback

Before starting with the CLI, you need to run the server on your docker host. 

You can start it without authentication or with TLS authentication: 

```
$ sudo ./server
```

or 

```
$ sudo ./server \
--cert-file=/path/to/cert.crt \
--key-file=/path/to/key.key
```

If you run self signed certificates, you need to add your root authority: 

```
$ sudo ./server \
--cert-file=/path/to/cert.crt \
--key-file=/path/to/key.key \
--root-ca=/path/to/root.crt 
```

And with client, you need to provide a json config file, called `goback.json` like this one: 

```json
{
    ...
    "tls": {
        "cert_file": "/path/to/client.local.tld.crt",
        "key_file": "/path/to/client.local.tld.key",
        "root_cert": "/path/to/rootCA.pem",
        "server_name": "server.local.tld"
    }
    ...
}
```

## Backup your volumes

First, you need to identify which volume (id or name) you want to backup.

```shell
awesome-guy@data-instance-001:~$ docker volume ls
DRIVER              VOLUME NAME
local               098ee8a3201b24011817b955c1b445707b0156112bc4db1c2600c9e743dbadc6
local               0bfdac778e6f4f3057910e613fd3ab6033cfb00b766810e8734c67ce303cc4f8
local               data_example_1
local               data_example_2
```

Let's suppose, I want to save `data_example_1` on `/mnt/path`, I will run this command:

```shell
awesome-guy@data-instance-001:~$ goback save data_example_1 /mnt/path
```

Pretty simple isn't it ?

Now, I want to save `data_example_1` each thirty minutes:

```shell
awesome-guy@data-instance-001:~$ goback schedule data_example_1 /mnt/path  "0 30 * * * *"
```

Following this cron table:

```
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
```

(For this cron table, credits go to [godoc](https://godoc.org/github.com/robfig/cron))

Finally, you can get a list of your current volumes backing up:

```shell
awesome-guy@data-instance-001:~$ goback schedule list
volume:"data_example_1" cron:"0 30 * * * *"
```

## gRPC support

Server is designed with gRPC APIs, so you can find proto file and generate your client, to use goback directly in your workflow. 

## Metrics

List of Prometheus metrics:
  * Counter `number_cron_job`: Number of cron job