FROM maven:3.6.3-adoptopenjdk-15
RUN addgroup -S spring && adduser -S spring -G spring
USER spring:spring
ARG JAR_FILE=target/*.jar
COPY ${JAR_FILE} app.jar
ENTRYPOINT ["java","-jar","/app.jar", "--allowed-namespaces=${ALLOWED_NAMESPACES}"]