# go-multilang-api

### Description
A demo API that demonstrates how to manage multiple languages in your service.
The project also demonstrates how to deal with pluralization. 
Built using ``gotext``.

There're two supported languages in the repo: ```en-US```, ```fr-FR```.

### Prerequisites
Make sure that ```gotext``` is installed:
```terminal
go install golang.org/x/text/cmd/gotext@latest
```

### How-to
To run service using default port 8080:
```terminal
make run
```
Make a request using different languages
```terminal
curl -H "Accept-Language: en-US" localhost:8080
curl -H "Accept-Language: fr-FR" localhost:8080
```

### Workflow
A typical workflow when working with ```gotext```:

- Extract new strings from project to translate: ``make gen``
- Edit ``internal/translation/locales`` to add translations.
- ``make gen`` to re-build message catalog (``internal/translation/catalog.go``).  

When new languages were added:
- Extract the new language strings from project to translate: ``make gen``
- Create messages to translate by copying output from the prev step (``de-DE`` as an example): 
  ```terminal
    cp internal/translation/locales/de-DE/out.gotext.json internal/translation/locales/de-DE/messages.gotext.json
  ```
- Edit ``internal/translation/locales`` to add translations.
- ``make gen`` to re-build message catalog (``internal/translation/catalog.go``).
###