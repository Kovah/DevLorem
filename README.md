# DevLorem

Real quotes ready to copy and paste. No more "Lorem ipsum dolor". Get some quotes from President Obama, 
Samuel L Jackson, Daisy Ridley or Morgan Freeman!

![Preview Screenshot](/preview.png)

## Get Quotes via API

Required URL structure:
`your-domain.com/api/[int]`

* The results will be returns as Json by default.
* The maximum allowed amount of quotes is `99`.
* Append `paragraphs=true` as a query parameter to show paragraph tags in the output.
* Append `format=text` as a query parameter to get plain text output.


#### Examples

* `your-domain.com/api/5` would get you 5 quotes in JSON without `<p>` tags
* `your-domain.com/api/15?paragraphs=true` would get you 15 quotes in JSON including the `<p>` tags
* `your-domain.com/api/30?paragraphs=true&format=text` would get you 30 quotes as plain text including the `<p>` tags
* `your-domain.com/api/60?format=text` would get you 60 quotes as plain text without `<p>` tags


---


### Contribution

Want to contribute more source texts? Please create a pull request for the new file that should follow the exact same 
styling like the ones that are already available:

* Filename must contain the name (lowercase, spaces replaced with dashes).
* The JSON strcuture must match the existing files.
* Do not use any <p> tags in the quotes.


---


# Development and Compilation

## Development

You need the following packages on your machine to be able to work on DevLorem and compile the binary:

* github.com/spf13/cobra
* github.com/gorilla/mux
* github.com/GeertJohan/go.rice
* github.com/GeertJohan/go.rice/rice

This can be done easily by running `go mod download` in the current project directory.

To test the package, compile it and then run the resulting executable. I recommend using [Goland]()
for local development. 

## Compilation

To compile the binary and run it locally in the same folder, run the following command:

```
go build
```

To generate a single binary without any dependencies, you have to generate the needed content file for that:
```
rice embed-go
# then run the build command
go build
```

---

DevLorem is a project by [Kovah](https://kovah.de) | [Contributors](https://github.com/Kovah/DevLorem/graphs/contributors)
