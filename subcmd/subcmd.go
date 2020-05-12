package subcmd

import (
	"fmt"
	"os"

	"github.com/vim-jp/slacklog-generator/subcmd/serve"
)

// Run runs one of sub-commands.
func Run() error {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println(`Usage: go run . {subcmd}
  Subcmd:
    convert-exported-logs
    download-emoji
    download-files
    generate-html`)
		return nil
	}

	args := os.Args[2:]
	subCmdName := os.Args[1]
	switch subCmdName {
	case "convert-exported-logs":
		return ConvertExportedLogs(args)
	case "download-emoji":
		return DownloadEmoji(args)
	case "download-files":
		return DownloadFiles(args)
	case "generate-html":
		return GenerateHTML(args)
	case "serve":
		return serve.Run(args)
	}

	return fmt.Errorf("unknown subcmd: %s", subCmdName)
}
