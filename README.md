<h1 align="center">DevLorem</h1>

<p align="center"><b>Real quotes ready to copy and paste. No more "Lorem ipsum dolor".</b><br>Get some quotes from President Obama, 
Samuel L Jackson, Daisy Ridley or Morgan Freeman!</p>

<p>&nbsp;</p>

<img src="/preview.png" alt="DevLorem Preview">

<p>&nbsp;</p>


## About DevLorem

DevLorem is a small tool that generates Lorem Ipsum paragraphs from movie quotes, based on the different actors.  
It can be used as a command line tool, or can be run as a website that offers both a user interface to generate the
paragraphs and an API.


---


## Download DevLorem

The latest version of the tool can be found on the [releases page](https://github.com/Kovah/DevLorem/releases).
Please download the archive file of the latest release suitable for your operating system. Unpack the zip and make the
binary executable.  
Alternatively, you can use the Docker image. Details about the usage can be found further down.


## The command line tool

The command line tool can be used to quickly generate paragraphs within your terminal or in scripts. It is available
as the `generate` command and has some optional parameters.

```
devlorem generate
```

Usage:  
  devlorem generate [flags]

Flags:  
  -f, --format string        Format of the returned paragraphs, either json or text (default json)  
  -h, --help                 help for generate  
  -n, --number int           Number of paragraphs returned (default 5)  
  -p, --paragraphs           Show paragraph tags (<p> and </p>) in the generated paragraphs


#### Examples

* `devlorem generate` would get you 5 quotes in JSON without `<p>` tags
* `devlorem generate -n 15 -p` would get you 15 quotes in JSON including the `<p>` tags
* `devlorem generate -n 30 -p -f text` would get you 30 quotes as plain text including the `<p>` tags
* `devlorem generate -n 60 -f text` would get you 60 quotes as plain text without `<p>` tags


---


## The website (user interface + API)

The DevLorem executable ships with a built-in web server including all static assets. You can start the web server by
using the `serve` command. By default, the started web server listens to port 80 on your host.

```
$ devlorem serve
Starting HTTP server for DevLorem...
```

You can now open DevLorem in your browser under `http://localhost`.

If you want to use another port, you can use the optional bind flag. To change the port, specify the port including
preceding colon. In the following example we tell DevLorem to use the port 8090:

```
devlorem serve -b :8090
```


### Usage of the website API

Once the web server is started, the DevLorem API is available too. In the following examples I assume that you have set
up DevLorem behind a proxy with the domain `your-domain.com` and HTTPS configured.

Required URL structure:
`https://your-domain.com/api/[int]`

* The results will be returns as Json by default.
* The maximum allowed amount of quotes is `99`.
* Append `paragraphs=true` as a query parameter to show paragraph tags in the output.
* Append `format=text` as a query parameter to get plain text output.


#### Examples

* `https://your-domain.com/api/5` would get you 5 quotes in JSON without `<p>` tags
* `https://your-domain.com/api/15?paragraphs=true` would get you 15 quotes in JSON including the `<p>` tags
* `https://your-domain.com/api/30?paragraphs=true&format=text` would get you 30 quotes as plain text including the `<p>` tags
* `https://your-domain.com/api/60?format=text` would get you 60 quotes as plain text without `<p>` tags


---


## The DevLorem Docker image

DevLorem is also available as a [Docker image](https://hub.docker.com/r/kovah/devlorem). It is built using Alpine
Linux and is less than 10 MB large.

To use the command line tool, run the Docker image with the `generate` command. More details about the tool can be
found in the command line documentation above.

```
docker run --rm kovah/devlorem generate
```

To use the website, you have to additionally forward a port. Please notice that DevLorem does not support HTTPS
connections, so you probably need a proxy in front of it.  
You can start the web server by using the `serve` command. More details about the command can be found in the website 
documentation above.

```
docker run --rm -p 80:80 kovah/devlorem serve
```


---


## Contribution

Want to contribute more source texts? Please create a pull request for the new file that should follow the exact same 
styling like the ones that are already available:

* Filename must contain the name (lowercase, spaces replaced with dashes).
* The JSON structure must match the existing files.
* Do not use any <p> tags in the quotes.


---


## Development and Compilation

### Development

You need the following packages on your machine to be able to work on DevLorem and compile the binary:

* github.com/spf13/cobra
* github.com/gorilla/mux
* github.com/GeertJohan/go.rice
* github.com/GeertJohan/go.rice/rice

This can be done by running `go mod download` in the current project directory.

To test the package, compile it and then run the resulting executable. I recommend using Goland
for local development. 


### Compilation

To compile the binary and run it locally in the same folder, run the following command:

```
npm run build-prod
go build
```

To generate a single binary without any dependencies, you have to generate the needed content file for that:
```
rice embed-go
# then run the build command
go build
```

To test the current source code, run the following command:
```
go test -run ''
```


---


DevLorem is a project by [Kovah](https://kovah.de) | [Contributors](https://github.com/Kovah/DevLorem/graphs/contributors)
