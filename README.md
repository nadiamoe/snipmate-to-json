Convert snippets in SnipMate format to JSON, as used by VSCode and other snippet plugins.

Hacky, one-shot parser. Does not handle errors or badly-formatted files well.

## Run the thing

```console
$ go install go.nadia.moe/snipmate-to-json@latest
$ snipmate-to-json < markdown.snippets | jq > markdown.json
```
