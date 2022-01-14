package main

import (
    "archive/zip"
    "context"
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "path"
    "path/filepath"
    "runtime"
    "strings"
    "time"

    "github.com/pgaskin/kepubify/v4/kepub"
)

const UPLOAD_PATH_WINDOWS = "D:\\"
const UPLOAD_PATH_KOBO = "/mnt/onboard/kobofileserver"

var uploadPath string
var refreshScript string

type RequestData struct {
    converted bool
    fileName string
}

func responseString(msg string) string {
    return fmt.Sprintf("%s", msg)
}

func convertEPUB(converted bool, fileName string) (string, error) {
    if !strings.Contains(fileName, ".epub") {
        return fileName, nil
    }
    if strings.Contains(fileName, ".kepub.epub") {
        return fileName, nil
    }
    if !converted {
        return fileName, nil
    }

    converter := kepub.NewConverter()
    ctx := context.Background()

    r, err := zip.OpenReader(fileName)
    if err != nil {
        return "", err
    }
    defer func() {
        r.Close()
        os.Remove(fileName)
    }()

    newFileName := strings.Replace(fileName, ".epub", ".kepub.epub", -1)

    f, err := os.OpenFile(newFileName, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        return "", err
    }
    defer f.Close()

    err = converter.Convert(ctx, f, r)
    if err != nil {
        return "", err
    }

    return newFileName, nil
}

// You must add ExcludeSyncFolders settings to prevent from appearing books twice.
// /mmt/sd/kobofileserver and /mnt/onboard/kobofileserver
// ExcludeSyncFolders=(\\.(?!kobo|adobe).+|([^.][^/]*/)+\\..+|kobofileserver)
func notifyKoboRefresh() error {
    if runtime.GOOS != "windows" {
        cmd := exec.Command("/bin/sh", refreshScript)
        err := cmd.Run()
        if err != nil {
            return err
        }
    }
    return nil
}

func saveFile(r *http.Request) (RequestData, error) {
    var data RequestData

    if r.Method != "POST" {
        return data, fmt.Errorf("please use HTTP POST method to upload file")
    }

    err := r.ParseMultipartForm(32 << 20)
    if err != nil {
        return data, err
    }

    convertedStr := r.FormValue("upload-converted")
    converted := false
    if (convertedStr == "1") {
        converted = true
    }

    file, handler, err := r.FormFile("upload-file")
    if err != nil {
        return data, err
    }
    defer file.Close()

    fmt.Printf("uploadFile (%s) (%s) \n", handler.Filename, convertedStr)

    fileName := path.Join(uploadPath, handler.Filename)
    f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        return data, err
    }

    io.Copy(f, file)
    f.Close()

    data.converted = converted
    data.fileName = fileName

    return data, nil
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()

    requestData, err := saveFile(r)
    if err != nil {
        s := responseString(fmt.Sprintf("Error: (%v)", err))
        fmt.Fprintf(w, s)
        return
    }

    t2 := time.Now()

    finalFile, err := convertEPUB(requestData.converted, requestData.fileName)
    if err != nil {
        s := responseString(fmt.Sprintf("Error: (%v)", err))
        fmt.Fprintf(w, s)
        return
    }

    t3 := time.Now()

    /*
    // only for Elipsa now.
    // click connect button automatically.
    err = importBooks()
    if err != nil {
        s := responseString(fmt.Sprintf("Error: (%v), please use \"Import Books\" of NickelMenu", err))
        fmt.Fprintf(w, s)
        return
    }
    */

    s := responseString(
        fmt.Sprintf(
            "Uploading (%s) is successful, saveFile(%v), convertFile(%v)",
            finalFile,
            t2.Sub(t1),
            t3.Sub(t2),
        ),
    )
    fmt.Fprintf(w, s)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, homeHTML)
}

func main() {
    uploadPath = UPLOAD_PATH_KOBO
    if runtime.GOOS == "windows" {
        uploadPath = UPLOAD_PATH_WINDOWS
    }

	exePath, err := os.Executable()
    if err != nil {
        fmt.Println(err)
        return
    }

    exePath = filepath.Dir(exePath)

    refreshScript = path.Join(exePath, "refresh.sh")

    webPath := path.Join(exePath, "web")
    fs := http.FileServer(http.Dir(webPath))

    http.Handle("/web/", http.StripPrefix("/web/", fs))
    http.HandleFunc("/upload", uploadFile)
    http.HandleFunc("/", homePage)

    fmt.Printf("Listening on: 80, web path: (%s), uploading path: %s\n", webPath, uploadPath)

    err = http.ListenAndServe(":80", nil)
    if err != nil {
        fmt.Println(err)
        return
    }
}