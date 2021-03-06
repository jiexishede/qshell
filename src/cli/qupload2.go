package cli

import (
	"flag"
	"fmt"
	"github.com/qiniu/log"
	"os"
	"qshell"
)

func QiniuUpload2(cmd string, params ...string) {
	flagSet := flag.NewFlagSet("qupload2", flag.ExitOnError)

	var threadCount int64
	var srcDir string
	var fileList string
	var bucket string
	var putThreshold int64
	var keyPrefix string
	var ignoreDir bool
	var overwrite bool
	var checkExists bool
	var checkHash bool
	var checkSize bool
	var skipFilePrefixes string
	var skipPathPrefixes string
	var skipFixedStrings string
	var skipSuffixes string
	var upHost string
	var bindUpIp string
	var bindRsIp string
	var bindNicIp string
	var rescanLocal bool
	var logLevel string
	var logFile string

	flagSet.Int64Var(&threadCount, "thread-count", 0, "multiple thread count")
	flagSet.StringVar(&srcDir, "src-dir", "", "src dir to upload")
	flagSet.StringVar(&fileList, "file-list", "", "file list to upload")
	flagSet.StringVar(&bucket, "bucket", "", "bucket")
	flagSet.Int64Var(&putThreshold, "put-threshold", 0, "chunk upload threshold")
	flagSet.StringVar(&keyPrefix, "key-prefix", "", "key prefix prepended to dest file key")
	flagSet.BoolVar(&ignoreDir, "ignore-dir", false, "ignore the dir in the dest file key")
	flagSet.BoolVar(&overwrite, "overwrite", false, "overwrite the file of same key in bucket")
	flagSet.BoolVar(&checkExists, "check-exists", false, "check file key whether in bucket before upload")
	flagSet.BoolVar(&checkHash, "check-hash", false, "check hash")
	flagSet.BoolVar(&checkSize, "check-size", false, "check file size")
	flagSet.StringVar(&skipFilePrefixes, "skip-file-prefixes", "", "skip files with these file prefixes")
	flagSet.StringVar(&skipPathPrefixes, "skip-path-prefixes", "", "skip files with these relative path prefixes")
	flagSet.StringVar(&skipFixedStrings, "skip-fixed-strings", "", "skip files with the fixed string in the name")
	flagSet.StringVar(&skipSuffixes, "skip-suffixes", "", "skip files with these suffixes")
	flagSet.StringVar(&upHost, "up-host", "", "upload host")
	flagSet.StringVar(&bindUpIp, "bind-up-ip", "", "upload host ip to bind")
	flagSet.StringVar(&bindRsIp, "bind-rs-ip", "", "rs host ip to bind")
	flagSet.StringVar(&bindNicIp, "bind-nic-ip", "", "local network interface card to bind")
	flagSet.BoolVar(&rescanLocal, "rescan-local", false, "rescan local dir to upload newly add files")
	flagSet.StringVar(&logFile, "log-file", "", "log file")
	flagSet.StringVar(&logLevel, "log-level", "info", "log level")

	flagSet.Parse(params)

	uploadConfig := qshell.UploadConfig{
		SrcDir:           srcDir,
		FileList:         fileList,
		Bucket:           bucket,
		PutThreshold:     putThreshold,
		KeyPrefix:        keyPrefix,
		IgnoreDir:        ignoreDir,
		Overwrite:        overwrite,
		CheckExists:      checkExists,
		CheckHash:        checkHash,
		CheckSize:        checkSize,
		SkipFixedStrings: skipFixedStrings,
		SkipFilePrefixes: skipFilePrefixes,
		SkipPathPrefixes: skipPathPrefixes,
		SkipSuffixes:     skipSuffixes,
		UpHost:           upHost,
		BindUpIp:         bindUpIp,
		BindRsIp:         bindRsIp,
		BindNicIp:        bindNicIp,
		RescanLocal:      rescanLocal,
		LogFile:          logFile,
		LogLevel:         logLevel,
	}

	//check params
	if uploadConfig.SrcDir == "" {
		fmt.Println("Upload config no `--src-dir` specified")
		return
	}

	if uploadConfig.Bucket == "" {
		fmt.Println("Upload config no `--bucket` specified")
		return
	}

	if _, err := os.Stat(uploadConfig.SrcDir); err != nil {
		log.Error("Upload config `SrcDir` not exist error,", err)
		return
	}

	if threadCount < qshell.MIN_UPLOAD_THREAD_COUNT || threadCount > qshell.MAX_UPLOAD_THREAD_COUNT {
		fmt.Printf("Tip: you can set <ThreadCount> value between %d and %d to improve speed\n",
			qshell.MIN_UPLOAD_THREAD_COUNT, qshell.MAX_UPLOAD_THREAD_COUNT)

		if threadCount < qshell.MIN_UPLOAD_THREAD_COUNT {
			threadCount = qshell.MIN_UPLOAD_THREAD_COUNT
		} else if threadCount > qshell.MAX_UPLOAD_THREAD_COUNT {
			threadCount = qshell.MAX_UPLOAD_THREAD_COUNT
		}
	}

	qshell.QiniuUpload(int(threadCount), &uploadConfig)
}
