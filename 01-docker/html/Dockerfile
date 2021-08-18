# Para buildar com alias (Buildah não suporta --target):
# docker build -t nginx:1.0 . --target img-nginx
FROM nginx:latest as img-nginx
WORKDIR /app
RUN apt-get update && \
    apt-get install -y \
    vim
COPY public/ /usr/share/nginx/html

# docker build -t ubuntu:1.0 . --target img-ubuntu
FROM ubuntu:latest as img-ubuntu
ENTRYPOINT ["echo","Hello"]
CMD ["World"]