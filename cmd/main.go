package main

import (
    "archive/zip"
    "context"
    "fmt"
    "io"
    "io/fs"
    "net/http"
    "os"
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

type EnvSettings struct {
    uploadPath string
    downloadPath string
}

type UploadResult struct {
    Result string
    FileName string
    SavedTime string
    ConvertedTime string
}

func responseString(res UploadResult) string {
    // prevent from having error, just use fmt.
    return fmt.Sprintf(
        `{"Result":"%s","FileName":"%s","SavedTime":"%s","ConvertedTime":"%s"}`,
        res.Result,
        res.FileName,
        res.SavedTime,
        res.ConvertedTime,
    )
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
    var res UploadResult

    t1 := time.Now()

    requestData, err := saveFile(r)
    if err != nil {
        res.Result = fmt.Sprintf("Error: (%v)", err);
        s := responseString(res)
        fmt.Fprintf(w, s)
        return
    }

    t2 := time.Now()

    finalFile, err := convertEPUB(requestData.converted, requestData.fileName)
    if err != nil {
        res.Result = fmt.Sprintf("Error: (%v)", err);
        s := responseString(res)
        fmt.Fprintf(w, s)
        return
    }

    t3 := time.Now()

    res.FileName = finalFile
    res.SavedTime = t2.Sub(t1).String()
    res.ConvertedTime = t3.Sub(t2).String()

    s := responseString(res)
    fmt.Fprintf(w, s)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, homeHTML)
}

func checkAndMkdir(path string) (error) {
    _, err := os.Stat(path)
    if err == nil { return nil }

    if !os.IsNotExist(err) { return err }

    fmt.Println(path)

    err = os.Mkdir(path, fs.ModePerm)
    if err != nil { return err }

    return nil
}

func prepareEnv(env EnvSettings) error {
    err := checkAndMkdir(env.uploadPath)
    if err != nil { return err }

    err = checkAndMkdir(env.downloadPath)
    if err != nil { return err }

    return nil
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

    img := path.Join(exePath, "qrcode.png")
    url, err := generateQRCode(img)

    refreshScript = path.Join(exePath, "refresh.sh")

    webPath := path.Join(exePath, "web")
    fs := http.FileServer(http.Dir(webPath))

    // for downloading files from device.
    downloadPath := path.Join(exePath, "download")
    dlFs := http.FileServer(http.Dir(downloadPath))

    envSettings := EnvSettings {
        uploadPath: uploadPath,
        downloadPath: downloadPath,
    }

    err = prepareEnv(envSettings)
    if err != nil {
        fmt.Println(err)
        return
    }

    http.Handle("/web/", http.StripPrefix("/web/", fs))
    http.Handle("/download/", http.StripPrefix("/download/", dlFs))
    http.HandleFunc("/upload", uploadFile)
    http.HandleFunc("/", homePage)

    fmt.Printf("Listening on: %s:80, web path: (%s), uploading path: %s\n", url, webPath, uploadPath)

    err = http.ListenAndServe(":80", nil)
    if err != nil {
        fmt.Println(err)
        return
    }
}