# MPRIS Timer
[![image](https://github.com/user-attachments/assets/75651dc5-de7a-4244-974a-47ee69adac0f)](https://flathub.org/apps/io.github.efogdev.mpris-timer)

A timer app (CLI/GUI) with seamless GNOME integration accomplished by pretending to be a media player. \
Ultimately, serves the only purpose — to start a timer quickly and efficiently. \
Notifications included! Utilizing GTK4, Adwaita and MPRIS interface.

>MPRIS Timer aims to be as keyboard friendly as possible.
>Use navigation keys (arrows, tab, shift+tab, space, enter) or start inputting numbers right away.

![image](https://github.com/user-attachments/assets/3a6f6eb8-8e5f-4c16-a801-6e346bd4d100)

## Installation

```shell
flatpak install --user flathub io.github.efogdev.core-timer
```

## Preview

![2](https://github.com/user-attachments/assets/7be07479-85bb-44b1-9f6f-0fc85190c89e)

## CLI use

>No UI will be shown if run with `start` flag.
```text
Usage of mpris-timer:
  -color string
    	Progress color for the player (default "#F6D32D")
  -notify
    	Send desktop notification (default true)
  -silence int
    	Play this milliseconds of silence before the actual audio — might be helpful for audio devices that wake up not immediately
  -sound
    	Play sound (default true)
  -start int
    	Start the timer immediately
  -text string
    	Notification text (default "Time is up!")
  -title string
    	Name/title of the timer (default "Timer")
  -ui
    	Show timepicker UI (default true)
  -volume float
    	Volume [0-1] (default 1)
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
flatpak run org.flatpak.Builder --force-clean --sandbox --user --install --install-deps-from=flathub --ccache .build io.github.efogdev.core-timer.yml
```
