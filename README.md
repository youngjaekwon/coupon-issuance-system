## 개요
캠페인 시작 시점에 맞춰 선착순으로 쿠폰을 발급하는 시스템.
## 요구사항 정리
### 1. 기능적 요구사항
- [ ] 캠페인을 생성할 수 있어야함
- [ ] 특정 캠페인의 정보 + 발급된 쿠폰 코드 목록을 조회할 수 있어야 함
- [ ] 쿠폰을 발급할 수 있어야 함
### 2. 비기능적 요구사항
- [ ] 캠페인 시작 직후 초당 500~1,000건의 요청 처리 가능
- [ ] 초과 발급 금지
- [ ] 발급된 쿠폰 코드 중복 금지
- [ ] 캠페인 시작 시각 전에는 절대 발급 불가
- [ ] 동시성 테스트 가능 도구 혹은 스크립트 제공
### 3. 기술적 요구사항
- [ ] Go 언어 기반
- [ ] ConnectRPC 사용 필수
## 시스템 아키텍처
```pgsql
                         +------------------+
                         |    Load Test     |
                         |   (Concurrency)  |
                         +--------+---------+
                                  |
                             HTTP (ConnectRPC)
                                  |
                        +---------v----------+
                        |   Coupon API       |
                        | (Gin + ConnectRPC) |
                        +----+--------+------+
                             |        |
               +-------------+        +------------------+
               |                                     |
    +----------v-----------+         +---------------v-------------+
    |      PostgreSQL      |         |            Redis            |
    |  (Campaign + Coupon) |         | (Prewarm + Stock Control)   |
    +----------------------+         +-----------------------------+

```
## 디렉토리 구조
```text
.
├── cmd/                   # Entry point (main.go)
├── internal/
│   ├── handler/           # ConnectRPC 핸들러
│   ├── repository/        # DB 및 Redis repository
│   ├── service/           # 비즈니스 로직
│   ├── util/              # 코드 생성기 등 유틸
│   ├── infra/             # DB, Redis 초기화
│   └── config/            # 환경 설정 로딩
├── gen/                   # ConnectRPC 코드 생성 위치
├── proto/                 # .proto 파일
├── tests/                 # 단위 및 통합 테스트
└── docker-compose.yml     # 개발 환경 구성
```
## API 명세

## 빌드 및 실행 방법
### 1. 환경 변수 설정
`.env`파일 생성:
```env
APP_ENV=development
PORT=8000
DB_DRIVER=postgres
DATABASE_URL="postgres://user:password@postgres:5432/mydb?sslmode=disable"
REDIS_ADDRESS=redis:6379
REDIS_PASSWORD=1d3fa16689e58a0daa720628ed51052e42f954f393a48cd66b6a4c0b9696215d
REDIS_DB=0
```
### 2. Docker Compose 실행
```bash
docker-compose up -d --build
```
## 테스트 방법
### 1. 단위 테스트
```bash
go test ./...
```
### 2. 동시성 부하 테스트
```bash
cd scripts/loadtest/
go run main.go
```