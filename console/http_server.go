package console

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/notblessy/skelago/config"
	"github.com/notblessy/skelago/db"
	"github.com/notblessy/skelago/delivery/http"
	"github.com/notblessy/skelago/repository"
	"github.com/notblessy/skelago/usecase"
	"github.com/notblessy/skelago/utils"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var runHTTPServer = &cobra.Command{
	Use:   "httpsrv",
	Short: "run http server",
	Long:  `This subcommand is for starting the http server`,
	Run:   runHTTP,
}

func init() {
	rootCmd.AddCommand(runHTTPServer)
}

func runHTTP(cmd *cobra.Command, args []string) {
	psql := db.InitDB()
	defer db.CloseDB(psql)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	welcomeRepo := repository.NewWelcomeRepository(psql)
	welcomeUsecase := usecase.NewWelcomeUsecase(welcomeRepo)

	httpSvc := http.NewHTTPService()
	httpSvc.RegisterWelcomeUsecase(welcomeUsecase)

	httpSvc.Routes(e)

	logrus.Fatal(e.Start(":" + config.HTTPPort()))
}
