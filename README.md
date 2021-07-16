# Formatter : mobigen/gologger

커맨드라인 환경에서 보기 쉽도록 포맷을 커스터마이징을 하였다.  
베이스 로그 플랫폼을 sirupsen/logrus 이다.  

## Configuration:

```go
type Formatter struct {
	// TimestampFormat - 시간 프린트 포맷 
	TimestampFormat string

	// ShowFullLevel - show a full level [WARNING] instead of [WARN]
	ShowFullLevel bool
	// NoUppercaseLevel - no upper case for level value
	NoUppercaseLevel bool

	// ShowFields
	ShowFields bool
	// NoFieldsSpace
	NoFieldsSpace bool
}
```

## Usage

이 모듈을 사용하기 위해서는 다음 과정이 필요하다.  
참고 : [golang private repository link!](https://mingrammer.com/go-modules-private-repo)
```bash
$ GOPRIVATE=github.com/mobigen go get -v github.com/mobigen/gologger
```

소스 코드  
```go
import (
	formatter "github.com/mobigen/gologger"
	"github.com/sirupsen/logrus"
)

log := logrus.New()
log.SetFormatter(&formatter.Formatter{
	TimestampFormat : "2000-01-02 15:04:05.000"
	ShowFields:    true,
})
log.SetLevel(logrus.DebugLevel)
log.SetReportCaller(true)

log.Info("just info message")
// Output : 2021-07-16 16:48:39.882 [INFO] [main.go          :  33] info message

log.WithField("component", "rest").Warn("warn message")
// Output : 2021-07-16 16:48:39.883 [WARN] [main.go          :  36] [ component:rest ] warn message
```

See more examples in the [tests](./tests/formatter_test.go) file.

