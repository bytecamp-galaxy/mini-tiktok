package aliyun

import "fmt"

func snapshot(playUrl string) string {
	return fmt.Sprintf("%d?x-oss-process=video/snapshot,t_0,f_jpg,w_0,h_0,m_fast",
		playUrl)
}
