package http

import (
	stdhttp "net/http"
)
type stream struct{
	stdhttp.Request
}