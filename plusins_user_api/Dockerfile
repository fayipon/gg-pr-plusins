FROM alpine:3.22

# Define the project name | 定义项目名称
ARG PROJECT=plusins_user_api
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=plusins_user_api.yaml
# Define the author | 定义作者
ARG AUTHOR="example@example.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY ./${PROJECT}_api ./
COPY ./etc/${CONFIG_FILE} ./etc/

EXPOSE 8082

ENTRYPOINT ["./plusins_user_api_api", "-f", "etc/plusins_user_api.yaml"]