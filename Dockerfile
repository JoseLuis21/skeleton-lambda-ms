# BUILDING STAGE #
FROM public.ecr.aws/docker/library/golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN cd cmd/ms-lambda && go build -o /ms-lambda

# RUNNER WITH LAMBDA IMAGE #
FROM public.ecr.aws/lambda/provided:al2 AS runner
COPY --from=builder /ms-lambda .

#LAMBDA INSIGHTS#
RUN curl -O https://lambda-insights-extension.s3-ap-northeast-1.amazonaws.com/amazon_linux/lambda-insights-extension.rpm && \
    rpm -U lambda-insights-extension.rpm && \
    rm -f lambda-insights-extension.rpm
#END OF LAMBDA INSIGHTS#

ENTRYPOINT ["./ms-lambda"]