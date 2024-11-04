# MPRIS Timer

![image](https://github.com/user-attachments/assets/80c40dee-1a2f-4729-8f9b-89e5eeb934b9)

>MPRIS Timer is really keyboard friendly! It should be quite intuitive. \
>Use navigation (arrows, tab, shift+tab) or start inputting numbers right away.

![ezgif-2-342b92ca02](https://github.com/user-attachments/assets/11e564dc-951e-4cc7-8215-ca1160ee0c0c)

Run:

```shell
go run cmd/main.go -help
```

Build:
```shell
go build -ldflags="-s -w" -o ./.bin/app ./cmd/main.go
```

Flatpak:
```shell
go run github.com/dennwc/flatpak-go-mod@latest .
flatpak-builder --user --force-clean .build io.github.efogdev.mpris-timer.yml
```

## ToDo

1) Customizable presets
2) Preferences dialog
