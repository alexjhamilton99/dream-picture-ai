package sb

import (
	"os"

	"github.com/nedpals/supabase-go"
)

const BaseAuthURL = "https://mbeifufogwugxyzdrted.supabase.co/auth/v1/recover"

const ResetPasswordEndPoint = "auth/v1/recover"

var Client *supabase.Client

func Init() error {
	sbURL := os.Getenv("SUPABASE_URL")
	sbSecret := os.Getenv("SUPABASE_SECRET")
	Client = supabase.CreateClient(sbURL, sbSecret)
	return nil
}
