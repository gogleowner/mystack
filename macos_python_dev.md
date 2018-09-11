# macos에서 pyenv로 python 개발환경 구축

## pyenv 설치

```
$ brew install pyenv
$ vi ~/.zshrc
	 파일 내에 추가 export LC_ALL=en_US.UTF-8
   (from https://github.com/pyenv/pyenv)
$ echo -e 'if command -v pyenv 1>/dev/null 2>&1; then\n  eval "$(pyenv init -)"\nfi' >> ~/.zshrc
$ source ~/.zshrc
```

## pyenv를 이용하여 python v3.6.6 설치

```
$ pyenv install 3.6.6 // python v3.6.6 설치
$ pyenv global 3.6.6 // global 설정을 v3.6.6 으로 설정 
$ pyenv rehash // 최초에 global로 버전 설정시 해야함. 그 이후에는 안해도 됨.
```

- 예시

  ```
  $ pyenv global 3.6.6 && pyenv rehash && python --version
  Python 3.6.6
  $ pyenv global system && python --version // macos 기본 버전은 system
  Python 2.7.10
  ```

- pyenv 기본 커멘드
  - 설치된 python 버전목록
    ```
    $ pyenv versions
      system
    * 3.6.6 (set by /Users/{YOUR_USER_DIR/.pyenv/version)
    ```

## pipenv 설치

```
$ pip install pipenv
$ pipenv install
$ pipenv lock --pre
// To activate this project's virtualenv, run the following:
$ pipenv shell 
Loading .env environment variables…
Launching subshell in virtual environment. Type 'exit' or 'Ctrl+D' to return.
```
