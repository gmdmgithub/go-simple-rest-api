#First GO REST API

## Quick Start

### Write main.go
``` bash
# simply run
go run ./main.go
# simplify run - the project: 
go build
```

### Install mux router
``` bash
go get -u github.com/gorilla/mux
```
### Imports
import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
