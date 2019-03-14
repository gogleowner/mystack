# JVM Execution Parameters

간단하게 앱에서 GC를 모니터링하고자 할때 썻던 JVM 실행 파라미터 정리

## GC log
- -verbose:gc
  - GC로그를 남김
- -Xloggc:{fileName}
  - GC로그를 파일에 남김
- -XX:+PrintGCTimeStamps
  - GC로그에 TimeStamp를 남김

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

