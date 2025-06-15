package webhook

import "github.com/thinkgos/logger"

func MediaServerId(v string) logger.Field {
	return logger.String("serverId", v)
}

func App(v string) logger.Field {
	return logger.String("app", v)
}
func Schema(v string) logger.Field {
	return logger.String("schema", v)
}

func Vhost(v string) logger.Field {
	return logger.String("vhost", v)
}

func Ip(v string) logger.Field {
	return logger.String("ip", v)
}

func Port(v int) logger.Field {
	return logger.Int("port", v)
}

func StreamId(v string) logger.Field {
	return logger.String("streamId", v)
}

func TcpId(v string) logger.Field {
	return logger.String("tcpId", v)
}
