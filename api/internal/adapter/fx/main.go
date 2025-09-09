package fx

import "go.uber.org/fx"

func Main() {
	app := fx.New(
		ConfigModule,
		StorageModule,
		WCAModule,
		ServicesModule,
	)

	app.Run()
}
