FROM scratch
ADD timestampd /
EXPOSE 8080
CMD ["/timestampd"]
