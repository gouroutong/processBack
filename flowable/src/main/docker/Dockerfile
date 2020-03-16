#FROM anapsix/alpine-java
#ADD target/flowable-0.0.1-SNAPSHOT.jar flowable.jar
#ENTRYPOINT ["java","-jar","/flowable.jar"]

FROM anapsix/alpine-java
VOLUME /tmp
ADD flowable-1.0.2.jar app.jar
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/app.jar"]