# ivona-service
A simple micro service around Ivona TTS API


Endpoints
----------
   

- `/` [*POST*]
    - param **text**
        - text to TTS
    - param **voice**
        - voice to be used
    - response : converted audio stream with content-type 'audio/mpeg3'

- `/voices/` [*POST*]
    - param **name**
        - optional name to filter
    - param **language**
        - optional language to filter
    - param **gender**
        - optional gender to filter
    - response : list of voices with content-type 'application/json'
```json
{
  "Voices": [
        { 
            "Name": "string",
            "Language": "string",
            "Gender": "String"
        },
        { 
            "Name": "string",
            "Language": "string",
            "Gender": "String"
        }
    ],
  "RequestID": "ivona-request-id-string",
  "ContentType": "content type sent by ivona"
}
```


MIT License

