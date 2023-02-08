## Zecrey nft Concurrent test

Within 30 minutes, the number of transactions sent increased gradually from 1000 to 30000

## The test is divided into two modules

### Expectations are correct:

This module is not necessarily correct, and there is competition with each other

### Expected failure:

This module will test the wrong form of each parameter, which must be wrong

## Running

### Before running, configure the 'config. yaml' file. Except for the fields beginning with BoundaryStr, the correct values are required to be written in the file

### Start the main.go file under the test file directly

```
    go run ./mian.go
```

## Operating instructions:

1. Start the go process incrementally through the host to test the concurrency. If there is no abnormality in the nft
   background within 30 minutes, the test can be considered as passed
2. Output the log to the log file and check whether it meets the expectation by checking the log

