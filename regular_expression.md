# Regular Expression

## 전방탐색, 후방탐색

- 전방탐색 : `?=`

    - 예제문장 : `https://image.server.url/some-path-variable/20150914.jpg` 
    - `.jpg` 로 끝나는 파일에서 숫자를 가져오기 : `\d+(?=.jpg)`
    - 프로토콜 가져오기 : `\w+(?=:)`

- 후방탐색 : `?<=`

    - 예제문장 : `ABC01: $23.45`
    - $ 뒤의 숫자 가져오기 : `(?<=\$)[0-9.]+`

