<div align="center">
    <br><br><img src="https://raw.githubusercontent.com/echgo/logo/main/echGo.svg" width="200" />

# echGo

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/echgo/echgo.svg)](https://golang.org/) [![Go](https://github.com/echgo/echgo/actions/workflows/go.yml/badge.svg)](https://github.com/echgo/echgo/actions/workflows/go.yml) [![Docker Image CI](https://github.com/echgo/echgo/actions/workflows/docker-image.yml/badge.svg)](https://github.com/echgo/echgo/actions/workflows/docker-image.yml) [![CodeQL](https://github.com/echgo/echgo/actions/workflows/codeql.yml/badge.svg)](https://github.com/echgo/echgo/actions/workflows/codeql.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/echgo/echgo)](https://goreportcard.com/report/github.com/echgo/echgo) ![Docker Pulls](https://img.shields.io/docker/pulls/echgo/echgo) [![GitHub issues](https://img.shields.io/github/issues/echgo/echgo)](https://github.com/echgo/echgo/issues) [![GitHub forks](https://img.shields.io/github/forks/echgo/echgo)](https://github.com/echgo/echgo/network) [![GitHub stars](https://img.shields.io/github/stars/echgo/echgo)](https://github.com/echgo/echgo/stargazers) [![GitHub license](https://img.shields.io/github/license/echgo/echgo)](https://github.com/echgo/echgo/blob/master/LICENSE) 

This small Docker project is the easiest way to send notifications directly via .txt, .json, .yaml or .xml files to services like: [Gotify](https://gotify.net/), [Matrix](https://matrix.org/), [Telegram](https://telegram.org/), [Discord](https://discord.com/), [Slack](https://slack.com/), [Trello](https://trello.com/de), [Zendesk](https://www.zendesk.de/), [osTicket](https://osticket.com/), SMTP (Email) or Webhook.
    
</div>

## Quick start

Here you can find the instructions on how to set up the Docker container and define the files so that they can be read in correctly.

### Create the configuration files

First, we start the Docker container to create the configuration file. For this you can use the following command.

```console
docker run --name echgo-init -d --rm \
    -v /etc/echgo/configuration:/go/src/app/files/configuration \
    echgo/echgo:latest
```

If the container was started, then the directory **/etc/echgo/configuration** is created. Here you will find the configuration for the different communication paths. Please fill in and save this as required. You can find an example of the configuration file [here](https://github.com/echgo/echgo/blob/master/.echgo.yaml). If you want to adjust the configuration. You do not have to restart the Docker container again. The software reads the configuration once before each run, so it is always up-to-date.

The container is stopped automatically and removed.

### Start the service

Now we can start the service directly. To do this, please run the following command once.

```console
docker run --name echgo -d --restart always \
    -v /etc/echgo/configuration:/go/src/app/files/configuration \
    -v /var/lib/echgo/notification:/go/src/app/files/notification \
    echgo/echgo:latest
```

Now the service should run. With the command we map once the configuration file and the location of the notifications.

### Create notification

Now we create a notification to be sent to different channels. You can also enter only one channel. How these notification files are created later is up to you. With a bash script or from another program does not matter.

Currently you can store the following channels in the file, if they are configured: **gotify, matrix, telegram, discord, slack, trello, zendesk, osticket, smtp & webhook**.

It does not matter what the TXT file is called, which is stored in the directory **/var/lib/echgo/notification** defined by us. The only important thing is that we currently only support .txt files. But there are more file types planned.

**Important! With version v1.0.0 we have removed the text error channel and adjusted it to plural channels. This may need to be considered in the scripts. The data types are final now and should not be adjusted anymore!**

Now let's create a test.txt file in the directory with the following content:

```text
channels=gotify,telegram
headline=echGo
message=This is a test message from a txt file.
```

If we now want to send only one webhook, we can do this with the following text file.

```text
channels=webhook
headline=echGo
message=This is a test message from a txt file.
```

Instead of a TXT file you can now also store a JSON file in the following directory: **/var/lib/echgo/notification**. In this file one or more of the channels can be defined. **Important! The data format has been adapted from version v0.2.0. Please take into account.**

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

In this example we want to send the notification to [Gotify](https://gotify.net/) & [Telegram](https://telegram.org/). Thereby we set a headline and the message. You can also use a YAML file. This must then be structured as follows:

```yaml
channels:
    - "gotify"
    - "slack"
headline: "echGo"
message: "This is a test message from a yaml file."
```

If you prefer to use an XML file, you can do that as well. The file can be read in like the others. It must be in the following format:

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

Now echGo reads the files every minute and sends them to the specified channels.

## Run the service with updates & docker-compose

If you want to get updates for echGo automated, then this is surely exciting for you. Here we use [watchtower](https://github.com/containrrr/watchtower/) to update the container. Watchtower is defined so that it only updates containers with the label **com.centurylinklabs.watchtower.enable=true**. That means you don't have to worry about your other containers.

In order for the echGo service to start properly, you must either do [this](https://github.com/echgo/echgo#create-the-configuration-files) step once before or place [this file](https://raw.githubusercontent.com/echgo/echgo/master/.echgo.yaml) with your correct data on your system in the **/etc/echgo/configuration** directory.

Now you can download [this file](https://raw.githubusercontent.com/echgo/echgo/master/docker-compose.yaml) and start it via ssh in the upload directory with the command **docker-compose up -d**. Or you can copy the code from here.

```yaml
version: "3.5"
services:
    echgo:
        container_name: echgo
        networks:
            - default
        volumes:
            - /etc/echgo/configuration:/go/src/app/files/configuration
            - /var/lib/echgo/notification:/go/src/app/files/notification
        labels:
            - com.centurylinklabs.watchtower.enable=true
        restart: always
        image: echgo/echgo:latest
    watchtower:
        container_name: watchtower
        networks:
            - default
        volumes:
            - /var/run/docker.sock:/var/run/docker.sock
        ports:
            - 8080:8080
        restart: always
        image: containrrr/watchtower:latest
        command: --cleanup --include-restarting --rolling-restart --include-stopped --label-enable --interval 3600
networks:
    default:
```

[Here](https://docs.docker.com/compose/reference/) you can find a list of all docker-compose commands.

If you eventually want to run multiple servers with echgo, then this might still be interesting for you. Here I have set up a NFS server on which the echgo configuration file is located and create a mount on this server in the volume **echgo_configuration** and use this for the echgo container. A guide for NFS servers and how to use them can be found [here](https://ubuntu.com/server/docs/service-nfs). You can also start this with the usual commands like **docker-compose up -d**. But please remember to enter the IP of the client server at the NSF server before.

```yaml
version: "3.5"
services:
    echgo:
        container_name: echgo
        networks:
            - default
        volumes:
            - /var/lib/echgo/notification:/go/src/app/files/notification
            - echgo_configuration:/go/src/app/files/configuration
        labels:
            - com.centurylinklabs.watchtower.enable=true
        restart: always
        image: echgo/echgo:latest
    watchtower:
        container_name: watchtower
        networks:
            - default
        volumes:
            - /var/run/docker.sock:/var/run/docker.sock
        ports:
            - 8080:8080
        restart: always
        image: containrrr/watchtower:latest
        command: --cleanup --include-restarting --rolling-restart --include-stopped --label-enable --interval 3600
volumes:
    echgo_configuration:
        driver: local
        driver_opts:
            type: nfs
            o: nfsvers=4,addr=1.2.3.4,ro,async
            device: :/mnt/docker/echgo
networks:
    default:
```

If you want to use this docker-compose, just copy the part and save it in a docker-compose.yaml file. Then you can start directly with it.

## Planned channels

Here you will find channels we have planned or already implemented. If you think of another one, please send it to us.

- Twilio
- Ntfy
- Message Bird
- Pushover
- Microsoft Teams

## Added channels

Here you can find all added channels and in which version they are added.

- Matrix - _Added in version v0.0.3_
- Trello - _Added in version v0.0.4_
- Discord - _Added in version v0.0.6_
- Zendesk - _Added in version v0.0.7_
- osTicket - _Added in version v0.1.2_
- Slack - _Added in version v0.1.3_