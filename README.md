Desafio Tunts.Rocks
=============

Table of Contents
-----------------

-   [Install & Setup](#install-&-setup)
-   [Usage](#usage)
-   [License](#license)


Install & Setup
---------------
```html 
git clone https://github.com/ccallazans/desafio-tunts
cd desafio-tunts
```

Run locally:
``` 
go run ./cmd/cli/*.go
```
Run inside docker:
``` 
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.18.5-alpine3.16 go run ./cmd/cli/*.go
```

This will generate an excel file "result.xlsx" on workdir


License
-------

[MIT License](LICENSE)