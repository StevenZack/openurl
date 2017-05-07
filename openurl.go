package openurl
import (
    "errors"
    "os/exec"
    "runtime"
)
var unsupportedPlatformError = errors.New("Unsupported platform")
var invalidUrlError = errors.New("Invalid url")
func Open(url string) error {
    if len(url) == 0 {
        return invalidUrlError
    }
    switch runtime.GOOS {
    case "windows":
        if _, err := os.Stat("\"C:/ProgramData/Microsoft/Windows/Start Menu/Programs/Google Chrome\""); !os.IsNotExist(err) { //Chrome  installed on this computer
            file, _ := os.Create("./openurl.bat")
            file.WriteString("\"C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\Google Chrome.lnk\" --app=http://" + url)
            file.Close()
            cmd := exec.Command(".\\openurl.bat")
            if err := cmd.Run(); err != nil {
                return err
            }
        } else { //no chrome
            exec.Command("explorer", url).Run()
        }
    case "darwin":
        if err := exec.Command("google-chrome", "--app=http://"+url); err != nil {
            if err := exec.Command("google-chrome-stable", "--app=http://"+url); err != nil {
                if err := exec.Command("chromium", "--app=http://"+url); err != nil {
                    return exec.Command("open", url)
                }
            }
        }
    case "linux":
        if err := exec.Command("google-chrome", "--app=http://"+url); err != nil {
            if err := exec.Command("google-chrome-stable", "--app=http://"+url); err != nil {
                if err := exec.Command("chromium", "--app=http://"+url); err != nil {
                    return exec.Command("xdg-open", url)
                }
            }
        }
    default:
        return "", unsupportedPlatformError
    }
