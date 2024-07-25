package sb

import (
	"os"

	"github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func Init() error {
	sbURL := os.Getenv("SUPABASE_URL")
	sbSecret := os.Getenv("SUPABASE_SECRET")
	Client = supabase.CreateClient(sbURL, sbSecret)
	return nil
}
