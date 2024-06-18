# Battery indicator 
## A tool for displaying laptop battery status in i3 window manager status line written 

## Build and install

- Clone the repository 

```
git clone https://github.com/mikalai-dev/battery_indicator 
```

- Build the binary

```
go build -o battery main.go
```

- Copy the built tool to any apropriate directory (e.g. ~/.config/i3/bin/)

- Add the following lines to i3status.conf/i3blocks.conf

```
[battery]
command=~/.config/i3/scripts/battery
interval=30
```

replacing the path with path, where you've placed the ``battery`` binary
