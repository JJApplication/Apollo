package secure_manager

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/JJApplication/Apollo/logger"
	"go.uber.org/zap"
)

func TestSSHGuardWatcher(t *testing.T) {
	// 初始化 Logger 防止 panic
	if logger.Logger == nil {
		l, _ := zap.NewDevelopment()
		logger.Logger = l
		logger.LoggerSugar = l.Sugar()
	}

	// 获取 auth.log 绝对路径
	cwd, _ := os.Getwd()
	// 假设测试在 app/secure_manager 目录下运行
	logPath := filepath.Join(cwd, "auth.log")

	// 如果文件不存在，尝试使用相对路径（兼容不同测试运行环境）
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		logPath = "auth.log"
	}

	t.Logf("Testing with log file: %s", logPath)

	// 等待读取和解析
	// auth.log 文件较小，2秒足够
	time.Sleep(2 * time.Second)

	file, err := os.Open(logPath)
	if err != nil {
		return
	}
	defer file.Close()

	// 检查数据库是否有数据
	// 如果没有数据，说明是首次初始化或数据丢失，需要从头读取日志恢复数据
	// 如果有数据，则只监听新增日志

	reader := bufio.NewReader(file)

	// 正则匹配: Dec 15 07:46:26 ... Blocking "IP" for 3600 secs
	re := regexp.MustCompile(`^([A-Z][a-z]{2}\s+\d+\s+\d{2}:\d{2}:\d{2}).*sshguard\[\d+\]: Blocking "([^"]+)" for (\d+) secs`)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			time.Sleep(1 * time.Second)
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

			t.Logf("Testing SSHGuard block: %s", block)
		}
	}
}
