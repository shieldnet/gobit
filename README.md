# Go-Bit!
* *[UPBit](https://www.upbit.com/)* Auto Trading System with OpenAPI

# Features
* 현재 구현된 기능은 아래와 같습니다.
  * 시장가 매도/매수
  * 분봉 차트를 가져오는 기능
  * jwt-token 생성기

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
