# Chip-8 emulator implementation

A pet-project to learn more about emulation and computers.

## Examples

<img width="752" alt="Screen Shot 2022-01-03 at 21 18 11" src="https://user-images.githubusercontent.com/1260049/147993953-f4f33f91-f66b-489e-b013-fee2f0605bd0.png">
<img width="752" alt="Screen Shot 2022-01-03 at 21 19 04" src="https://user-images.githubusercontent.com/1260049/147993956-5f149b93-b20b-481e-b68b-905ba9c21a49.png">
<img width="752" alt="Screen Shot 2022-01-03 at 21 19 16" src="https://user-images.githubusercontent.com/1260049/147993957-3b8d2cf6-c8f4-4762-8efe-21546be61e9b.png">
<img width="752" alt="Screen Shot 2022-01-03 at 21 19 35" src="https://user-images.githubusercontent.com/1260049/147993962-63f1679f-4e27-4d4e-88d0-d90ae231dae6.png">



## Requirements

- [go 1.17](https://go.dev)
- [go-sdl2](https://github.com/veandco/go-sdl2#requirements)

## Running

```
go run main.go <path/to/rom>
```

## Keymapping


|   |   |   |   |  
|-|-|-|-|
| 1 | 2 | 3 | 4 |
| Q | W | E | R |
| A | S | D | F |
| Z | X | C | V |

## TODO

- [ ] Sound
- [ ] Refactor keypad events
- [ ] Tunning emulation speed

## Acknowledgment

- [Chip-8 Technical Reference](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM)
- [How to write an emulator (CHIP-8 interpreter)](https://multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/)

