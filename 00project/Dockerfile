# initialize image
FROM golang:alpine

# install git
RUN apk update && apk add git

# set working directory in container
WORKDIR /app

# copy all from host (this project directory) 
COPY . .

# validate our project
RUN go mod tidy

# build as binary
RUN go build -o binary

ENTRYPOINT [ "./binary" ]

# build as name my-app-image: docker build . -t my-app-image

# see images: docker images

# run image: docker run -i -t my-app-image

# -i for enable interactive model
# -t for enable tty

# add -e for add env
# run image: docker run -i -t -e PORT=8083 my-app-image

# add -p host_port:container_port
# for expose from container to host

###################################
# flag -t
# for tagging/name a image
# => docker build . -t <namaimage>:<tagnya>

# => docker container ls --a
# show all container list

# => docker container start <idContainer>
# run container

# docker container create -e PORT=8083 -p 8083:8083 my-app-image
# create container


#################################

# => docker exec -it <idContainer> sh
# connect/remote to container

# => docker exec -it <idContainer> <your command>
# connect/remote to container and exec command

#################################

# => docker run -itd --name belajarssh -e MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=YES -v D:/docker-volume:/var/lib/mysql -p 3306:3306 mariadb
# run mariadb image as cotainer

#################################
