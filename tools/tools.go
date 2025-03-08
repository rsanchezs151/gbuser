package tools

import (
	"fmt"
	"time"
)

func FechaMySQL() string {
	// obtener fecha actual
	t := time.Now()
	//formatear fecha 2025-03-07
	return fmt.Sprintf("%d-%02d-%02dT%2d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
