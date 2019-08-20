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
- `-a, --all` : 구동중, 멈춘 컨테이너 목록 보기
- `-q, --quiet` : 숫자 ID 만 출력

```
$ docker ps -q
1c07723538f7
692a0d927d6e
da1b1752b611
```
- `-s, --size` : 파일의 사이즈를 함께 출력

```
$ docker ps -s
CONTAINER ID    IMAGE      COMMAND       CREATED        STATUS                PORTS              NAMES            SIZE
1c07723538f7    blabla   "/bin/bash"   23 hours ago    Up 23 hours     0.0.0.0:8080->8080/tcp    blabla   912.1 kB (virtual 1.162 GB)
```

- `--format` : ps내용을 원하는 형태로 포맷팅할 수 있음
- `-f, --filter` : 특정 조건에 맞는 컨테이너 목록을 필터링. `key=value` 형태로 필터링하면 된다.
  - `$ docker ps -f status=exited` : 종료된 컨테이너 목록
  - `$ docker ps -f name={blabla}` : 컨테이너 이름으로 필터링

	```
	CONTAINER ID / IMAGE / COMMAND / CREATED / STATUS / PORTS / NAMES
	```

## docker logs
- 컨테이너의 로그 보기
- usage : `docker logs [OPTIONS] CONTAINER`
- Options
	- `-f, --follow` : follow log output
  - `--since` : 특정 시점부터의 로그를 보여준다.
    - `$ docker logs {CONTAINER ID} --since 2013-01-02T13:23:37`
    - `$ docker logs {CONTAINER ID} --since 42m`

## docker kill
- 구동중인 컨테이너를 shutdown
- usage : `$ docker kill {container id}`
- `$ docker ps -a` 명령어 실행시에 컨테이너가 남아있음

## docker images
- 서버에 존재하는 도커 이미지 리스트
- Usage : docker images [OPTIONS] [REPOSITORY[:TAG]]
- Options
	- `--all -a` : Show all images (default hides intermediate images)
	- `--quiet , -q` : Only show numeric IDs
	- ...


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
		
  - `--name` : 컨테이너에 이름 지정
  - `--hostname , -h` : 컨테이너에 호스트이름 설정
  - `--link` : 컨테이너 끼리 연결
    - ex) `--link="container_name:alias"`
  - `--rm` : Automatically remove the container when it exits
    - 컨테이너가 종료될때 자동으로 데이터를 지움
  - `--env , -e` : 환경변수 세팅
    - ex) `-e "SPRING_PROFILES_ACTIVE=blabla"`

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
- usage : `$ docker top CONTAINER [ps OPTIONS]`

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

## 모든 docker 컨테이너나 이미지 stop or remove
- `$ docker stop $(docker ps -a -q)`
- `$ docker rm $(docker ps -a -q)`
- `$ docker rmi $(docker images -a -q)`

## 종료된 docker 컨테이너들의 볼륨 삭제
- `$ docker rm -v $(docker ps -a -q -f status=exited)`

## docker volume ls
- Volume의 목록을 보여준다.
- Usage `$ docker volume ls [OPTIONS]`
- Options
  - `--filter, -f` : Provide filter values (e.g. ‘dangling=true’)
  - `--format` : Pretty-print volumes using a Go template
  - `--quiet, -q` : Only display volume names
- Volume : 도커 컨테이너에 의해 생성 및 사용되는 데이터를 유지하는 메커니즘
  ![types-of-mounts-volume](https://docs.docker.com/storage/images/types-of-mounts-volume.png)
  - https://docs.docker.com/storage/volumes/

## docker build
- Dockerfile로부터 도커 이미지를 빌드한다.
- Usage : `$ docker build [OPTIONS] PATH | URL | -`
- Options
  - `--rm true` : 이미지 빌드가 성공한 후 임시 컨테이너를 삭제한다.
  - `-t --tag=""` : 이미지의 이름과 태그를 명시한다. ex) `abc-project:0.0.1`
  - `--build-arg` : 빌드 시점에 Dockerfile 에 `ENV` 값을 세팅한다.
  - `--file , -f` : 빌드할 Dockerfile 파일이름 지정 (Default is ‘PATH/Dockerfile’)
    - `$ docker build -f {PATH/your-dockerfile-name}`

## docker system prune
- 중지된 컨테이너, dangling 이미지, 사용하지 않는 네트워크를 삭제한다.

```
$ docker system prune
WARNING! This will remove:
  - all stopped containers
  - all networks not used by at least one container
  - all dangling images
  - all dangling build cache

Are you sure you want to continue? [y/N] y
Deleted Containers:
3174e566b9122d25394a8d4c8fd4e971bace5e29bb75ff6e036e5dc9fcb54bae
...

Deleted Networks:
..

Deleted Images:
..

Total reclaimed space: 3.294GB
```

- 사용하지 않는 volume을 제거하려면 `--volumes` 옵션을 추가한다.

```
$ docker system prune --volumes
WARNING! This will remove:
  - all stopped containers
  - all networks not used by at least one container
  - all volumes not used by at least one container
  - all dangling images
  - all dangling build cache

Are you sure you want to continue? [y/N] y
Deleted Volumes:
..

Total reclaimed space: 74.45MB
```

## References
- [Docker 한글 문서 / 영상 모음집](http://documents.docker.co.kr)
- [Top 10 Docker CLI commands you can’t live without](https://medium.com/the-code-review/top-10-docker-commands-you-cant-live-without-54fb6377f481)
- https://blog.docker.com/2019/07/intro-guide-to-dockerfile-best-practices/

