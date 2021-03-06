# microhttpd

Micro-micro-micro httpd server powered by Go's net/http.

Configuration
------------

## Port

Use the flag `-port` to set the port to run the httpd.

Example:

```
$ microhttpd -port 8080
```

The default port is `80`.

## Files

Add your static files to the directory "www/" and you are ready to roll! It's also possible to pass a custom directory to the flag `-serve`.

Example:

```
$ microhttpd -serve /path/to/my/static/files/
```
Docker
------

There's a Docker image available on the Docker Hub. Just use `microwaves/microhttpd:latest`
in your Dockerfiles or Kubernetes deployment.

Author
------

Stephano Zanzin <@microwaves>
