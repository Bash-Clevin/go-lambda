FROM golang:1.20rc3-bullseye as dev
WORKDIR /app
COPY . /app
RUN go build -o go_lambda

FROM public.ecr.aws/lambda/go:1.2023.04.04.13
COPY --from=dev /app/go_lambda ${LAMBDA_TASK_ROOT}

CMD ["./go_lambda"]