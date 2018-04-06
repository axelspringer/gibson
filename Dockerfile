FROM scratch
ADD bin/gibson_linux_amd64 /gibson
CMD ["/gibson"]
