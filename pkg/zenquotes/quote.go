package zenquotes

import (
	"fmt"
)

type Quote struct {
	Q string `json:"q,omitempty"`
	A string `json:"a,omitempty"`
	H string `json:"h,omitempty"`
}

func (q Quote) String() string {
	return fmt.Sprintf("%s (c)%s", q.Q, q.A)
}
