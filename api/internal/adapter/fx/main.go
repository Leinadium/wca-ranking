package fx

import "go.uber.org/fx"

func Main() {
	app := fx.New(
		ServicesModule,
	)

	app.Run()
}
