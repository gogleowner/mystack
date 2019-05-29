# Netty

Netty는 비동기 이벤트 네트워크 애플리케이션 프레임워크다.

## 소개
TBD ...

## 동작방식
TBD ...

## Worker Thread
- 실제 동작을 수행하는 쓰레드 그룹
- 디폴트 워커 개수는 `코어수 * 2` 이다.
  - https://github.com/netty/netty/blob/00afb19d7a37de21b35ce4f6cb3fa7f74809f2ab/transport/src/main/java/io/netty/channel/MultithreadEventLoopGroup.java#L41

# 만날수 있는 오류
- io.netty.channel.unix.Errors$NativeIoException: writeAddress(..) failed: Connection reset by peer
  - 발생 위치 클래스&메소드 : https://github.com/netty/netty/blob/f17bfd0f64189d91302fbdd15103788bf9eabaa2/transport-native-unix-common/src/main/java/io/netty/channel/unix/Errors.java#L72
  - Connection reset by peer 오류는, client 가 요청한 것을 서버에서 응답하려고 했으나 client의 브라우저가 닫히거나 다른 페이지로 이동했을때 발생한 오류이다.
  - 사내 개발서버에 올라가있는 앱이 health check 할때 간헐적으로 발생하고 있는 오류인데, 정확한 원인을 잘 모르겠다. 코딩으로 해결하기 보다는 인프라적으로 해결해야할 것 같은 느낌이 든다.


# References
- https://sungjk.github.io/2016/11/08/NettyThread.html
- https://gompangs.tistory.com/105
- https://okky.kr/article/228975

