package microservices

import (
	// load the microserver drivers
	_ "github.com/dtm-labs/dtmdriver-dapr"
	_ "github.com/dtm-labs/dtmdriver-ego"
	_ "github.com/dtm-labs/dtmdriver-kratos"
	_ "github.com/dtm-labs/dtmdriver-polaris"
	_ "github.com/dtm-labs/dtmdriver-springcloud"
	_ "github.com/orcswang-lang/dtmdriver-gozero"
	_ "github.com/zhufuyi/dtmdriver-sponge"
)
