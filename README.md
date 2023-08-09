# Service Notification / Messaging
This features uses as interface gateway to sending Email or whatsapp (Not Tested)

## Features
- Send Mail
  - Sending email using goMail
  - Create customValidator using validator go-playground
- Send Whatsapp
  - Using Internal thirdparty vendor
  - still implemented

## Environtment Variable

```env
CONFIG_SMTP_HOST="smtp.gmail.com"
CONFIG_SMTP_PORT="587"
CONFIG_AUTH_EMAIL="user@gmail.com"
CONFIG_AUTH_PASSWORD="yoursecretcode"

POSTGRES_HOST="localhost"
POSTGRES_PORT="5432"
POSTGRES_DBNAME="HServiceNotification"
POSTGRES_USER="root"
POSTGRES_PASSWORD="yourpassword"
POSTGRES_SSLMODE="disable"

BASE_URL="http://localhost"
BASE_PORT=":8181"
WEB_BASE_PORT=":3000"

GRPC_PORT=":5454"
```

## AUTHOR
- Budi Arifianto (arifianto.budi@jec.co.id)
