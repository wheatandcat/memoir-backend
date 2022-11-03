package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"
)

const verifyCustomTokenURL = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s"

type User struct {
	UserToken string `yaml:"userToken"`
}

type LoginYaml struct {
	Vars User
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("読み込み出来ませんでした: %v", err)
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile("../serviceAccount.json")
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.CustomToken(ctx, "some-uid")
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	req, err := json.Marshal(map[string]interface{}{
		"token":             token,
		"returnSecureToken": true,
	})
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	apiKey := os.Getenv("FIREBASE_API_KEY")

	resp, err := postRequest(fmt.Sprintf(verifyCustomTokenURL, apiKey), req)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	var respBody struct {
		IDToken string `json:"idToken"`
	}
	if err := json.Unmarshal(resp, &respBody); err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	ly := LoginYaml{}
	ly.Vars.UserToken = respBody.IDToken

	if err := WriteOnFile("./login.yaml", ly); err != nil {
		log.Fatalf("WriteOnFile: %v\n", err)
	}

	if err = os.Chmod("./login.yaml", 0600); err != nil {
		log.Fatalf("OSコマンドで失敗: %v", err)
	}
}

func postRequest(url string, req []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req)) //nolint:gosec
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status code: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func WriteOnFile(fileName string, data interface{}) error {
	// ここでデータを []byte に変換しています。
	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	// []byte をファイルに上書きしています。
	err = os.WriteFile(fileName, buf, os.ModeExclusive)
	if err != nil {
		return err
	}
	return nil
}
