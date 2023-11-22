package service

import (
	"fmt"
	"log"
	"regexp"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/global"
	"golang.org/x/crypto/ssh"
)

type OTAService struct {
	Status  string
	Message string
}

func parseVersion(version string) string {
	// Regular expression to extract version components
	re := regexp.MustCompile(`^v(\d+)\.(\d+)\.(\d+)([-\.]?\d*)$`)
	matches := re.FindStringSubmatch(version)

	// Check if the version string is valid
	if matches == nil {
		return "Invalid version format"
	}

	// Extracting the version components
	versionMajor := matches[1]
	versionBuild := matches[2]
	versionDev := matches[3]
	versionSub := matches[4]

	// Formatting the output
	output := fmt.Sprintf("VERSION_MAJOR=%s\nVERSION_BUILD=%s\nVERSION_DEV=%s\nVERSION_SUB=\"%s\"\n\nCASAOS_NAME=\"ZimaOS\"\nCASAOS_ID=\"zimaos\"\n\nDEPLOYMENT=\"development\"\n", versionMajor, versionBuild, versionDev, versionSub)
	return output
}

func (s *OTAService) Build(version string, releaseNote string) error {
	sshConfig := &ssh.ClientConfig{
		User: global.GlobalConfig.BuildUser, // 替换为实际的用户名
		Auth: []ssh.AuthMethod{
			ssh.Password(global.GlobalConfig.BuildPassword), // 替换为实际的密码
			// 或使用密钥
			// ssh.PublicKeys(privateKey()),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不检查服务器的公钥
	}

	s.Status = "building"
	// 连接到SSH服务器
	connection, err := ssh.Dial("tcp", global.GlobalConfig.BuildHost, sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

	// 创建SSH会话
	session, err := connection.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: ", err)
	}
	defer session.Close()

	releaseVersion := parseVersion(version)
	commands := []string{
		"echo -e \"" + releaseVersion + "\" > " + global.GlobalConfig.BuildPath + "/buildroot-external/meta",
		"echo -e \"" + releaseNote + "\" > " + global.GlobalConfig.BuildPath + "/buildroot-external/release-note.md",
		"cd " + global.GlobalConfig.BuildPath + " && " + global.GlobalConfig.BuildPath + "/scripts/enter-no-it.sh make zimacube",
	}

	// 依次执行每个命令
	for _, cmd := range commands {
		fmt.Println(cmd)
		output, err := runCommand(connection, cmd)
		if err != nil {
			s.Status = "failed"
			s.Message = err.Error() + " " + output
			fmt.Println("Failed to run command ", cmd, err, output)
			return err
		}
		fmt.Printf("Output of '%s':\n%s\n", cmd, output)
	}

	s.Status = "success"
	s.Message = "构建完成"
	return nil
}

func runCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
