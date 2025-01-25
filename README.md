# http-axe

1. **requests:**
  - GET(URL)
  - POST(URL, Content-Type+Encoding, bytesBody)
  - PUT(URL, bytesBody)
  - DELETE(URL, bytesBody)
  - HEAD(URL)
  - OPTIONS(URL, bytesBody)

2. **handlers:**
  - ServerStart(AddressPort)
  - HttpFunc(Endpoint, func(any) any)
