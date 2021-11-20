package app

func Boot(options Options) {
	core := createCore(options)
	core.AppLogger.Info().Interface("config", core.Config).Msg("app boot")

	startHttpServer(core.HttpServer, core.HttpLogger)
}
