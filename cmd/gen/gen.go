package gen

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/skyrocketOoO/gorm-plugin/columnname"
	"github.com/skyrocketOoO/gorm-plugin/tablename"
	"github.com/skyrocketOoO/serverx/internal/boot"
	"github.com/skyrocketOoO/serverx/internal/domain"
	database "github.com/skyrocketOoO/serverx/internal/service/database"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "",
	Long:  ``,
	// Args:  cobra.MinimumNArgs(1),
	Run: Gen,
}

func Gen(cmd *cobra.Command, args []string) {
	database.New()
	if err := boot.Run(); err != nil {
		log.Fatal().Msgf("Initialization failed: %v", err)
	}

	db, err := database.New()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
	defer database.Close(db)

	if err := tablename.GenTableNamesCode(db,
		"internal/gen/table/table.go"); err != nil {
		log.Fatal().Msgf("%v", err)
	}

	tablenames, err := tablename.GetTableNames(db)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	if err := columnname.GenTableColumnNamesCode(db, tablenames,
		"internal/gen/column/column.go"); err != nil {
		log.Fatal().Msgf("%v", err)
	}

	fmt.Println("Done")
}

func init() {
	Cmd.Flags().
		StringVarP(&domain.Database, `database`, "d", "postgres", `"postgres", "sqlite", "mysql"`)
}
