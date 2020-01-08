# AWS Reinvent

# Infrastructure Service Update (Compute, Storage, Networking, Container)

- AWS Reinvent 2019 발표 자료 : [https://aws.amazon.com/ko/new/reinvent/](https://aws.amazon.com/ko/new/reinvent/)
- AWS : 광범위하고 깊이 있는 클라우드 플랫폼들을 제공한다.
    - 기술 및 비즈니스
    - 마켓 플레이스
    - 분석, 웹서비스, DevOps, 모바일 서비스, IoT, 인공지능, 엔터프라이즈, 하이브리드 환경, 마이그레이션
    - 인프라, 본 서비스 (컴퓨팅, 데이텁ㅔ이스, CDN, 스토리지, 네트워킹), 보안 & 규정
- AWS 컴퓨팅 가상화 역사
    - Xen → Nitro → 2015년 annaprunalabs을 이수하여 Nitro system 준비.
    - Nitro
        - Nitro card : VPC 네트워킹, Amazon EBS, Instance Stroage
        - Nitro Security Chip : 메인보드에 통합, 하드웨어 자원보호, H/W Root of Trust
        - Nitro Hypervisor : 경량 하이퍼바이저 메모리 / CPU 할당, 베어메탈 급 성능
    - [http://www.annapurnalabs.com/](http://www.annapurnalabs.com/) 인수.

## 컴퓨팅 서비스 업데이트

- 약 270 이상의 인스턴스 타입을 제공

### EC2 인스턴스 타입을 다양하게 제공

- C → 컴퓨팅 특화
- R → 메모리 특화
- M → General Purpose
- G → Graphic 특화

### Amazon EC2 Inf1 Instance 출시

- 클라우드 상에서 가장 빠르고 저렴한 머신러닝 추론 성능 제공
- AWS Inferentia
- Inf1은 GPU로 구동되는 G4인스턴스와 비교하여 최대 3배 높은 처리량, 최대 40%의 낮은 추론 당 비용 제공

### AWA Neuron

- Inferentia에서 고성능 딥러닝을 가능하게 하는 S/W Suite 제공 ([github.com/aws/aws-neuron-sdk](http://github.com/aws/aws-neuron-sdk))

### AWAS Compute Optimizer

기계학습 기반 추천 엔진을 사용하여 워크로드에 적합한 최적의 Amazon EC2 인스턴스 및 Auto Scaling 그룹 식별 가능

- 워크로드에 가장 적합한 AWS 리소스 쉽게 식별, AWS 워크로드에 대한 비용 절감 및 성능 향상
- M,C,R,T 및 X 인스턴스 제품군에서 140개 이상의 EC2 인스턴스 유형 옵션 지원

### AWS Graviton2 Processor

- 클라우드 기반 ARM 기반 프로세서 : 저전력 프로세서. 직접도가 높고 비용이 저하됨.
- M6g, C6g, R6g

### AWS Nitro Enclaves

- EC2 인스턴스 내에서 보안에 민감한 데이터를 추가로 보호하기 위해 추가 격리 구역 생성 가능

### Amazon EC2 Image Builder

- Amazon Machine Image 구축 및 유지 위한 서비스
- VM 이미지 생성 자동화 / 필수 보안 구성 요소로 이미지 생성 및 보안 패치 자동화 / 프로덕션 적용 전 테스팅 손쉽게 수행.
- [https://aws.amazon.com/ko/image-builder/](https://aws.amazon.com/ko/image-builder/)

### Amazon Braket

- 과학자와 개발자가 양자컴퓨팅을 쉽게 탐색하고 실험할 수 있는 완전 관리형 서비스
- 양자 알고리즘 설계&테스트&실행 단일 환경 / 다양한 양자 하드웨어 기술 실험 / 하이브리드 양자 및 클래식 알고리즘 실행 / 양자 컴퓨팅 전문가와의 도움 제공
- [https://aws.amazon.com/ko/blogs/korea/amazon-braket-get-started-with-quantum-computing/](https://aws.amazon.com/ko/blogs/korea/amazon-braket-get-started-with-quantum-computing/)

### AWS Outposts

- AWS 인프라 및 주요 서비스, API 및 개발 도구를 고객 온프레미스로 확장하는 완전 관리형 서비스
- [https://aws.amazon.com/ko/outposts/](https://aws.amazon.com/ko/outposts/)

### AWS Wavelength

- 5G 네트워크 엣지에서 AWS 컴퓨팅 및 스토리지를 사용하여 5G 기반 모바일 디바이스 및 사용자에게 저 지연 APP에 최적 성능 제공 가능.
- AWS 리전과 통신사 5G 망과 직접 연결, 로컬 컴퓨팅&스토리지&DB 및 기타 서비스 제공, 5G기반 새로운 모바일 앱 경험 제공

## 스토리지 서비스 업데이트

- Amazon S3 : object 스토리지
- Amazon EBS : Block level 스토리지
- Amazon EFS , Amazon FSx for Windows File Server , Amazon FSx for Lustre : 파일 스토리지
- [https://aws.amazon.com/ko/products/storage/](https://aws.amazon.com/ko/products/storage/)
- AWS 스토리지 서비스의 중심은 S3

### New S3 Access Points

- 단일 S3버킷에 있는 데이터 셋에 대한 공유를 더욱 손쉽게 제어
- 동일 S3버킷에 저장된 데이터에 대해 사용자, 앱 별 엔드포인트 제공. 다양한 App들이 공유데이터를 사용하는 데이터 레이크
- 동일 리전이면 복제.
    - 동일 리전에서 데이터에 대한 자동화된 비동기 복제 지원
- Amazon S3 Replication Time Control : 99.99%의 데이터를 15분 이내에 복제

### EBS Fast Snapshot Restore (FSR)

- 6x 낮아진 RTO(Recovery time objective)
- 정기적인 스냅샷을 생성. 대략 60분... 최대 10개의 볼륨까지 한꺼번에 복구. FSR은 스냅샷 생성 전,중,후 언제나 적용 가능.

## 네트워킹 서비스 업데이트

- 리전, 가용존 (AZ), VPC
- AWS 글로벌 네트워크로 리전간에 연결이 되어있다.

### Ingress Routing

- 인입되는 트래픽을 최종 목적지 도달 전  3rd party 어플라이언스 또는 AWS 서비스로 라우팅.
- 클라우드에서는 L3 레이어를 통해 게이트웨이 역할을 할 수 밖에 없었는데 이번에 Ingress 기반 라우팅을 제공하게 됨.
- Ingress route table → NGFW (Next Generation Firewall) → APP

### AWS Transit Gateway Inter-Region Peering

- 멀티 AWS 리전 기반 네트워크 게이트웨이를 연결하여 글로벌 네트워크 구축 가능
- 리전 이중화를 위해 멀티 리전간 리소스 공유, 데이터 복제
- 단일 장애 지점이나 대역폭 병목 현상 없이 리전 간 트래픽 암호화
- 허브 & Spoke 모델의 네트워크를 여러 AWS 리전에서 사용 가능케함.
- AWS 글로벌 네트워크를 통한 프라이빗, 고성능 연결 구성으로 Scale out이 가능하게 함.

### AWS Accelerated Site-to-Site VPN

### Global Network 모니터링

- 리전간의 병목 등 모니터링.

### Multicast on AWS Transit Gateway

- 멀티케스트 : 여러 그룹의 도메인에 동시에 데이터를 전송할 수 있게함.

## 컨테이너 서비스 업데이트

### AWS Cloud 컨테이너 옵션

- Non AWS Cloud 흐름
    - Docker Host
    - 쿠버네티스이 사용률이 높아짐. : 서비스가 많아지면 노드 관리가 어려워진다.
- AWS Cloud 흐름
    - 2012 Amazon ECS
    - 2017 Fargate
    - 2018 Amazon EKS
    - 2019 Managed Node Groups
    - 컨테이너 로드맵 : [https://github.com/aws/containers-roadmap/projects/1](https://github.com/aws/containers-roadmap/projects/1)

### AWS Cloud 컨테이너 구성 옵션

- 관리 : ECS, EKS (Elastic Kubernetes Service)
    - ECS : AWS 전용 컨테이너 오케스트레이션 서비스
    - EKS : 쿠버네티스를 직접 구성 및 운영해야하는 복잡성 감소.
- 호스팅 : EC2, Fargate
- 이미지 보관소 : Amazon Elastic Container Registry

### AWS 컨테이너 서비스 주요 기능 업데이트

- Fargate Spot : Fargate 표준 가격에서 최대 70% 할인된 예비 용량.
    - APP이 stateless 한 경우 사용할 수 있음.
- Amazon ECS Capacity Providers

### Amazon ECS를 위한 Firelens

- 컨테이너를 위한 로깅 시스템
- 컨테이너 별로 발생하는 로그를 S3 등으로 전송

# Serverless and Application Service

## Serverless and Application

- 과거부터 현재까지의 어플리케이션 개발 흐름
    - 물리 서버 → virtualization → containerization → serverless
    - 인프라 레이어는 점점 추상화되고, 비즈니스 로직 작성에 집중할 수 있는 환경으로 발전하고 있다.
- 서버리스 운영 모델의 이점
    - 서버 운영 및 관리할 필요 X
    - 자동 스케일링
    - 쓴 만큼만 비용 청구
    - 고가용성 & 보안
- 기본 아키텍처
    - Event Source
        - Chanages in data state
        - Requests to Endpoints
        - Changes in resource state
    - Function : lambda 함수
        - 여러 언어로 작성.
    - Services
- Circuit breaker and many more by jeremy day
- 인프라 관리 - 코드로 관리.
    - AWS CloudFormation : YAML, JSON 등으로 정의.
    - Serverless Application Model : 코드 형태로 좀더 간소화하게
    - AWS Cloud Development Kit : 각 언어로 작성후, 인프라는 sdk 를 통해 코드로 정의
- 배포 전략 : 서버리스 CI CD 전략.

## 서버리스 서비스 업데이트

### Lambda - Provisioned Concurrency

- cold-start 문제를 해결해달라는 요청이 많았다. 일정 개수의 람다를 미리 띄워놓는 형태로 서비스를 활성화한다. 일정 개수가 돌아가게 해야하는 요금은 추가로 부과된다.
- 함수 시작 시간을 밀리초 단위로 유지하면서 기능을 초기화하고 하이퍼-레디 상태로 유지.
- 개발자가 프로비저닝된 동시성 설정 권한 제어 가능

### API Gateway

- REST & WebSocket 지원
- Flexible auth options
- Throttling , usage tiers
- caching

### HTTP APIs for API Gateway

- 기존엔 REST API만 제공했었는데, API 프록시 역할로만 사용하는 용도로 쓸 수 있다.

### AWS Step Functions

- 실패 작업 재실행, 순차 실행, 예외처리, 데이터 기반 처리, 병렬 실행.
- 생성 방법 : JSON 으로 워크플로우 정의, 콘솔에서 시각화 제공, 실행 상태 모니터링
- Express Workflow : 대용량 데이터 처리를 단기간에 수행하는 워크플로우 작성시 사용
    - 초당 10만건 이상 처리 가능.

### Amazon EventBridge

- 다양한 소스의 데이터와 APP을 쉽게 연결하는 서버리스 이벤트 버스

## 개발 관련 업데이트

- AWS 개발자들의 노하우  [https://aws.amazon.com/ko/builders-library](https://aws.amazon.com/ko/builders-library)
- AWS Amplify Overview : 클라우드 기반 웹앱 및 모바일앱을 개발
    - Amplify Framework에는 라이브러리, CLI 툴체인, UI 구성요소가 포함. 이를 통해 AppSync, Poinpoint와 같은 AWS 서비스들과 통합.
- Amplify for iOS & Android ; MOB317
- Amplify Datastore
    - GraphQL을 사용하여 모바일, 웹앱과 클라우드간 데이터를 자동으로 동기화하는 디바이스 기반의 영구스토리지 엔진
- Amplify Prediction : AI/ML 경험이 없어도 AL/ML 서비스 쉽게 활용
    - 말하기, 번역, OCR, 이미지 객체정보 수집.

