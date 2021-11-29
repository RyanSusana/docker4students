FROM ubuntu

RUN apt-get update
RUN apt-get install apache2-utils -y

COPY ddos.sh .
CMD ["./ddos.sh"]