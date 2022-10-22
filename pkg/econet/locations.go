package econet

type Locations struct {
	Results struct {
		Locations []struct {
			AWAY           bool   `json:"@AWAY"`
			AWAYCONFIG     bool   `json:"@AWAYCONFIG"`
			LOCATIONINFO   string `json:"@LOCATION_INFO"`
			LOCATIONNAME   string `json:"@LOCATION_NAME"`
			LOCATIONSTATUS string `json:"@LOCATION_STATUS"`
			VACATION       bool   `json:"@VACATION"`
			VACATIONCONFIG bool   `json:"@VACATIONCONFIG"`
			WEATHER        string `json:"@WEATHER"`
			WEATHERF       int    `json:"@WEATHER_F"`
			WEATHERI       string `json:"@WEATHER_I"`
			Equiptments    []struct {
				ACTIVE     bool   `json:"@ACTIVE"`
				ALERTCOUNT int    `json:"@ALERTCOUNT"`
				AWAY       bool   `json:"@AWAY"`
				AWAYCONFIG bool   `json:"@AWAYCONFIG"`
				AWAYMSG    string `json:"@AWAY_MSG"`
				BCONFIG    []struct {
					Align string `json:"align,omitempty"`
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
					Title string `json:"title,omitempty"`
				} `json:"@BCONFIG"`
				CONFIG []struct {
					Align string `json:"align"`
					Name  string `json:"name"`
					Title string `json:"title"`
					Type  string `json:"type"`
					Value struct {
						Constraints struct {
							Bgcolor      string        `json:"bgcolor"`
							EnumText     []string      `json:"enumText"`
							EnumTextIcon []interface{} `json:"enumTextIcon"`
							Fontcolor    string        `json:"fontcolor"`
							Icon         string        `json:"icon"`
							LowerLimit   int           `json:"lowerLimit"`
							UpperLimit   int           `json:"upperLimit"`
						} `json:"constraints"`
						Status string `json:"status"`
						Value  int    `json:"value"`
					} `json:"value"`
				} `json:"@CONFIG"`
				CONNECTED bool `json:"@CONNECTED"`
				DRACTIVE  struct {
					Constraints struct {
						Dialog []struct {
							Message string `json:"message"`
							Title   string `json:"title"`
							Value   int    `json:"value"`
						} `json:"dialog"`
					} `json:"constraints"`
					Value string `json:"value"`
				} `json:"@DRACTIVE"`
				MODE struct {
					Constraints struct {
						EnumText     []string `json:"enumText"`
						EnumTextIcon []string `json:"enumTextIcon"`
						LowerLimit   int      `json:"lowerLimit"`
						UpperLimit   int      `json:"upperLimit"`
					} `json:"constraints"`
					Status string `json:"status"`
					Value  int    `json:"value"`
				} `json:"@MODE"`
				MODECONFIG struct {
					Constraints struct {
						Bgcolor      string        `json:"bgcolor"`
						EnumText     []string      `json:"enumText"`
						EnumTextIcon []interface{} `json:"enumTextIcon"`
						Fontcolor    string        `json:"fontcolor"`
						Icon         string        `json:"icon"`
						LowerLimit   int           `json:"lowerLimit"`
						UpperLimit   int           `json:"upperLimit"`
					} `json:"constraints"`
					Status string `json:"status"`
					Value  int    `json:"value"`
				} `json:"@MODECONFIG"`
				MODEIMAGE string `json:"@MODEIMAGE"`
				NAME      struct {
					Constraints struct {
						StringLength int `json:"stringLength"`
					} `json:"constraints"`
					Value string `json:"value"`
				} `json:"@NAME"`
				RESUME         bool   `json:"@RESUME"`
				RUNNING        string `json:"@RUNNING"`
				SCHEDULE       bool   `json:"@SCHEDULE"`
				SCHEDULERESUME string `json:"@SCHEDULERESUME"`
				SCHEDULESTATUS string `json:"@SCHEDULESTATUS"`
				SETPOINT       struct {
					Constraints struct {
						Error        []interface{} `json:"error"`
						IsConversion bool          `json:"isConversion"`
						LowerLimit   int           `json:"lowerLimit"`
						Unit         int           `json:"unit"`
						UpperLimit   int           `json:"upperLimit"`
						Warning      []struct {
							Message string `json:"message"`
							Value   int    `json:"value"`
						} `json:"warning"`
					} `json:"constraints"`
					Value int `json:"value"`
				} `json:"@SETPOINT"`
				STATUS  string `json:"@STATUS"`
				TCONFIG []struct {
					Align string `json:"align"`
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"@TCONFIG"`
				TYPE         string   `json:"@TYPE"`
				VACATION     bool     `json:"@VACATION"`
				Actions      []string `json:"actions"`
				DeviceName   string   `json:"device_name"`
				DeviceType   string   `json:"device_type"`
				MacAddress   string   `json:"mac_address"`
				SerialNumber string   `json:"serial_number"`
			} `json:"equiptments"`
			LocationID string `json:"location_id"`
		} `json:"locations"`
	} `json:"results"`
	Success bool   `json:"success"`
	Logs    string `json:"logs"`
	Stack   string `json:"stack"`
}
