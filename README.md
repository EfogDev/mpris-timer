# MPRIS Timer

![image](https://github.com/user-attachments/assets/80c40dee-1a2f-4729-8f9b-89e5eeb934b9)

>MPRIS Timer is really keyboard friendly! It should be quite intuitive. \
>Use navigation (arrows, tab, shift+tab) or start inputting numbers right away.

![ezgif-3-054839fb4c](https://github.com/user-attachments/assets/7994f964-b18e-4254-a141-9b5c149b1483)

Run:

```shell
go run cmd/main.go -help
```

Build:
```shell
go build -tags native,waylan -ldflags="-s -w" -o ./.bin/app ./cmd/main.go
```

Flatpak:
```shell
go run github.com/dennwc/flatpak-go-mod@latest .
flatpak run org.flatpak.Builder --force-clean --sandbox --user --install --install-deps-from=flathub --ccache --mirror-screenshots-url=https://dl.flathub.org/media/ .build io.github.efogdev.mpris-timer.yml
```

## ToDo

1) Customizable presets
2) Preferences dialog
