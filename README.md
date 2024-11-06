# MPRIS Timer
[![image](https://github.com/user-attachments/assets/75651dc5-de7a-4244-974a-47ee69adac0f)](https://flathub.org/apps/io.github.efogdev.mpris-timer)

A timer app (CLI/GUI) with seamless GNOME integration accomplished by pretending to be a media player. \
Ultimately, serves the only purpose — to start a timer quickly and efficiently. \
Notifications included! Utilizing GTK4, Adwaita and MPRIS interface.

>MPRIS Timer is really keyboard friendly! It should be quite intuitive. \
>Use navigation (arrows, tab, shift+tab) or start inputting numbers right away.

![image](https://github.com/user-attachments/assets/80c40dee-1a2f-4729-8f9b-89e5eeb934b9)

![ezgif-3-054839fb4c](https://github.com/user-attachments/assets/7994f964-b18e-4254-a141-9b5c149b1483)

## Installation

```shell
flatpak install --user flathub io.github.efogdev.mpris-timer
```

## Development

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
flatpak run org.flatpak.Builder --force-clean --sandbox --user --install --install-deps-from=flathub --ccache .build io.github.efogdev.mpris-timer.yml
```
