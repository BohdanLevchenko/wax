# Wax
Wax is a simple and minimalistic http redirector.

The main idea was to create compact crossplatform http server which redirects http requests according to own configuration.

Jumpstart with:

```shell
  ./build.sh
  bin/wax
```

Configuration is dead simple:
```json
{
  "mappings": [{
      "source": "/yandex",
      "target": "http://www.yandex.ru"
    }, {
      "source": "/google",
      "target": "http://www.google.com"
    }, {
      "source": "/space",
      "target": "http://www.nasa.gov"
    }
  ]
}
```

It's enables the following redirections:

```shell
  http://localhost:8001/yandex => http://www.yandex.ru
  http://localhost:8001/google => http://www.google.com
  http://localhost:8001/space => http://www.nasa.gov
```

By default wax respond with 301 Moved permanently status code and listen on 8001 port, but this can be overridden with -code and -port command line parameters.

See wax -help for details.
