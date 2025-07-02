package handler

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/yeawyow/gateway/service"
)

type LineTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type LineProfile struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

func LineHandler(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "code not provided",
		})
	}

	client := resty.New()

	// ขอ token จาก LINE
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "authorization_code",
			"code":          code,
			"redirect_uri":  os.Getenv("URL_CALLBACK_APP"), // ต้องตรงกับที่ LINE Dev ตั้งไว้
			"client_id":     os.Getenv("LINE_CHANNEL_ID"),
			"client_secret": os.Getenv("LINE_CHANNEL_SECRET"),
		}).
		Post("https://api.line.me/oauth2/v2.1/token")

	if err != nil || resp.StatusCode() != 200 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "LINE token request failed",
			"body":  resp.String(),
		})
	}

	var tokenResp LineTokenResponse
	if err := json.Unmarshal(resp.Body(), &tokenResp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot parse token response",
		})
	}

	// ดึงข้อมูล profile
	profileResp, err := client.R().
		SetHeader("Authorization", "Bearer "+tokenResp.AccessToken).
		Get("https://api.line.me/v2/profile")

	if err != nil || profileResp.StatusCode() != 200 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot get LINE profile",
			"body":  profileResp.String(),
		})
	}

	var profile LineProfile
	if err := json.Unmarshal(profileResp.Body(), &profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot parse profile response",
		})
	}
    userp ,err := service.AuthLineService(profile.UserID)
if err != nil {
	return c.Status(500).JSON(fiber.Map{"error": "internal error"})
}
if userp == nil {
	return c.Status(404).JSON(fiber.Map{"error": "user not found","user":profile.UserID,"profile":profile})
}
	
	// ส่ง token + profile กลับไป
	return c.JSON(fiber.Map{
		"token":   tokenResp,
		// "profile": profile,
		"user_idapp":userp,
		"profile":profile,
	})
}
