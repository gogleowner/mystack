# macos에서 pyenv로 python 개발환경 구축

## pyenv 설치

```
$ brew install pyenv
$ vim zshconfig 수정
	 파일 내에 추가 export LC_ALL=en_US.UTF-8
(from https://github.com/pyenv/pyenv)
$ echo -e 'if command -v pyenv 1>/dev/null 2>&1; then\n  eval "$(pyenv init -)"\nfi' >> ~/.zshrc
$ source ~/.zshrc
```

## python 3.6.6 설치

```
$ pyenv install 3.6.6
$ pyenv global 3.6.6
$ pyenv rehash
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