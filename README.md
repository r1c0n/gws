# Gamma Web Server - GWS

A lightweight and fast web server written in Go.

## Features

- Supports HTTP and HTTPS
- Easy & quick setup
- Easy to use configuration file

## Installation

Download the latest release from GitHub, or find the latest release on our website: [https://www.gammaws.gq](https://www.gammaws.gq)

## Usage

To start Gamma Web Server, double click on the .exe file, or run it through the command-line.

```cmd
.\gws.exe
```

To see your website, visit the link showed on the command-line below the product information. By default, the link is [localhost:8080](localhost:8080).

## Configuration

Gamma Web Server uses a JSON configuration file to customize its behavior. You can specify the port, the document root & more in the configuration file. Find the configuration file in the project's root directory. It will be called **config.json**.

Here is an example of what the **config.json** should look like.

```json
{
    "port": ":80",
    "tls_config": {
      "cert_file": "server.crt",
      "key_file": "server.key"
    },
    "static_dir": "public",
    "repo_config": {
      "version": "1.1.0",
      "author": "Official B",
      "product": "Gamma Web Server",
      "repository": "https://github.com/gamma-gws/gws"
    }
  }  
```

## Misc. Info

![](https://img.shields.io/github/go-mod/go-version/r1con/gws) ![](https://img.shields.io/github/v/release/r1c0n/gws?include_prereleases) ![](https://img.shields.io/github/license/r1c0n/gws)
