# 도커
도커를 접하면서 처음 겪는 일들을 기록하는 일종의 cheat sheet.

## docker ps
- 구동중인 컨테이너 목록
- `-a` : 구동중, 멈춘 컨테이너 목록 보기
- `--format` : ps내용을 원하는 형태로 포맷팅할 수 있음

	```
	CONTAINER ID / IMAGE / COMMAND / CREATED / STATUS / PORTS / NAMES
	```

## docker logs
- 컨테이너의 로그 보기
- usage : `docker logs [OPTIONS] CONTAINER`
- Options
	- `-f, --follow` : follow log output

## docker kill
- 구동중인 컨테이너를 shutdown

## docker images
- 서버에 존재하는 도커 이미지 리스트

## docker exec
- Run a command in a running container
- usage : `$ docker exec [OPTIONS] CONTAINER COMMAND [ARGS....]`
- 예 : `$ docker exec -it 컨테이너ID /bin/bash`
	- 해당 컨테이너로 `/bin/bash` 를 실행한다.

## docker run
- usage : `$ docker run [OPTIONS] IMAGE[:TAG|@DIGEST] [COMMAND] [ARG...]`
- Options
	- `-d` : containers started in detached mode