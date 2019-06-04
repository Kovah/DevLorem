# DevLorem

Show paragraphs of real text ready to copy and paste. No more "Lorem ipsum dolor", get some lines from President Obama, 
Jim Carrey or Morgan Freeman.

![Preview Screenshot](/preview.png)

## Get Quotes via API

Required URL structure:
`your-domain.com/api/[int][/p][/json]`

* [int] = optional, number of paragraphs you want
* [/p] = optional, select if the `<p>` tags should be included
* [/json] = optional, output the data in JSON format

#### Examples

* `your-domain.com/api/15/p/json` would get you 15 paragraphs in JSON including the `<p>` tags
* `your-domain.com/api/5/json` would get you 5 paragraphs in JSON without `<p>` tags
* `your-domain.com/api/100/p` would get you 100 paragraphs as plain text including the `<p>` tags
* `your-domain.com/api/100` would get you 100 paragraphs as plain text without `<p>` tags

---

### Contribution

Want to contribute more source texts? Please create a pull request for the new file that should follow the exact same 
styling like the ones that are already available:

* Name the file by the person or thing you want to reference.
* The text file should not contain special characters but can contain spaces.
* Place all paragraphs into `<p></p>` tags.

---

DevLorem is a project by [Kovah](https://kovah.de) | [Contributors](https://github.com/Kovah/DevLorem/graphs/contributors)
