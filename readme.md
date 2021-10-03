# Unofficial Go SDK for GoPay Payments REST API

## Installation

```
go get https://github.com/apparently-studio/gopay-go-api
```

## Basic usage

```go
client := gopay.NewClient("my id", "my secret", "my goid", "gateway url")
```

Create Payment

```go
payment, err := gopay.CreatePayment(&client, gopay.PaymentBody{ /* define your payment */ })
if err != nil {
    log.Println(err)
    return
}

log.Println(payment.GWURL)
```

Get Payment

```go
payment, err := gopay.GetPayment(&client, int64(id))
if err != nil {
    log.Println(err)
    return
}

log.Println(payment)
```

Refund Payment

```go
err := gopay.RefundPayment(&client, int64(id), 50000)
if err != nil {
    log.Println(err)
    return
}

log.Println("Refund has been successful.")
```