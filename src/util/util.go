package util

import (
       "time"
)

func GetCurrentSecond() int64{
     return time.Now().Unix()
}