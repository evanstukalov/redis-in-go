package commands

import (
	"context"
	"fmt"
	"io"

	"github.com/codecrafters-io/redis-starter-go/internal/config"
	"github.com/codecrafters-io/redis-starter-go/internal/utils"
)

func (c *KeysCommand) handleAll(
	ctx context.Context,
	conn io.Writer,
	config config.Config,
	args []string,
) {
	fileContent := utils.ReadFile(config.RedisDir + "/" + config.RedisDbFileName)
	conn.Write([]byte(fmt.Sprintf("*1\r\n$%d\r\n%s\r\n", len(fileContent), fileContent)))
}
