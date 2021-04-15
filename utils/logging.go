package utils

import (
  "io"
  "log"
  "os"
)

func LoggingSettings(logFile string) {

  // 読み、書き、追記
  logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

  if err != nil {
    log.Fatalln(err)
  }

  // ログの書き込み先を標準出力とログファイルに指定
  multiLogFile := io.MultiWriter(os.Stdout, logfile)

  // ログフォーマットの指定
  log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

  log.SetOutput(multiLogFile)
}