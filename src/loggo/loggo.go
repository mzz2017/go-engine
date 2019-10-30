package loggo

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	LEVEL_DEBUG = iota
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
)

type Config struct {
	Level     int
	Prefix    string
	MaxDay    int
	NoLogFile bool
}

var gConfig Config

func Ini(config Config) {
	initbefore := false
	if gConfig.MaxDay != 0 {
		fmt.Println("loggo had ini before " + gConfig.Prefix)
		initbefore = true
	}

	gConfig = config
	if gConfig.Prefix == "" {
		panic("log prefix is empty")
	}

	if strings.Contains(gConfig.Prefix, "_") {
		panic("log prefix contain _")
	}

	Warn("loggo Ini")

	if !initbefore {
		go loopCheck(gConfig)
	}
}

func Debug(format string, a ...interface{}) {
	if gConfig.Level <= LEVEL_DEBUG {
		str := genLog(LEVEL_DEBUG, format, a...)
		if !gConfig.NoLogFile {
			file := openLog(LEVEL_DEBUG)
			defer file.Close()
			file.WriteString(str)
		}
		fmt.Print(str)
	}
}

func Info(format string, a ...interface{}) {
	if gConfig.Level <= LEVEL_INFO {
		str := genLog(LEVEL_INFO, format, a...)
		if !gConfig.NoLogFile {
			file := openLog(LEVEL_INFO)
			defer file.Close()
			file.WriteString(str)
		}
		if gConfig.Level <= LEVEL_DEBUG {
			if !gConfig.NoLogFile {
				file1 := openLog(LEVEL_DEBUG)
				defer file1.Close()
				file1.WriteString(str)
			}
		}
		fmt.Print(str)
	}
}

func Warn(format string, a ...interface{}) {
	if gConfig.Level <= LEVEL_WARN {
		str := genLog(LEVEL_WARN, format, a...)
		if !gConfig.NoLogFile {
			file := openLog(LEVEL_WARN)
			defer file.Close()
			file.WriteString(str)
		}
		if gConfig.Level <= LEVEL_INFO {
			if !gConfig.NoLogFile {
				file1 := openLog(LEVEL_INFO)
				defer file1.Close()
				file1.WriteString(str)
			}
		}
		if gConfig.Level <= LEVEL_DEBUG {
			if !gConfig.NoLogFile {
				file2 := openLog(LEVEL_DEBUG)
				defer file2.Close()
				file2.WriteString(str)
			}
		}
		fmt.Print(str)
	}
}

func Error(format string, a ...interface{}) {
	if gConfig.Level <= LEVEL_ERROR {
		str := genLog(LEVEL_ERROR, format, a...)
		if !gConfig.NoLogFile {
			file := openLog(LEVEL_ERROR)
			defer file.Close()
			file.WriteString(str)
		}
		if gConfig.Level <= LEVEL_WARN {
			if !gConfig.NoLogFile {
				file0 := openLog(LEVEL_WARN)
				defer file0.Close()
				file0.WriteString(str)
			}
		}
		if gConfig.Level <= LEVEL_INFO {
			if !gConfig.NoLogFile {
				file1 := openLog(LEVEL_INFO)
				defer file1.Close()
				file1.WriteString(str)
			}
		}
		if gConfig.Level <= LEVEL_DEBUG {
			if !gConfig.NoLogFile {
				file2 := openLog(LEVEL_DEBUG)
				defer file2.Close()
				file2.WriteString(str)
			}
		}
		fmt.Print(str)
	}
}

func genLog(level int, format string, a ...interface{}) string {
	file, funcName, line := getFunc()
	t := time.Now().Format(time.RFC3339Nano)
	str := fmt.Sprintf(format, a...)
	ret := fmt.Sprintf("[%v] [%v] [%v:%v] [%v] %v\n", levelName(level), t, file, line, funcName, str)
	return ret
}

func getFunc() (string, string, int) {
	// Ask runtime.Callers for up to 5 pcs, including runtime.Callers itself.
	pc := make([]uintptr, 5)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return "NIL", "NIL", 0
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	n = 5

	// Loop to get frames.
	// A fixed number of pcs can expand to an indefinite number of Frames.
	for i := 0; i < n; i++ {
		frame, more := frames.Next()
		if i == n-1 {
			return frame.File, frame.Function, frame.Line
		}
		if !more {
			break
		}
	}
	return "NIL", "NIL", 0
}

func levelName(level int) string {
	switch level {
	case LEVEL_DEBUG:
		return "DEBUG"
	case LEVEL_INFO:
		return "INFO"
	case LEVEL_WARN:
		return "WARN"
	case LEVEL_ERROR:
		return "ERROR"
	}
	return "NIL"
}

func NameToLevel(name string) int {
	switch strings.ToUpper(name) {
	case "DEBUG":
		return LEVEL_DEBUG
	case "INFO":
		return LEVEL_INFO
	case "WARN":
		return LEVEL_WARN
	case "ERROR":
		return LEVEL_ERROR
	}
	return -1
}

func openLog(level int) os.File {
	date := time.Now().Format("2006-01-02")
	fileName := gConfig.Prefix + "_" + levelName(level) + "_" + date + ".log"
	f, e := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if e != nil {
		panic(e)
	}
	return *f
}

func checkDate(config Config) {
	now := time.Now().Format("2006-01-02")
	nowt, _ := time.Parse("2006-01-02", now)
	nowunix := nowt.Unix()
	filepath.Walk("./", func(path string, f os.FileInfo, err error) error {

		if f == nil || f.IsDir() {
			return nil
		}

		if !strings.HasSuffix(f.Name(), ".log") {
			return nil
		}

		strs := strings.Split(f.Name(), "_")
		if strs == nil || len(strs) < 3 {
			return nil
		}

		if strs[0] != config.Prefix {
			return nil
		}

		date := strs[2]
		date = strings.TrimRight(date, ".log")

		t, e := time.Parse("2006-01-02", date)
		if e != nil {
			Error("loggo delete Parse file fail %v %v %v", f.Name(), date, err)
			return nil
		}
		tunix := t.Unix()
		if nowunix-tunix > int64(config.MaxDay)*24*3600 {
			err := os.Remove(f.Name())
			if e != nil {
				Error("loggo delete log file fail %v %v", f.Name(), err)
				return nil
			}
		}

		return nil
	})
}

func loopCheck(config Config) {
	for {
		checkDate(config)
		time.Sleep(time.Minute)
	}
}
