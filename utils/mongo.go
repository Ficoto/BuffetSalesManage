package utils

// Mongo connection setting
type mongoConn struct {
	authConn     authConn
	AuthSource   string
	ReplicateSet string
	URI          string
}

func (mc mongoConn) preStr() string {
	return "mongodb://"
}

func (mc *mongoConn) generateURI() {
	URI := mc.preStr() + mc.authConn.uri + "/"
	option := ""
	if mc.AuthSource != "" && mc.ReplicateSet != "" {
		option = "authSource=" + mc.AuthSource
		option += "&replicaSet=" + mc.ReplicateSet
	} else {
		if mc.AuthSource != "" {
			option = "authSource=" + mc.AuthSource
		}
		if mc.ReplicateSet != "" {
			option = "replicaSet=" + mc.ReplicateSet
		}
	}
	if option != "" {
		URI += "?" + option
	}
	mc.URI = URI
}
