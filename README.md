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

DevLorem is a project by [Kovah](https://kovah.de) | [Contributors](https://github.com/Kovah/DevLorem/graphs/contributors)
