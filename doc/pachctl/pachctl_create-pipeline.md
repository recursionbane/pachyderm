## ./pachctl create-pipeline

Create a new pipeline.

### Synopsis


Create a new pipeline from a spec, the spec looks like this
{
  "pipeline": {
    "name": "name"
  },
  "transform": {
    "image": "",
    "cmd": [
      "cmd",
      "args..."
    ],
    "stdin": ""
  },
  "shards": "1",
  "inputs": [
    {
      "repo": {
        "name": "in_repo"
      },
      "reduce": false
    }
  ]
}

```
./pachctl create-pipeline -f pipeline.json
```

### Options

```
  -f, --file="-": The file containing the pipeline, - reads from stdin.
```

### Options inherited from parent commands

```
      --log-flush-frequency=5s: Maximum number of seconds between log flushes
```

### SEE ALSO
* [./pachctl](./pachctl.md)	 - 

###### Auto generated by spf13/cobra on 16-Jan-2016
