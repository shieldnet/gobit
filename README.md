# Go-Bit!
* *[UPBit](https://www.upbit.com/)* Auto Trading System with OpenAPI

# Features
## 구현 작업 진행상황
### 자산
* [ ] [전체 계좌 조회](https://docs.upbit.com/reference#%EC%A0%84%EC%B2%B4-%EA%B3%84%EC%A2%8C-%EC%A1%B0%ED%9A%8C)
### 주문
* [ ] [개별 주문 조회](https://docs.upbit.com/reference#%EA%B0%9C%EB%B3%84-%EC%A3%BC%EB%AC%B8-%EC%A1%B0%ED%9A%8C)
* [ ] [주문 리스트 조회](https://docs.upbit.com/reference#%EC%A3%BC%EB%AC%B8-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C)
* [ ] [주문 취소 접수](https://docs.upbit.com/reference#%EC%A3%BC%EB%AC%B8-%EC%B7%A8%EC%86%8C)
* [ ] [주문하기](https://docs.upbit.com/reference#%EC%A3%BC%EB%AC%B8%ED%95%98%EA%B8%B0)
### 출금
* [ ] [출금 리스트 조회](https://docs.upbit.com/reference#%EC%A0%84%EC%B2%B4-%EC%B6%9C%EA%B8%88-%EC%A1%B0%ED%9A%8C)
* [ ] [개별 출금 조회](https://docs.upbit.com/reference#%EA%B0%9C%EB%B3%84-%EC%B6%9C%EA%B8%88-%EC%A1%B0%ED%9A%8C)
* [ ] [출금 가능 정보](https://docs.upbit.com/reference#%EC%B6%9C%EA%B8%88-%EA%B0%80%EB%8A%A5-%EC%A0%95%EB%B3%B4)
* [ ] [코인 출금하기](https://docs.upbit.com/reference#%EC%BD%94%EC%9D%B8-%EC%B6%9C%EA%B8%88%ED%95%98%EA%B8%B0)
* [ ] [원화 출금하기](https://docs.upbit.com/reference#%EC%9B%90%ED%99%94-%EC%B6%9C%EA%B8%88%ED%95%98%EA%B8%B0)
### 입금
* [ ] [입금 리스트 조회](https://docs.upbit.com/reference#%EC%9E%85%EA%B8%88-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C)
* [ ] [개별 입금 조회](https://docs.upbit.com/reference#%EA%B0%9C%EB%B3%84-%EC%9E%85%EA%B8%88-%EC%A1%B0%ED%9A%8C)
* [ ] [입금 주소 생성 요청](https://docs.upbit.com/reference#%EC%9E%85%EA%B8%88-%EC%A3%BC%EC%86%8C-%EC%83%9D%EC%84%B1-%EC%9A%94%EC%B2%AD)
* [ ] [전체 입금 주소 조회](https://docs.upbit.com/reference#%EC%A0%84%EC%B2%B4-%EC%9E%85%EA%B8%88-%EC%A3%BC%EC%86%8C-%EC%A1%B0%ED%9A%8C)
* [ ] [개별 입금 주소 조회](https://docs.upbit.com/reference#%EA%B0%9C%EB%B3%84-%EC%9E%85%EA%B8%88-%EC%A4%8F-%EC%A1%B0%ED%9A%8C)
* [ ] [원화 입금하기](https://docs.upbit.com/reference#%EC%9B%90%ED%99%94-%EC%9E%85%EA%B8%88%ED%95%98%EA%B8%B0)

### 서비스 정보
* [ ] [입출금 현황](https://docs.upbit.com/reference#%EC%9E%85%EC%B6%9C%EA%B8%88-%ED%98%84%ED%99%A9)
* [ ] [API 키 리스트 조회](https://docs.upbit.com/reference#open-api-%ED%82%A4-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C)

### 시세 종목 조회
* [ ] [마켓 코드 조회](https://docs.upbit.com/reference#%EB%A7%88%EC%BC%93-%EC%BD%94%EB%93%9C-%EC%A1%B0%ED%9A%8C)

### 시세 캔들 조회
* [ ] [분 캔들 조회](https://docs.upbit.com/reference#%EB%B6%84minute-%EC%BA%94%EB%93%A4-1)
* [ ] [일 캔들 조회](https://docs.upbit.com/reference#%EC%9D%BCday-%EC%BA%94%EB%93%A4-1)
* [ ] [주 캔들 조회](https://docs.upbit.com/reference#%EC%A3%BCweek-%EC%BA%94%EB%93%A4-1)
* [ ] [월 캔들 조회](https://docs.upbit.com/reference#%EC%9B%94month-%EC%BA%94%EB%93%A4-1)

### 시세 체결 조회
* [ ] [최근 체결 내역](https://docs.upbit.com/reference#%EC%B5%9C%EA%B7%BC-%EC%B2%B4%EA%B2%B0-%EB%82%B4%EC%97%AD)

### 시세 Ticker 조회
* [ ] [현재가 정보](https://docs.upbit.com/reference#%EC%8B%9C%EC%84%B8-ticker-%EC%A1%B0%ED%9A%8C)

### 시세 호가 정보 (Orderbook) 조회
* [ ] [호가 정보 조회](https://docs.upbit.com/reference#%EC%8B%9C%EC%84%B8-%ED%98%B8%EA%B0%80-%EC%A0%95%EB%B3%B4orderbook-%EC%A1%B0%ED%9A%8C)
## How to use?
* `go`를 설치합니다. [Download and install go](https://golang.org/doc/install)
* `${GOPATH}/src/github.com/shieldnet` 디렉토리를 만들어줍니다.
  * FORK해서 사용하시는 경우, `${GOPATH}/src/github.com/<your_name>` 디렉토리를 만들어줍니다.
* 해당 디렉토리로 이동한 후, `git clone https://github.com/shieldnet/gobit.git`을 진행해줍니다.
* `cd gobit`을 통해 gobit 폴더로 이동해줍니다.
* `go get` 명령어를 실행해 `go.mod`의 패키지를 다운로드해줍니다.
* [Upbit OpenAPI 관리](https://upbit.com/mypage/open_api_management) 페이지에서 Open API Key를 발급받습니다.
  * SecretKey는 발급받은 직후가 아니면 다시 볼 수 없습니다. 주의하시기 바랍니다.
* `jwt-maker`폴더의 `keys.go.sample`에 발급받은 SecretKey와 AccessKey를 입력해 준 다음 `keys.go.sample`을 `keys.go`로 변경해줍니다.
* 전략을 본인의 취향에 맞게 수정합니다. (`strategy.go.sample` 파일을 `strategy.go`로 변경해준 다음 작업합니다.)
* `go run main.go`로 전략을 실행합니다.

## 사용 결과 화면
![image](https://user-images.githubusercontent.com/9548599/111020890-4e1e1380-840c-11eb-8c59-69141c5f7c9b.png)


## Caution(주의사항)
* 본 자동 매매 프로그램을 이용해서 손해가 발생하더라도, 저는 책임지지 않습니다.
* 본 자동 매매 프로그램을 이용하는 순간 위의 주의사항에 동의한 것으로 간주합니다.
* 배포시 소스코드 제공의무(Reciprocity)와 범위는, `gobit`에 포함된 모든 모듈과 소스코드에 해당합니다.
  * 단, 개인적으로 작성한 `strategy/strategy.go`와 `jwt-maker/keys.go`는 해당하지 않습니다.
  * 본 소프트웨어의 상업적, 다수를 대상으로 한 이윤 창출 목적의 이용을 금지합니다. 단, 개인 사용자의 개인 투자 목적으로는 사용 가능합니다. 

## Contribution
* 기본적으로, 개인 프로젝트이기 때문에 Contribution은 받지 않습니다만, 버그나 기능 추가에 관한 Issue는 환영하고 있습니다.
