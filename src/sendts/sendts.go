package main

import (
	"fmt"
	"strings"
	"os"
	"os/exec"
	"io/ioutil"
)

// 사용자 이름이나 이런것들은 서버마다 다를 수 있기에 이건 조정해줘야하거나 동기화시켜야해
// 버퍼에서 ip나 경로등은 인자로 받거나 grpc통신 등등으로 받아서 구현하면 될듯
// 누군가 악성 .ts파일을 업로드 하면 어떡하지
// 로그를 사용하는, 남기는 방식을 추가해야함
const (
	VIDEO_PATH       string = "/mnt/c/Users/winix/Videos/독수리"
	CLIENT_NAME      string = "roo"
	CLIENT_FILE_PATH string = "/tmp/test"
	PRIVATE_KEY_PATH string = "/home/roo/.ssh/tmpkey"
	REMOTE_SERVER_IP string = "192.168.27.128"
)

func main() {
	fmt.Println("[START]");

	for {
		// 데몬으로 동작
		// 너무 빠르니 스트리밍 속도에 따라서 지연시키는 기능도 추가하자
		files, err := ioutil.ReadDir(VIDEO_PATH)
		if err != nil {
			fmt.Println(err)
		} else if len(files) == 0 {
			fmt.Println("[No file]")
		}

		for _, file := range files {
			// 파일인 경우만 진행
			if file.IsDir() != true {
				// 특정 확장자만 진행
				if strings.HasSuffix(file.Name(), ".ts") == true || strings.HasSuffix(file.Name(), ".m3u8") == true {

					buf := fmt.Sprintf("scp -i %s -r %s/%s %s@%s:%s", PRIVATE_KEY_PATH, VIDEO_PATH, file.Name(), CLIENT_NAME, REMOTE_SERVER_IP, CLIENT_FILE_PATH)
					cmd := exec.Command("/bin/sh", "-c", buf)
					fmt.Println(buf)

					// 에러면 지나가고 성공했으면 백업폴더로 옮긴다. 전송 실패하면 다시 시도하는 코드를 추가해야함. 모든 소스에서 전송 실패, 에러 처리에 대한 기능을 추가해야함
					err := cmd.Run()
					if err != nil {
						fmt.Println(err)
						continue
					}

					// 보낸 .ts들만 다른 디렉터리에 보관. 함수 계속 쓰기 싫으면 flag같은걸로 처리해도 좋을듯
					// 위의 조건문과 겹쳐서 이건 분기를 조정해도 좋을듯
					if strings.HasSuffix(file.Name(), ".ts") == true {

						err := os.Rename(fmt.Sprintf("%s/%s", VIDEO_PATH, file.Name()), fmt.Sprintf("%s완료/%s", VIDEO_PATH, file.Name()))
						if err != nil {
							fmt.Println(err);
						}
					}
				}
			}
		}
	}
	fmt.Println("[END]");
}

/*-------------------------------------------------------------------------------
[Reference]

1. mkdir
	err := os.Mkdir("testDir", 0777)
2. mv
	cmd := exec.Command("/bin/sh", "-c", "mv testDir SecondDir")
	cmd.Run()
3. rename
	err := os.Rename("test.txt", "bak/test.txt")
	fmt.Println(err);
-------------------------------------------------------------------------------*/
