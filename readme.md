#

## needed files 

`key.pem` -- from google accounts

`gcal.ini` looks like
> serviceEmail=blah@foobar.iam.gserviceaccount.com
> pemFile=key.pem


## local testing
```
GOPATH="$GOPATH:$HOME/src/dbexperements/go/" go test -v lncdcal_test.go
```
