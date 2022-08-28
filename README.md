<div align="center">
    <br><br><img src="https://raw.githubusercontent.com/echgo/logo/main/logo.svg" width="200" />

# echGo

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/echgo/echgo.svg)](https://golang.org/) [![Go](https://github.com/echgo/echgo/actions/workflows/go.yml/badge.svg)](https://github.com/echgo/echgo/actions/workflows/go.yml) [![Docker Image CI](https://github.com/echgo/echgo/actions/workflows/docker-image.yml/badge.svg)](https://github.com/echgo/echgo/actions/workflows/docker-image.yml) [![CodeQL](https://github.com/echgo/echgo/actions/workflows/codeql.yml/badge.svg)](https://github.com/echgo/echgo/actions/workflows/codeql.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/echgo/echgo/v2)](https://goreportcard.com/report/github.com/echgo/echgo/v2) ![Docker Pulls](https://img.shields.io/docker/pulls/echgo/echgo) [![GitHub issues](https://img.shields.io/github/issues/echgo/echgo)](https://github.com/echgo/echgo/issues) [![GitHub forks](https://img.shields.io/github/forks/echgo/echgo)](https://github.com/echgo/echgo/network) [![GitHub stars](https://img.shields.io/github/stars/echgo/echgo)](https://github.com/echgo/echgo/stargazers) [![GitHub license](https://img.shields.io/github/license/echgo/echgo)](https://github.com/echgo/echgo/blob/master/LICENSE) 

This small Docker project is the easiest way to send notifications directly via .txt, .json or .xml files to services like: [Gotify](https://gotify.net/), [Pushover](https://pushover.net/), [Matrix](https://matrix.org/), [Telegram](https://telegram.org/), [Discord](https://discord.com/), [Slack](https://slack.com/), [Trello](https://trello.com/de), [Zendesk](https://www.zendesk.de/), [osTicket](https://osticket.com/), [twillo](https://www.twilio.com/), SMTP (Email) or Webhook. 
    
</div>

## Quick start

Here you can find the instructions on how to set up the Docker container and define the files so that they can be read in correctly.

### Create the configuration files

First, we start the Docker container to create the configuration file. For this you can use the following command.

```console
docker run --name echgo-init -d --rm \
    -v /etc/echgo/configuration:/app/files/configuration \
    echgo/echgo:latest
```

If the container was started, then the directory `/etc/echgo/configuration` is created. Here you will find the configuration for the different communication paths. Please fill in and save this as required. If you want to adjust the configuration. You do not have to restart the Docker container again. The software reads the configuration once before each run, so it is always up-to-date. [Here](https://github.com/echgo/echgo/wiki/Configuration-File#demo-data) you find once the file with demo data, so that you see, how the file looks. If you need more information about the individual keys and tokens, you can find them [here](https://github.com/echgo/echgo/wiki/Environment-Variables#channel-variables) in the channel variables.

The container is stopped automatically and removed.

### Start the service

Now we can start the service directly. To do this, please run the following command once.

```console
docker run --name echgo -d --restart always \
    -v /etc/echgo/configuration:/app/files/configuration \
    -v /var/lib/echgo/notification:/app/files/notification \
    echgo/echgo:latest
```

Now the service should run. With the command we map once the configuration file and the location of the notifications.

### Further adjustments

#### Adjust interval

If you like to adjust the interval of 15 seconds, you can do this with the following variable `INTERVAL`. This must be of type `integer`, otherwise it will not be taken into account. The whole thing looks like this.

```console
docker run --name echgo -d --restart always \
    -e INTERVAL=5 \
    -v /etc/echgo/configuration:/app/files/configuration \
    -v /var/lib/echgo/notification:/app/files/notification \
    echgo/echgo:latest
```

With these settings, the service now reads the notifications every 5 seconds.

#### Use environments

If you want to use environments instead of the json configuration, you can force this with the variable `USE_ENVIRONMENT`. This value is a `boolean`. This is the way json configuration is taken no longer.

```console
docker run --name echgo -d --restart always \
    -e USE_ENVIRONMENT=true \ 
    -v /var/lib/echgo/notification:/app/files/notification \
    echgo/echgo:latest
```

Now you can use the following variables for the different services. If a service is running notification file but has no access data stored, the notification will not be executed. A list of all variables can be found [here](https://github.com/echgo/echgo/wiki/Environment-Variables#channel-variables).

### Create notification

Now we create a notification to be sent to different channels. You can also enter only one channel. How these notification files are created later is up to you. With a bash script or from another program does not matter.

The only important thing is that the file is placed in this folder `/var/lib/echgo/notification`. The name of the file does not matter. It only matters that the file extension and the file format are correct. Currently, we can read the following formats: `.txt, .json & .xml`. 

You can store the following channels in the file, if they are configured: `gotify, pushover, matrix, telegram, discord, slack, trello, zendesk, osticket, twillo, smtp & webhook`. These are always specified in an array. That means you can address one or more channels with one notification file. Now let's look at the currently available file formats and how you can configure them.

#### TXT file

Now let's have a look at an example .txt file. The name of the file can always be freely chosen. It is only important that the data are set per line and that these are stored as key value pairs.

Notification channels can be listed comma separated. If you want to use only one channel, you don't have to use a comma.

```text
channels=gotify,telegram
headline=echGo
message=This is a test message from a txt file.
```

#### JSON file

Here you can find an example for a .json file. Here you can also enter several or only one channel. The structure of the file must be followed please. If you want to know more about JSON, you can find the official site [here](https://www.json.org/json-en.html).

```json
{
    "channels": [
        "gotify",
        "matrix",
        "zendesk"
    ],
    "headline": "Nice headline",
    "message": "This is a test message from a json file."
}
```

#### XML file

The file type .xml can also be used. The structure of the file looks as follows. If you need to XML, you can find it [here](https://www.xml.com/).

```xml
<data>
    <channels>
        <type>gotify</type>
        <type>discord</type>
    </channels>
    <headline>echGo</headline>
    <message>This is a test message from a xml file.</message>
</data>
```

Now echGo reads the files every 15 seconds and sends them to the specified channels. It is also possible to read in several files of different types at the same time.

## Run the service with updates & docker-compose

If you want to get updates for echGo automated, then this is surely exciting for you. Here we use [watchtower](https://github.com/containrrr/watchtower/) to update the container. Watchtower is defined so that it only updates containers with the label `com.centurylinklabs.watchtower.enable=true`. That means you don't have to worry about your other containers.

In order for the echGo service to start properly, you must either do [this](https://github.com/echgo/echgo#create-the-configuration-files) step once before.

Now you can create a docker-compose.yml and start it via ssh in the upload directory with the command `docker-compose up -d` order from version 2 with `docker compose up -d`. Or you can copy the code from here.

```yaml
version: "3.9"
services:
    watchtower:
        container_name: watchtower
        volumes:
            - /var/run/docker.sock:/var/run/docker.sock
        expose:
            - 8080
        restart: always
        image: containrrr/watchtower:latest
        command: --cleanup --include-restarting --rolling-restart --include-stopped --label-enable --interval 3600
    echgo:
        container_name: echgo
        environment:
            - TZ=Europe/Berlin
        volumes:
            - /etc/echgo/configuration:/app/files/configuration
            - /var/lib/echgo/notification:/app/files/notification
        labels:
            - com.centurylinklabs.watchtower.enable=true
        depends_on:
            watchtower:
              condition: service_started
        restart: always
        image: echgo/echgo:latest
```

[Here](https://docs.docker.com/compose/reference/) you can find a list of all docker-compose commands.

If you eventually want to run multiple servers with echgo, then this might still be interesting for you. Here I have set up a NFS server on which the echgo configuration file is located and create a mount on this server in the volume `echgo_configuration` and use this for the echgo container. A guide for NFS servers and how to use them can be found [here](https://ubuntu.com/server/docs/service-nfs). But please remember to enter the IP of the client server at the NSF server before you start the services via docker compose file.

```yaml
version: "3.9"
services:
    watchtower:
        container_name: watchtower
        volumes:
            - /var/run/docker.sock:/var/run/docker.sock
        expose:
            - 8080
        restart: always
        image: containrrr/watchtower:latest
        command: --cleanup --include-restarting --rolling-restart --include-stopped --label-enable --interval 3600
    echgo:
        container_name: echgo
        environment:
            - TZ=Europe/Berlin
        volumes:
            - /var/lib/echgo/notification:/app/files/notification
            - echgo_configuration:/app/files/configuration
        labels:
            - com.centurylinklabs.watchtower.enable=true
        depends_on:
          watchtower:
            condition: service_started
        restart: always
        image: echgo/echgo:latest
volumes:
    echgo_configuration:
        driver: local
        driver_opts:
            type: nfs
            o: nfsvers=4,addr=1.2.3.4,ro,async
            device: :/mnt/docker/echgo
```

If you want to use this docker-compose, just copy the part and save it in a docker-compose.yml file. Then you can start directly with it.

## Planned channels

Here you will find channels we have planned or already implemented. If you think of another one, please send it to us.

- Microsoft Teams
- WhatsApp Business
- HubSpot

## Added channels

Here you can find all added channels and in which version they are added.

- Matrix - _Added in version v0.0.3_
- Trello - _Added in version v0.0.4_
- Discord - _Added in version v0.0.6_
- Zendesk - _Added in version v0.0.7_
- osTicket - _Added in version v0.1.2_
- Slack - _Added in version v0.1.3_
- twilio - _Added in version v1.0.3_
- Pushover - _Added in version v1.1.7_ 

## Special thanks

Thanks to [JetBrains](https://github.com/JetBrains) for supporting me with this and other [open source projects](https://www.jetbrains.com/community/opensource/#support).
