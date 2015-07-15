IPFS API Spec Tests
===================

This repo contains tests to determine if a given implementation of IPFS conforms to the [IPFS API Spec](https://github.com/ipfs/api)

This test expects the following environment variables

```
IPFS_HOST localhost
IPFS_PORT 5001
```

*You can set the values to whatever you wish*

**Install**

```
$ go get -u github.com/travisperson/ipfs-api-test
$ ipfs-api-test
```
