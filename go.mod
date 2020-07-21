module github.com/EDDYCJY/go-gin-example

go 1.14

require (
	github.com/astaxie/beego v1.12.2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.57.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/jinzhu/gorm v1.9.14
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => /Users/heyichen/Documents/go-gin/pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => /Users/heyichen/Documents/go-gin/middleware
	github.com/EDDYCJY/go-gin-example/models => /Users/heyichen/Documents/go-gin/models
	github.com/EDDYCJY/go-gin-example/pkg/e => /Users/heyichen/Documents/go-gin/pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => /Users/heyichen/Documents/go-gin/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => /Users/heyichen/Documents/go-gin/pkg/util
	github.com/EDDYCJY/go-gin-example/routers => /Users/heyichen/Documents/go-gin/routers

)
