package service

import (
	"fmt"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/global"
	"golang.org/x/crypto/ssh"
)

type OTAService struct{}

func (s *OTAService) Build() error {
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

	// 执行命令
	command := "cd /home/ctrdh/operating-system && /home/ctrdh/operating-system/scripts/enter-no-it.sh make zimacube_recovery" // 替换为要执行的命令
	// command := "ls" // 替换为要执行的命令

	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("Failed to run: %s", err, string(output))
	}
	fmt.Println(string(output))

	return nil
}
