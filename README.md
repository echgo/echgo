<div align="center">
    <br><br><img src="https://raw.githubusercontent.com/echgo/logo/main/echGo.png" width="200" />

# echGo

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/echgo/echgo.svg)](https://golang.org/) [![Go](https://github.com/echgo/echgo/actions/workflows/go.yml/badge.svg)](https://github.com/echgo/echgo/actions/workflows/go.yml) [![Docker Image CI](https://github.com/echgo/echgo/actions/workflows/docker-image.yml/badge.svg)](https://github.com/echgo/echgo/actions/workflows/docker-image.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/echgo/echgo)](https://goreportcard.com/report/github.com/echgo/echgo) ![Docker Pulls](https://img.shields.io/docker/pulls/echgo/echgo) [![GitHub issues](https://img.shields.io/github/issues/echgo/echgo)](https://github.com/echgo/echgo/issues) [![GitHub forks](https://img.shields.io/github/forks/echgo/echgo)](https://github.com/echgo/echgo/network) [![GitHub stars](https://img.shields.io/github/stars/echgo/echgo)](https://github.com/echgo/echgo/stargazers) [![GitHub license](https://img.shields.io/github/license/echgo/echgo)](https://github.com/echgo/echgo/blob/master/LICENSE) 

This small Docker project is the easiest way to send notifications directly via .txt files to services like: [Gotify](https://gotify.net/), [Telegram](https://telegram.org/), SMTP (Email) or Webhook.
    
</div>

## Quick start

Here you can find the instructions on how to set up the Docker container and define the files so that they can be read in correctly.

### Create the configuration files

First, we start the Docker container to create the configuration file. For this you can use the following command.

```console
docker run --name echgofig -d --rm \
    -v /etc/echgo/configuration:/go/src/app/files/configuration \
    echgo/echgo:latest
```

If the container was started, then the directory **/etc/echgo/configuration** is created. Here you will find the configuration for the different communication paths. Please fill in and save this as required.

The container is stopped automatically and can now be removed.

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

It does not matter what the TXT file is called, which is stored in the directory **/var/lib/echgo/notification** defined by us. The only important thing is that we currently only support .txt files. But there are more file types planned.

Now let's create a test.txt file in the directory with the following content:

```text
channel=gotify,telegram
headline=echGo test 
message=This is a test message!
```

If we now want to send only one webhook, we can do this with the following text file.

```text
channel=webhook
headline=Nice webhook headline! 
message=This is the best message for the webhook.
```

In this example we want to send the notification to [Gotify](https://gotify.net/) & [Telegram](https://telegram.org/). Thereby we set a headline and the message.  Now echGo reads the files every minute and sends them to the specified channels.




