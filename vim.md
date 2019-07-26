# vim

## 여러 라인 수정

1. `Ctrl + V` 로 blockwise-visual mode 켜기
2. 수정하고자 하는 라인 만큼 블록
3. `Shift + i` 후 수정하고자 하는 텍스트 입력
4. `ESC` 버튼을 누르면 수정하고자 하는 라인까지 수정이 모두 됨

## 스크롤

- `Ctrl + D` : Scroll window Downwards in the buffer.
- `Ctrl + U` : Scroll window Upwards in the buffer.

## 문자열 치환

- `:%s/word/new_word/g` : 파일 전체의 word 문자열을 new\_word 로 변경한다.
- `:%s/word/new_word/gc` : 문자열 변경여부를 사용자에게 매번 묻는다.

