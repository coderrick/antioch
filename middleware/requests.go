package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// lock();
	unlock()

}

func lock() {
	url := "https://api-production.august.com/remoteoperate/11CE24327D5649E0BCC8E79AE3D15D1E/lock?=,,,"

	req, _ := http.NewRequest("PUT", url, nil)

	req.Header.Add("x-august-api-key", "a15e098f-0088-489a-a106-5de86657900c")
	req.Header.Add("x-august-access-token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpbnN0YWxsSWQiOiIiLCJhcHBsaWNhdGlvbklkIjoiIiwidXNlcklkIjoiMmQ4MmUzNTctYzJlYy00YzM3LTkxMjctYWQ4NjdiMWJkZTdmIiwidkluc3RhbGxJZCI6ZmFsc2UsInZQYXNzd29yZCI6dHJ1ZSwidkVtYWlsIjp0cnVlLCJ2UGhvbmUiOnRydWUsImhhc0luc3RhbGxJZCI6ZmFsc2UsImhhc1Bhc3N3b3JkIjpmYWxzZSwiaGFzRW1haWwiOmZhbHNlLCJoYXNQaG9uZSI6ZmFsc2UsImlzTG9ja2VkT3V0IjpmYWxzZSwiY2FwdGNoYSI6IiIsImVtYWlsIjpbXSwicGhvbmUiOltdLCJleHBpcmVzQXQiOiIyMDIwLTAxLTMwVDE4OjQ1OjQ3LjczM1oiLCJ0ZW1wb3JhcnlBY2NvdW50Q3JlYXRpb25QYXNzd29yZExpbmsiOiIiLCJpYXQiOjE1NzAwNDE5NDcsImV4cCI6bnVsbCwib2F1dGgiOnsiYXBwX25hbWUiOiJUQyBIYWNrYXRob24iLCJjbGllbnRfaWQiOiI1YzI4NmRlYi1lMWIzLTRiMTItODZjZC00MmY1YWMwYmEwMDEiLCJhcGlLZXkiOiJhMTVlMDk4Zi0wMDg4LTQ4OWEtYTEwNi01ZGU4NjY1NzkwMGMiLCJyZWRpcmVjdF91cmkiOiJodHRwczovL3BhcnRuZXJzLmF1Z3VzdC5jb20vb2F1dGh0ZXN0L2NiIn19.Ry-BYuQcv7im2x1wF4Hj48f-nynHtW1sFVpQGdTI41E")
	req.Header.Add("accept-version", "0.0.1")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "PostmanRuntime/7.15.2")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "cc503d8f-7c18-47bc-815a-9209ddc3ac71,e2bcf72f-a860-42ba-b7ab-c97c4fe81cce")
	req.Header.Add("Host", "api-production.august.com")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func unlock() {
	url := "https://api-production.august.com/remoteoperate/11CE24327D5649E0BCC8E79AE3D15D1E/unlock"

	req, _ := http.NewRequest("PUT", url, nil)

	req.Header.Add("x-august-api-key", "a15e098f-0088-489a-a106-5de86657900c")
	req.Header.Add("x-august-access-token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpbnN0YWxsSWQiOiIiLCJhcHBsaWNhdGlvbklkIjoiIiwidXNlcklkIjoiMmQ4MmUzNTctYzJlYy00YzM3LTkxMjctYWQ4NjdiMWJkZTdmIiwidkluc3RhbGxJZCI6ZmFsc2UsInZQYXNzd29yZCI6dHJ1ZSwidkVtYWlsIjp0cnVlLCJ2UGhvbmUiOnRydWUsImhhc0luc3RhbGxJZCI6ZmFsc2UsImhhc1Bhc3N3b3JkIjpmYWxzZSwiaGFzRW1haWwiOmZhbHNlLCJoYXNQaG9uZSI6ZmFsc2UsImlzTG9ja2VkT3V0IjpmYWxzZSwiY2FwdGNoYSI6IiIsImVtYWlsIjpbXSwicGhvbmUiOltdLCJleHBpcmVzQXQiOiIyMDIwLTAxLTMwVDE4OjQ1OjQ3LjczM1oiLCJ0ZW1wb3JhcnlBY2NvdW50Q3JlYXRpb25QYXNzd29yZExpbmsiOiIiLCJpYXQiOjE1NzAwNDE5NDcsImV4cCI6bnVsbCwib2F1dGgiOnsiYXBwX25hbWUiOiJUQyBIYWNrYXRob24iLCJjbGllbnRfaWQiOiI1YzI4NmRlYi1lMWIzLTRiMTItODZjZC00MmY1YWMwYmEwMDEiLCJhcGlLZXkiOiJhMTVlMDk4Zi0wMDg4LTQ4OWEtYTEwNi01ZGU4NjY1NzkwMGMiLCJyZWRpcmVjdF91cmkiOiJodHRwczovL3BhcnRuZXJzLmF1Z3VzdC5jb20vb2F1dGh0ZXN0L2NiIn19.Ry-BYuQcv7im2x1wF4Hj48f-nynHtW1sFVpQGdTI41E")
	req.Header.Add("accept-version", "0.0.1")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "PostmanRuntime/7.15.2")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "abf46895-6448-4dcc-91be-478fea8976b2,d059dc83-fa1a-4c74-a81d-3a951b67d835")
	req.Header.Add("Host", "api-production.august.com")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
