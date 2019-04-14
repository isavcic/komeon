# komeon

komeon: Kubernetes InitContainer that waits until wanted Labels are found among the running, dependent Pods

## Usage

The Labels are specified via an environment variable `POD_LABELS` in the following format:

```bash
POD_LABELS="app=foo,bar=baz;app=quux,waldo=fred"
```

In the above example, two sets of Labels are specified and both sets must be found among the running Pods for the InitContainer to exit with code 0, thus successfully initializing the Pod.

Labels are scanned every 2s.

## Alternatives

komeon was inspired by [Yogesh Lonkar's InitContainer implementation](https://github.com/yogeshlonkar/pod-dependency-init-container) and can serve as a simplified, yet extended drop-in replacement, as komeon supports Label sets separated by a semicolon.
