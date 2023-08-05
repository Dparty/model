package uptime

import (
	"fmt"
	"testing"

	"github.com/Dparty/common/constants"
	"github.com/Dparty/model"
	"github.com/Dparty/model/common"
	"github.com/spf13/viper"
)

func TestCreateMonitor(t *testing.T) {
	var err error
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("databases fatal error config file: %w", err))
	}
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	db, err := model.NewConnection(user, password, host, port, database)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Monitor{})
	n := &Monitor{
		Notifications: Notifications{Notification{
			Type: constants.EMAIL,
			EmailReceivers: &EmailReceivers{
				To: common.StringList{"a", "b"},
			},
		}},
	}
	db.Save(&n)
	b := &Monitor{}
	db.Find(b, n.ID)
	t.Log(b.Notifications[0].EmailReceivers.To)
}
