package service

import (
	"fmt"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/global"
	"golang.org/x/crypto/ssh"
)

type OTAService struct{}

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

	// 连接到SSH服务器
	connection, err := ssh.Dial("tcp", global.GlobalConfig.BuildHost, sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

	// 创建SSH会话
	session, err := connection.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	commands := []string{
		"echo " + version + " > " + global.GlobalConfig.BuildPath + "/buildroot-external/meta2",
		"echo -e \"" + releaseNote + "\" > " + global.GlobalConfig.BuildPath + "/buildroot-external/release-note-2.md",
		"cd " + global.GlobalConfig.BuildPath + " && " + global.GlobalConfig.BuildPath + "/scripts/enter-no-it.sh make zimacube_recovery",
	}

	// 依次执行每个命令
	for _, cmd := range commands {
		fmt.Println(cmd)
		output, err := runCommand(connection, cmd)
		if err != nil {
			log.Fatalf("Failed to run command '%s': %s", cmd, err)
		}
		fmt.Printf("Output of '%s':\n%s\n", cmd, output)
	}

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
