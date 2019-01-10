# Memcached

# 데이터구조

## 캐시 항목을 찾기 위한 해시테이블

- bucket 배열
- 2^k - 1 를 통해 hash 마스크 처리
- hash_value & hash_mask 로 hash값이 들어있는 bucket 을 신속하게 찾아냄
- bucket들은 NULL로 끝나는 SingleLinkedList

## eviction 순서를 결정하는 LRU doubly linked list

- head → Most Recently Used
- tail → Least Recently Used
- eviction할때는 tail 부터 검사

## key, value, 플래그, 포인터 담고 있는 캐시 데이터 구조

- 해시테이블에서 bucket 당 single linked-list를 가리키는 포인터
- LRU에서 double linked-list에 사용되는 포인터
- reference counter : 캐시항목에 동시에 접근하는 쓰레드수
- 캐시 항목 상태 플래그
- 키
- 값
- 값 길이 (byte)

## 캐시 항목 데이터 관리자 slab allocator

- 캐시 항목에 대한 메모리 관리 기능
- 캐시 항목은 크기가 작아서 시스템 call(malloc/free)를 사용한 메모리 할당 / 해제는 속도가 느리고 thrashing 현상 발생 가능성이 있음

    기억장치의 페이지 부재가 비정상적으로 많이 발생하여 CPU가 프로그램 처리보다 페이지 교체에 더 많은 시간을 보내 성능히 급격히 줄어드는 현상. 멀티 프로세싱 기능을 갖춘 시스템에서 가상 메모리의 페이지 부재(Page Fault)가 너무 많이 발생하여 프로세스 실행보다 페이지 교체에 더 많은 시간을 소모하는 현상

- slab은 많은 항목들을 포함할 수 있는 큰 메모리 chunk
- 1,024 byte 메모리 chunk의 slab는 64바이트 이하 캐시 항목을 16개까지 저장 가능
- 캐시 항목이 접근될 때마다 slab allocator는 저장할 값 크기를 확인하고 수용할 수 있을 만큼의 큰 slab내의 캐시 항목을 돌려줌.

# Reference

- [https://d2.naver.com/helloworld/151047](https://d2.naver.com/helloworld/151047)
