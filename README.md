# retry

Simple retry tool.

## Installation

	go get [-u] github.com/suzuken/retry

## Usage

```
$ retry
Usage: retry <command>

  -initialInterval int
        retry interval(s) (default 1)
  -maxElapsedTime int
        Max Elapsed Time(s) is limit of backoff steps. If the job spends over this, job makes stopped. If set 0, the job will never stop. (default 10000)
  -maxInterval int
        cap of retry interval(s) (default 1000)
```

Retry list files until success.

```
$ retry ls /tmp/
```

If command failed, retry using exponential backoff. Strategy of backoff is depends on [cenkalti/backoff](https://github.com/cenkalti/backoff).

```
$ retry command_not_found
2015/08/26 16:55:14 err: exec: "command_not_found": executable file not found in $PATH
2015/08/26 16:55:15 err: exec: "command_not_found": executable file not found in $PATH
2015/08/26 16:55:17 err: exec: "command_not_found": executable file not found in $PATH
2015/08/26 16:55:20 err: exec: "command_not_found": executable file not found in $PATH
...
```

You can specify `maxElapsedTime`, as below:

```
$ retry -initialInterval=2 -maxElapsedTime=8 command_not_found
2015/08/26 16:57:37 err: exec: "command_not_found": executable file not found in $PATH
2015/08/26 16:57:39 err: exec: "command_not_found": executable file not found in $PATH
2015/08/26 16:57:43 err: exec: "command_not_found": executable file not found in $PATH
2015/08/26 16:57:49 err: exec: "command_not_found": executable file not found in $PATH
operation failed: exec: "command_not_found": executable file not found in $PATH
```

## LICENSE

MIT
