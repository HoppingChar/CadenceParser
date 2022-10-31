# CadenceParser

```
go run main.go
```

通过HTTP POST调用解析方法，例如：

```
POST http://127.0.0.1:8080/parse HTTP/1.1</br>

pub contract FlowToken {}
```

Response:
```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Mon, 31 Oct 2022 15:54:40 GMT
Content-Length: 428
Connection: close

{
  "program": {
    "Type": "Program",
    "Declarations": [
      {
        "Type": "CompositeDeclaration",
        "Access": "AccessPublic",
        "CompositeKind": "CompositeKindContract",
        "Identifier": {
          "Identifier": "FlowToken",
          "StartPos": {
            "Offset": 13,
            "Line": 1,
            "Column": 13
          },
          "EndPos": {
            "Offset": 21,
            "Line": 1,
            "Column": 21
          }
        },
        "Conformances": null,
        "Members": {
          "Declarations": null
        },
        "DocString": "",
        "StartPos": {
          "Offset": 0,
          "Line": 1,
          "Column": 0
        },
        "EndPos": {
          "Offset": 24,
          "Line": 1,
          "Column": 24
        }
      }
    ]
  }
}
```
