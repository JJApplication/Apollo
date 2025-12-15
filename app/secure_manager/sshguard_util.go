package secure_manager

import (
	"bufio"
	"github.com/JJApplication/Apollo/logger"
	"io"
	"os"
	"regexp"
	"time"
)

// SSHGuardBlock 屏蔽信息结构体
type SSHGuardBlock struct {
	IP       string // 屏蔽IP
	Time     string // 屏蔽时间
	Duration string // 屏蔽时长(秒)
}

var (
	AuthLogPath = "/var/log/auth.log"
)

// StartSSHGuardWatcher 启动SSHGuard日志监控
// logPath: 日志文件路径，默认为 /var/log/auth.log
func StartSSHGuardWatcher() {
	if AuthLogPath == "" {
		AuthLogPath = "/var/log/auth.log"
	}
	logger.LoggerSugar.Infof("%s Starting SSH Guard Watcher...", SecureManagerPrefix)
	go watchAuthLog(AuthLogPath)
}

func watchAuthLog(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	// 检查数据库是否有数据
	// 如果没有数据，说明是首次初始化或数据丢失，需要从头读取日志恢复数据
	// 如果有数据，则只监听新增日志
	existingData := GetSecureList()
	shouldReadAll := len(existingData) == 0

	if !shouldReadAll {
		// 移动到文件末尾
		file.Seek(0, io.SeekEnd)
	}

	reader := bufio.NewReader(file)

	// 正则匹配: Dec 15 07:46:26 ... Blocking "IP" for 3600 secs
	re := regexp.MustCompile(`^([A-Z][a-z]{2}\s+\d+\s+\d{2}:\d{2}:\d{2}).*sshguard\[\d+\]: Blocking "([^"]+)" for (\d+) secs`)

	for {
		// 安全日志的解析影响性能 无需过快
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// 检查日志轮转 (Log Rotation)
				// 1. 获取当前文件路径的最新状态
				newFi, statErr := os.Stat(path)
				// 2. 获取当前持有文件的状态
				curFi, curStatErr := file.Stat()

				if statErr == nil && curStatErr == nil {
					// 如果两个文件对象不代表同一个文件（Inode变化），说明发生了轮转
					if !os.SameFile(curFi, newFi) {
						file.Close()
						newFile, openErr := os.Open(path)
						if openErr == nil {
							file = newFile
							reader = bufio.NewReader(file)
							continue // 立即开始读取新文件
						}
					}
				}

				time.Sleep(60 * time.Second)
				continue
			}
			time.Sleep(60 * time.Second)
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) == 4 {
			now := time.Now()
			tm, _ := time.Parse(time.Stamp, matches[1])
			currentYear := now.Year()
			currentMonth := now.Month()
			if tm.Month().String() > currentMonth.String() {
				// 去年的数据
				currentYear = currentYear - 1
			}
			finalDate := time.Date(
				currentYear,
				tm.Month(),
				tm.Day(),
				tm.Hour(),
				tm.Minute(),
				tm.Second(),
				0, now.Location(),
			).Format(time.DateTime)

			block := SSHGuardBlock{
				Time:     finalDate,
				IP:       matches[2],
				Duration: matches[3],
			}

			// 存储到数据库
			saveSecureAudit(&SecureAudit{
				SecureType: TypeSSH,
				BlockIP:    block.IP,
				BlockStart: block.Time,
				BlockTime:  block.Duration,
				Remark:     "SSHGuard Block",
			})
		}
	}
}
