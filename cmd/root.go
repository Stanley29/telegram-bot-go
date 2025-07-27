package cmd

import (
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/spf13/cobra"
    tele "gopkg.in/telebot.v4"
)

// ✨ Завантажує .env при ініціалізації пакету
func init() {
    err := godotenv.Load()
    if err != nil {
        log.Println(".env файл не знайдено або не вдалося завантажити")
    }
}

// 🔧 Головна команда CLI
var rootCmd = &cobra.Command{
    Use:   "telegram-bot",
    Short: "Telegram bot powered by Go",
    Run: func(cmd *cobra.Command, args []string) {
        token := os.Getenv("TELE_TOKEN")
        if token == "" {
            log.Fatal("TELE_TOKEN is not set")
        }

        bot, err := tele.NewBot(tele.Settings{
            Token:  token,
            Poller: &tele.LongPoller{Timeout: 10 * time.Second},
        })
        if err != nil {
            log.Fatal(err)
        }

        bot.Handle("/start", func(c tele.Context) error { 
            return c.Send("Вітаю, Serhii! Бот працює ✅")
        })

        bot.Handle(tele.OnText, func(c tele.Context) error {
            return c.Send("Привіт! Я працюю 🎉")
        })

        bot.Start()
    },
}

// 🏁 Викликає команду
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
