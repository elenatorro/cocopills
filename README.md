# Cocopills

Just another project I will never finish, no hype needed

## Set up

Clone project

```sh
$ git clone git@github.com:elenatorro/cocopills.git
```

Build docker image locally:

```sh
$ docker build -t cocopills/cocopills-docker cocopills-docker
```

The project is in the `cocopills` directory. Run docker image **outside** the directory:

```sh
$ docker run --network host -v /cocopills:/mnt/cocopills -it cocopills/cocopills-docker
```

You'll be in the docker image. Go to the `cocopills` directory **in** in the image and start the server:

```sh
docker@youruser:/$ cd /mnt/cocopills
docker@youruser:/$ go run server/app.go
```

Go to [http://localhost:8080/ping](http://localhost:8080/ping) to test it works
