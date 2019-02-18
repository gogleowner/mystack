# 도커
도커를 접하면서 처음 겪는 일들을 기록하는 일종의 cheat sheet.

## docker 설치후 실행시 이슈

```
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock
```
- docker가 root 계정으로 설치되었을 때 root가 아닌 계정으로 실행하고자 하면 발생하는 오류이다.
- docker의 실행 계정 속성을 변경후 유저 로그아웃후 재 로그인한다.
- `$ sudo usermod -a -G docker ${USER_NAME}`
- `$ sudo service docker restart`


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
- usage : `docker kill {container id}`
- `$ docker ps -a` 명령어 실행시에 컨테이너가 남아있음

## docker images
- 서버에 존재하는 도커 이미지 리스트

## docker exec
- Run a command in a running container
- usage : `$ docker exec [OPTIONS] CONTAINER COMMAND [ARGS....]`
- 예 : `$ docker exec -it 컨테이너ID /bin/bash`
	- 해당 컨테이너로 `/bin/bash` 를 실행한다.

### docker shell에 root로 접속
- `$ docker exec -u 0 -it {MY_CONTAINER} /bin/bash`
- 참고 : https://docs.docker.com/engine/reference/run/

```
root (id = 0) is the default user within a container. The image developer can create additional users. Those users are accessible by name. When passing a numeric ID, the user does not have to exist in the container.
```


## docker run
- usage : `$ docker run [OPTIONS] IMAGE[:TAG|@DIGEST] [COMMAND] [ARG...]`
- Options
	- `-d` : containers started in detached mode
	- `-p` : 포트포워딩
		- ex) `$ docker run -d -p 5000:5000 {image}`
		- `{image}` image를 detached모드로 실행하고 포트를 5000으로 포트포워딩 하라.

## docker rmi
- 도커 이미지 삭제
- usage : `$ docker rmi [OPTIONS] IMAGE [IMAGE...]`
- Options
	- `-f` : 실행중 혹은 종료된 container도 함께 삭제

## docker stats
- 도커 컨테이너들의 리소스 사용률 보기
- usage : `$ docker stats [OPTIONS] [CONTAINER...]`
- Options
	- `-a` : 모드 컨테이너

## docker top
- 도커 컨테이너에서 실행중인 프로세스 목록
- usage : `docker top CONTAINER [ps OPTIONS]`

## docker rm
- 도커 컨테이너 삭제
- usage : `$ docker rm [OPTIONS] CONTAINER [CONTAINER...]`
- Options
  - `-f, --force` : 컨테이너가 실행중이어도 강제 삭제
  - `-v, --volumes` : 컨테이너에 할당된 볼륨 영역 삭제
  - `-l, --link` : 특정 링크 삭제

## docker cp
- 도커 컨테이너 - 로컬 파일시스템 간의 파일/디렉토리 복사
- usage : `$ docker cp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-  /  $ docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH`
- Options
  - `--archive , -a` : Archive mode (copy all uid/gid information)
  - `--follow-link , -L` : Always follow symbol link in SRC\_PATH

## docker resource constraints
- docker 컨테이너에 리소스 사용량을 조절할 수 있다.
- [https://docs.docker.com/config/containers/resource_constraints/](https://docs.docker.com/config/containers/resource_constraints/)
