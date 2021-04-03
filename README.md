# Cocopills

Just another project I will never finish, no hype needed

## Set up

Clone project

```sh
$ git clone git@github.com:elenatorro/cocopills.git
```

Build docker image locally:

```sh
$ cd cocopills
$ docker build -t cocopills/cocopills-docker cocopills-docker
```

Run the docker image in the current directory, setting a custom 'cocopills' volume:

```sh
$ docker run --network host -v "$(pwd):/cocopills" -it cocopills/cocopills-docker
```

You'll be in the docker image. Go to the `cocopills` directory **in** in the image and start the server:

```sh
docker@youruser:/$ cd cocopills
docker@youruser:/$ go run server/app.go
```

Go to [http://localhost:8080](http://localhost:8080)
