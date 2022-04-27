# QueryParameter-to-Header Middleware
Traefik middleware to add a header with the value of a given URL query parameter

`queryParameter` and `header` keys are configurable. It is assummed that there is only one query parameter with the given key. If there are several, the first one in the query string is chosen. If there is no matching query parameter, no header is added.

Tested with Traefik v2.6.3.

Example Traefik config:
``` yaml
[...]
experimental:
  plugins:
    queryParameterToHeader:
      moduleName: github.com/corticph/queryparameter-to-header
      version: v1.0.0

http:
[...]
  routers:
    whoami:
      rule: "Host(`whoami.localhost`)"
      service: "whoami"
      entryPoints:
        - web
      middlewares:
        - add-header-from-query
  
  middlewares:
    add-header-from-query:
      plugin:
        queryParameterToHeader:
          queryParameter: "v"
          header: "X-Version"
```