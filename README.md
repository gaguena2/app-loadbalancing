docker build -t gaguena/loadbalancing .


docker run --rm -p 8080:8080 --name loadbalancing gaguena/loadbalancing 