## Adding Godeps

If your change adds godeps, you'll need to save the new dependencies and commit them. We use `godep save -r` to vendor dependencies because this allows the package to be `go get`'able.

```console
$ godep restore # You only need to do this if `godep save` fails with unkown packages.
$ godep save -r ./...
$ git add Godeps/_workspace/src
```
