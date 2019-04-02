# JVM Execution Parameters

간단하게 앱에서 GC를 모니터링하고자 할때 썻던 JVM 실행 파라미터 정리

## GC log
- -verbose:gc
  - GC로그를 남김
- -Xloggc:{fileName}
  - GC로그를 파일에 남김
- -XX:+PrintGCTimeStamps
  - GC로그에 java 앱의 구동시간 기준의 time stamp값을 남김
- -XX:+PrintGCDateStamps
  - GC로그에 날짜에 대한 date stamp를 남김
- -XX:+PrintGCDetails
  - GC로그에 GC 전후에 대한 각 영역별 용량 변화를 남김
- -XX:+PrintHeapAtGC
  - GC 발생 전후의 Heap에 대한 정보를 상세하게 기록

## Heap Dump
- -XX:+HeapDumpBeforeFullGC
  - Full GC 발생 전 Heap Dump 파일을 남김
- -XX:+HeapDumpAfterFullGC
  - Full GC 발생 후 Heap Dump 파일을 남김
- -XX:+HeapDumpOnOutOfMemoryError
  - OOM 에러 발생시 Heap Dump 파일을 남김

## Heap Size
- -Xms{sizeOfInitialHeapSize}
- -Xms{sizeOfMaxHeapSize}
- -XX:MetaspaceSize={sizeOfInitialMetaspaceSize}

