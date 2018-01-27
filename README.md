## GoBack

Keep an eye on your Docker volumes !


## How to backup your volumes

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
awesome-guy@data-instance-001:~$ goback save data_example_1 /mnt/path each "0 30 * * * *"
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
awesome-guy@data-instance-001:~$ goback volume ls
DRIVER              VOLUME NAME
/mnt/path           data_example_1
```
