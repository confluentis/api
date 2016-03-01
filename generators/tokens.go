package generators

import (
    "os/exec"
    "log"
)

func GenerateToken() string {
    out, err := exec.Command("uuidgen").Output()

    if err != nil {
        log.Fatal(err)
    }

    token := string(out[:])

    return token
}
