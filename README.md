# bcrypt-password-hash

This program create safe hashed value for storing password. 

To learn how to use this program just call it and read the help

## Build

You need golang installed

If you have installed golang in a custom location 
please open the file .env and change GOROOT accordingly.

Then just run this command from this directory:

```shell
GOPATH=$PWD go get -t ./...
go install bcrypt-password-hash
```

Totally ignore any errors... believe me it should work.

Now you compile you can call it by doing so from this directory:

```shell
./bin/bcrypt-password-hash hash 'vlad' 14
or
./bin/bcrypt-password-hash check '$2a$15$nm2fGmbuDv/CDBgJUGF9OOaVEgYERkhA1IRdg0f/12yYhJ/qtA50W' 'vlad'
```

